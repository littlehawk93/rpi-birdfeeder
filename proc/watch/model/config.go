package model

import (
	"time"

	"github.com/dhowden/raspicam"
)

// MotionSensorConfig configuration parameters for the IR motion detection sensor
type MotionSensorConfig struct {
	SignalPin string `mapstructure:"signal"`
}

// RangeFinderConfig configuration parameters fro the ultrasonic rangefinder sensor
type RangeFinderConfig struct {
	EchoPin    string `mapstructure:"echo"`
	TriggerPin string `mapstructure:"trigger"`
}

// WatchConfig configuration parameters for the Watch process. Configuration parameters are expected to not frequently change between application launches.
// The values are provided via a configuration file rather than through the command line
type WatchConfig struct {
	MotionSensor              *MotionSensorConfig `mapstructure:"motion_sensor"`
	RangeFinderSensor         *RangeFinderConfig  `mapstructure:"range_finder"`
	CameraConfig              *CameraConfig       `mapstructure:"camera"`
	OutputFolder              string              `mapstructure:"output_dir"`
	MinCaptureIntervalSeconds int                 `mapstructure:"min_capture_interval"`
}

// CameraConfig configuration parameters for the raspberry pi camera.
type CameraConfig struct {
	CaptureIntervalMillis int `mapstructure:"capture_interval"`
	CaptureIntervalCount  int `mapstructure:"capture_count"`
	Saturation            int `mapstructure:"saturation"`
	Brightness            int `mapstructure:"brightness"`
	Contrast              int `mapstructure:"contrast"`
	Sharpness             int `mapstructure:"sharpness"`
	Quality               int `mapstructure:"quality"`
	ExposureCompensation  int `mapstructure:"ev"`
}

// AsStill takes the data from this Camera Config and creates and populates a new still capture command
func (me CameraConfig) AsStill() *raspicam.Still {

	s := raspicam.NewStill()

	s.Timeout = time.Duration(me.CaptureIntervalCount) * time.Millisecond * time.Duration(me.CaptureIntervalMillis)
	s.Timelapse = me.CaptureIntervalCount

	s.Preview = raspicam.Preview{
		Mode: raspicam.PreviewDisabled,
	}

	s.Camera.Sharpness = me.wrapProperty(me.Sharpness, -100, 100)
	s.Camera.Contrast = me.wrapProperty(me.Contrast, -100, 100)
	s.Camera.Brightness = me.Brightness
	s.Camera.ExposureCompensation = me.ExposureCompensation

	s.Camera.ExposureMode = raspicam.ExposureAuto
	s.Camera.MeteringMode = raspicam.MeteringAverage
	s.Camera.AWBMode = raspicam.AWBAuto

	s.Quality = me.Quality
	s.Raw = false
	s.Encoding = raspicam.EncodingJPEG

	return s
}

func (me CameraConfig) wrapProperty(val, min, max int) int {

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}
