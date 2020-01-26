package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/motion"
)

// Run the "main" function for the watch process
func Run(config *conf.ApplicationConfig) {

	watchConfig := config.WatchConfig

	motionSensor, err := motion.NewSensor(watchConfig.MotionSensor.SignalPin, onMotionDetected)

	if err != nil {
		log.Fatalf("Error initializing motion sensor: %s\n", err.Error())
	}

	motionSensor.Begin()
	fmt.Scanln()
}

func onMotionDetected() {
	fmt.Printf("[%s] MOTION DETECTED\n", time.Now().Format("2006-01-02 3:04:05.9999"))
}
