package watch

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dhowden/raspicam"
	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/motion"
)

// Run the "main" function for the watch process
func Run(config *conf.ApplicationConfig) {

	lastCapture := time.Now()

	watchConfig := config.WatchConfig

	motionSensor, err := motion.NewSensor(watchConfig.MotionSensor.SignalPin, makeOnMotionDetected(watchConfig, &lastCapture))

	if err != nil {
		log.Fatalf("Error initializing motion sensor: %s\n", err.Error())
	}

	motionSensor.Begin()
	fmt.Scanln()
}

func makeOnMotionDetected(config *model.WatchConfig, lastCapture *time.Time) func() {

	cameraCfg := config.CameraConfig
	outputDir := config.OutputFolder

	return func() {

		if _, err := os.Stat(outputDir); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(outputDir, 0x755)
			} else {
				log.Fatalf("Unable to access directory '%s': %s\n", outputDir, err.Error())
			}
		}

		if time.Now().Sub(*lastCapture).Seconds() >= float64(config.MinCaptureIntervalSeconds) {

			log.Printf("Motion detected: capturing %d images\n", config.CameraConfig.CaptureIntervalCount)

			c := make(chan error)

			go func() {
				for err := range c {
					log.Printf("CAMERA ERROR: %s\n", err.Error())
				}
			}()

			s := cameraCfg.AsStill()

			fileName := createFileName(outputDir)

			f, err := os.Create(fileName)

			if err != nil {
				log.Fatalf("Unable to create file '%s': %s\n", fileName, err.Error())
			}

			defer f.Close()

			raspicam.Capture(s, f, c)

			*lastCapture = time.Now()
		} else {
			log.Println("Motion detected, but last image capture was too recent")
		}
	}
}

func createFileName(dir string) string {
	return filepath.Join(dir, strings.Join([]string{time.Now().Format("20060102030405"), "CAPTURE", "%04d.jpg"}, "_"))
}
