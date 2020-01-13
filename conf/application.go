package conf

import (
	watchmodel "github.com/littlehawk93/rpi-birdfeeder/proc/watch/model"
)

// ApplicationConfig configuration parameters for the rpi-birdfeeder application
type ApplicationConfig struct {
	WatchConfig *watchmodel.WatchConfig `mapstructure:"watch"`
	LogConfig   *LogConfig              `mapstructure:"log"`
}
