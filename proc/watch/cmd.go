package watch

import (
	"fmt"
	"log"
	"sync"
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

	running := true

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for running {
			fmt.Printf("%5.3f\n", pingSensor.Ping())
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	fmt.Scanln()
	running = false
	wg.Wait()
}
