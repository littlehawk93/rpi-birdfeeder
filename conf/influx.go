package conf

import (
	"net/url"

	influx "github.com/influxdata/influxdb1-client"
)

// InfluxConfig configuration parameters for writing data to InfluxDB
type InfluxConfig struct {
	Database    string            `mapstructure:"db"`
	Measurement string            `mapstructure:"measurement"`
	Hostname    string            `mapstructure:"host"`
	Username    string            `mapstructure:"user"`
	Password    string            `mapstructure:"pass"`
	Tags        map[string]string `mapstructure:"tags"`
}

// CreateClient creates a new Influx DB client based on this config's properties
func (me InfluxConfig) CreateClient() (*influx.Client, error) {

	return influx.NewClient(influx.Config{
		URL: url.URL{
			Scheme: "http",
			Host:   me.Hostname,
		},
		Username: me.Username,
		Password: me.Password,
	})
}
