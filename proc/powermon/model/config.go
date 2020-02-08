package model

// PowerMonConfig configuration parameters for the power monitor command
type PowerMonConfig struct {
	RefreshIntervalSeconds int                `mapstructure:"refresh_interval"`
	PowerSensor            *PowerSensorConfig `mapstructure:"power_sensor"`
	InfluxMeasurement      string             `mapstructure:"measurement"`
	InfluxTags             map[string]string  `mapstructure:"tags"`
}

// PowerSensorConfig configuration parameters for the INA219 power sensor
type PowerSensorConfig struct {
	Address uint16 `mapstructure:"address"`
	Bus     string `mapstructure:"bus"`
}
