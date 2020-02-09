package model

import "github.com/littlehawk93/rpi-birdfeeder/camera"

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
	CameraConfig              *camera.Config      `mapstructure:"camera"`
	OutputFolder              string              `mapstructure:"output_dir"`
	MinCaptureIntervalSeconds int                 `mapstructure:"min_capture_interval"`
}
