package powermon

import (
	"log"
	"time"

	influx "github.com/influxdata/influxdb1-client"

	"github.com/littlehawk93/go-ina219"
	"github.com/littlehawk93/rpi-birdfeeder/conf"
)

// Run the "main" function for the powermon process
func Run(config *conf.ApplicationConfig) {

	powerConfig := config.PowerMonConfig

	db, err := config.InfluxConfig.CreateClient()

	if err != nil {
		log.Fatalf("Unable to connect to InfluxDB: %s\n", err.Error())
	}

	ps, err := ina219.NewSensor(powerConfig.PowerSensor.Address, powerConfig.PowerSensor.Bus)

	if err != nil {
		log.Fatalf("Error initializing INA219 sensor: %s\n", err.Error())
	}

	defer ps.Close()

	sps, err := power.NewSensor(powerConfig.SolarPowerSensor.Address, powerConfig.SolarPowerSensor.Bus)

	if err != nil {
		log.Fatalf("Error initializing solar INA219 sensor: %s\n", err.Error())
	}

	defer sps.Close()

	sensors := []power.Sensor{*ps, *sps}

	labels := []string{"Bus Voltage", "Current Draw", "Power Draw", "Solar Bus Voltage", "Solar Current Draw", "Solar Power Draw"}

	funcs := []func() (float64, error){
		ps.GetBusVoltage,
		ps.GetCurrent,
		ps.GetPower,
		reverseSensorReading(sps.GetBusVoltage),
		reverseSensorReading(sps.GetCurrent),
		sps.GetPower,
	}

	for true {

		bp, err := generateBatchPoints(config.InfluxConfig.Database, config.PowerMonConfig.Measurement, config.PowerMonConfig.Tags, labels, funcs...)

		if err != nil {
			ps.Close()
			log.Fatalf("Unable to read sensor value(s): %s\n", err.Error())
		}

		if res, err := db.Write(bp); err != nil || (res != nil && res.Error() != nil) {
			ps.Close()

			if err != nil {
				log.Fatalf("Unable to write Influx data points: %s\n", err.Error())
			} else {
				log.Fatalf("Unable to write Influx data points: %s\n", res.Error().Error())
			}
		}

		if err = setPowerSavingMode(true, sensors); err != nil {
			ps.Close()
			sps.Close()
			log.Fatalf("Unable to enable INA219 Power Saving Mode: %s\n", err.Error())
		}

		time.Sleep(time.Duration(powerConfig.RefreshIntervalSeconds) * time.Second)

		if err = setPowerSavingMode(false, sensors); err != nil {
			ps.Close()
			sps.Close()
			log.Fatalf("Unable to disable INA219 Power Saving Mode: %s\n", err.Error())
		}
	}
}

func setPowerSavingMode(enabled bool, sensors []power.Sensor) error {

	for _, s := range sensors {
		if err := s.SetPowerSavingMode(enabled); err != nil {
			return err
		}
	}
	return nil
}

func reverseSensorReading(f func() (float64, error)) func() (float64, error) {

	return func() (float64, error) {
		val, err := f()

		if err != nil {
			return val, err
		}
		return val * -1.0, nil
	}
}

func generateBatchPoints(db, measurement string, tags map[string]string, labels []string, funcs ...func() (float64, error)) (influx.BatchPoints, error) {

	bp := influx.BatchPoints{
		Database: db,
		Points: []influx.Point{
			influx.Point{
				Time:        time.Now(),
				Measurement: measurement,
				Tags:        tags,
				Fields:      make(map[string]interface{}),
			},
		},
	}

	for i, f := range funcs {
		val, err := f()

		if err != nil {
			return bp, err
		}

		bp.Points[0].Fields[labels[i]] = val
	}

	return bp, nil
}
