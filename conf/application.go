package conf

import (
	powermonmodel "github.com/littlehawk93/rpi-birdfeeder/proc/powermon/model"
	watchmodel "github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
)

// ApplicationConfig configuration parameters for the rpi-birdfeeder application
type ApplicationConfig struct {
	PowerMonConfig *powermonmodel.PowerMonConfig `mapstructure:"powermon"`
	WatchConfig    *watchmodel.WatchConfig       `mapstructure:"watch"`
	LogConfig      *LogConfig                    `mapstructure:"log"`
	InfluxConfig   *InfluxConfig                 `mapstructure:"influx"`
}
