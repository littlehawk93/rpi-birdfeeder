package powermon

import (
	"log"
	"time"

	influx "github.com/influxdata/influxdb1-client"
	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/ina219"
)

// Run the "main" function for the powermon process
func Run(config *conf.ApplicationConfig) {

	powerConfig := config.PowerMonConfig

	db, err := config.InfluxConfig.CreateClient()

	if err != nil {
		log.Fatalf("Unable to connect to InfluxDB: %s\n", err.Error())
	}

	bus, err := i2creg.Open("")

	if err != nil {
		log.Fatalf("Failed to open i2c bus: %s\n", err.Error())
	}

	ps, err := ina219.New(bus, &ina219.Opts{
		Address:       int(powerConfig.PowerSensor.Address),
		SenseResistor: 100 * physic.MilliOhm,
		MaxCurrent:    3200 * physic.MilliAmpere,
	})

	if err != nil {
		log.Fatalf("Error initializing INA219 sensor: %s\n", err.Error())
	}

	sps, err := ina219.New(bus, &ina219.Opts{
		Address:       int(powerConfig.SolarPowerSensor.Address),
		SenseResistor: 100 * physic.MilliOhm,
		MaxCurrent:    3200 * physic.MilliAmpere,
	})

	if err != nil {
		log.Fatalf("Error initializing solar INA219 sensor: %s\n", err.Error())
	}

	for true {
		powerVals, err := ps.Sense()

		if err != nil {
			log.Fatalf("Unable to read sensor value(s): %s\n", err.Error())
		}

		solarVals, err := sps.Sense()

		if err != nil {
			log.Fatalf("Unable to read solar sensor value(s): %s\n", err.Error())
		}

		bp := influx.BatchPoints{
			Database: config.InfluxConfig.Database,
			Points: []influx.Point{
				influx.Point{
					Time:        time.Now(),
					Measurement: powerConfig.Measurement,
					Tags:        powerConfig.Tags,
					Fields:      make(map[string]interface{}),
				},
			},
		}

		bp.Points[0].Fields["Bus Voltage"] = float64(powerVals.Voltage) / float64(physic.Volt)
		bp.Points[0].Fields["Current Draw"] = float64(powerVals.Current) / float64(physic.MilliAmpere)
		bp.Points[0].Fields["Power Draw"] = float64(powerVals.Power) / float64(physic.MilliWatt)
		bp.Points[0].Fields["Solar Bus Voltage"] = float64(solarVals.Voltage) / float64(physic.Volt)
		bp.Points[0].Fields["Solar Current Draw"] = float64(solarVals.Voltage) / float64(physic.MilliAmpere)
		bp.Points[0].Fields["Solar Power Draw"] = float64(solarVals.Power) / float64(physic.MilliWatt)

		if res, err := db.Write(bp); err != nil || (res != nil && res.Error() != nil) {
			if err != nil {
				log.Fatalf("Unable to write Influx data points: %s\n", err.Error())
			} else {
				log.Fatalf("Unable to write Influx data points: %s\n", res.Error().Error())
			}
		}

		time.Sleep(time.Duration(powerConfig.RefreshIntervalSeconds) * time.Second)
	}
}
