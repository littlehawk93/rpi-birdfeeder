package conf

// InfluxConfig configuration parameters for writing data to InfluxDB
type InfluxConfig struct {
	Hostname string `mapstructure:"host"`
}
