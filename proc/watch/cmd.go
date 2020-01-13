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

	signalPin := gpio.NewPin(config.MotionSensor.SignalPin)
	signalPin.Input()

	signalPin.Watch(gpio.EdgeBoth, onPinEdgeDetected)
	defer signalPin.Unwatch()

	fmt.Scanln()
}

func onPinEdgeDetected(p *gpio.Pin) {

	dateStr := time.Now().Format("1/2/06 03:04:05 PM")

	if p.Read() == gpio.Low {
		fmt.Printf("[%s] Pin %d LOW\n", dateStr, p.Pin())
	} else {
		fmt.Printf("[%s] Pin %d HIGH\n", dateStr, p.Pin())
	}
}
