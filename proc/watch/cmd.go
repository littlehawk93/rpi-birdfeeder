package watch

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/littlehawk93/go-sr501"
	"github.com/littlehawk93/rpi-birdfeeder/conf"
	"github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
	"github.com/littlehawk93/rpi-birdfeeder/sensors/camera"
)

// Run the "main" function for the watch process
func Run(config *conf.ApplicationConfig) {

	lastCapture := time.Now()

	watchConfig := config.WatchConfig

	sensor, err := sr501.NewSensor(watchConfig.MotionSensor.SignalPin, makeOnMotionDetected(watchConfig, &lastCapture))

	if err != nil {
		log.Fatalf("Error initializing motion sensor: %s\n", err.Error())
	}

	defer sensor.Close()

	sensor.Begin()
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

			fileName := createFileName(outputDir)

			if err := camera.CaptureTimelapse(fileName, cameraCfg); err != nil {
				log.Fatalf("Error capturing timelapse images:\n%s\n", err.Error())
			}

			*lastCapture = time.Now()
		} else {
			log.Println("Motion detected, but last image capture was too recent")
		}
	}
}

func createFileName(dir string) string {
	return filepath.Join(dir, strings.Join([]string{time.Now().Format("0601020304"), "CAPTURE", "%02d.jpg"}, "_"))
}
