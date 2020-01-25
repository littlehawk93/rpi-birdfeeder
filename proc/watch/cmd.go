package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
	"github.com/warthog618/gpio"
)

// Run generates the "main" function for the watch process
func Run(config *model.WatchConfig) {

	err := gpio.Open()

	if err != nil {
		log.Fatalln(err)
	}

	defer gpio.Close()

	pingSensor := model.NewPingSensor(config.RangeFinderSensor.EchoPin, config.RangeFinderSensor.TriggerPin)

	defer pingSensor.Close()

	for true {
		fmt.Printf("%5.3f\n", pingSensor.Ping())
		time.Sleep(500 * time.Millisecond)
	}
}
