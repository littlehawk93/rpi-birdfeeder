package powermon

import (
	"fmt"
	"log"
	"time"

	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/power"
)

// Run the "main" function for the powermon process
func Run(config *conf.ApplicationConfig) {

	powerConfig := config.PowerMonConfig

	powerSensor, err := power.NewSensor(powerConfig.PowerSensor.Address, powerConfig.PowerSensor.Bus)

	if err != nil {
		log.Fatalf("Error initializing INA219 sensor: %s\n", err.Error())
	}

	defer powerSensor.Close()

	running := true

	go func() {
		for running {

			if err := printSensorValues([]string{"Bus Voltage (V)", "Shunt Voltage (mV)", "Current (mA)", "Power (mW)"}, powerSensor.GetBusVoltage, powerSensor.GetShuntVoltage, powerSensor.GetCurrent, powerSensor.GetPower); err != nil {
				powerSensor.Close()
				log.Fatalf("Error reading sensor value: %s\n", err.Error())
			}

			time.Sleep(time.Duration(powerConfig.RefreshIntervalSeconds) * time.Second)
		}
	}()

	fmt.Scanln()
	running = false
}

func printSensorValues(labels []string, funcs ...func() (float64, error)) error {

	for i, f := range funcs {

		val, err := f()

		if err != nil {
			return err
		}

		fmt.Printf("%s: %6.3f\n", labels[i], val)
	}
	fmt.Println()

	return nil
}
