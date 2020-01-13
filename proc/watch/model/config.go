package model

// MotionSensorConfig configuration parameters for the IR motion detection sensor
type MotionSensorConfig struct {
	SignalPin int `mapstructure:"signal"`
}

// WatchConfig configuration parameters for the Watch process. Configuration parameters are expected to not frequently change between application launches.
// The values are provided via a configuration file rather than through the command line
type WatchConfig struct {
	MotionSensor *MotionSensorConfig `mapstructure:"motion_sensor"`
}
