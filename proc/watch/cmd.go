package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/motion"
	"github.com/warthog618/gpio"
)

// Run generates the "main" function for the watch process
func Run(config *model.WatchConfig) {

	err := gpio.Open()

	if err != nil {
		log.Fatalln(err)
	}

	defer gpio.Close()

	motionSensor := motion.NewSensor(config.MotionSensor.SignalPin, onMotionDetected)

	defer motionSensor.Close()
	motionSensor.Begin()

	fmt.Scanln()
}

func onMotionDetected() {
	fmt.Printf("[%s] MOTION DETECTED\n", time.Now().Format("2006-01-02 3:04:05.9999"))
}
