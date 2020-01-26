package proc

import (
	"github.com/littlehawk93/rpi-birdfeeder/conf"
)

// CommandProcess defines a sub-command process function
type CommandProcess func(cfg *conf.ApplicationConfig)
