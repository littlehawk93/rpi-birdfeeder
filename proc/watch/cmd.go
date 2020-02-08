package watch

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/raspicam"
	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/motion"
)

// Run the "main" function for the watch process
func Run(config *conf.ApplicationConfig) {

	watchConfig := config.WatchConfig

	motionSensor, err := motion.NewSensor(watchConfig.MotionSensor.SignalPin, makeOnMotionDetected(config.WatchConfig.CameraConfig, config.WatchConfig.OutputFolder))

	if err != nil {
		log.Fatalf("Error initializing motion sensor: %s\n", err.Error())
	}

	motionSensor.Begin()
	fmt.Scanln()
}

func makeOnMotionDetected(config *model.CameraConfig, outputDir string) func() {

	return func() {
		c := make(chan error)

		go func() {
			for err := range c {
				log.Printf("CAMERA ERROR: %s\n", err.Error())
			}
		}()

		s := config.AsStill()

		fileName := filepath.Join(outputDir, "CAPTURE.jpg")

		f, err := os.Create(fileName)

		if err != nil {
			log.Fatalf("Unable to create file '%s': %s\n", fileName, err.Error())
		}

		defer f.Close()

		raspicam.Capture(s, f, c)
	}
}
