package powermon

import (
	"fmt"
	"log"
	"time"

	influx "github.com/influxdata/influxdb1-client"

	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/power"
)

// Run the "main" function for the powermon process
func Run(config *conf.ApplicationConfig) {

	powerConfig := config.PowerMonConfig

	db, err := config.InfluxConfig.CreateClient()

	if err != nil {
		log.Fatalf("Unable to connect to InfluxDB: %s\n", err.Error())
	}

	ps, err := power.NewSensor(powerConfig.PowerSensor.Address, powerConfig.PowerSensor.Bus)

	if err != nil {
		log.Fatalf("Error initializing INA219 sensor: %s\n", err.Error())
	}

	defer ps.Close()

	running := true

	go func() {
		for running {

			bp, err := generateBatchPoints(config.InfluxConfig.Database, config.InfluxConfig.Measurement, config.InfluxConfig.Tags, []string{"Bus Voltage", "Current Draw", "Power Draw"}, ps.GetBusVoltage, ps.GetCurrent, ps.GetPower)

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

			if err = ps.SetPowerSavingMode(true); err != nil {
				ps.Close()
				log.Fatalf("Unable to enable INA219 Power Saving Mode: %s\n", err.Error())
			}

			time.Sleep(time.Duration(powerConfig.RefreshIntervalSeconds) * time.Second)

			if err = ps.SetPowerSavingMode(false); err != nil {
				ps.Close()
				log.Fatalf("Unable to disable INA219 Power Saving Mode: %s\n", err.Error())
			}
		}
	}()

	fmt.Scanln()
	running = false
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
