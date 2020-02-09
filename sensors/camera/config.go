package camera

import "fmt"

// Config configuration parameters for the raspberry pi camera.
type Config struct {
	CaptureInterval      int `mapstructure:"capture_interval"`
	CaptureIntervalCount int `mapstructure:"capture_count"`
	Saturation           int `mapstructure:"saturation"`
	Brightness           int `mapstructure:"brightness"`
	Contrast             int `mapstructure:"contrast"`
	Sharpness            int `mapstructure:"sharpness"`
	Quality              int `mapstructure:"quality"`
	ExposureCompensation int `mapstructure:"ev"`
}

func (me Config) params() []string {

	params := make([]string, 0)
	params = append(params, "--nopreview")

	params = addConfigIntProperty(params, "saturation", wrapConfigProperty(me.Saturation, -100, 100))
	params = addConfigIntProperty(params, "sharpness", wrapConfigProperty(me.Sharpness, -100, 100))
	params = addConfigIntProperty(params, "brightness", wrapConfigProperty(me.Brightness, 0, 100))
	params = addConfigIntProperty(params, "contrast", wrapConfigProperty(me.Contrast, -100, 100))
	params = addConfigIntProperty(params, "quality", wrapConfigProperty(me.Quality, 0, 100))
	params = addConfigIntProperty(params, "ev", wrapConfigProperty(me.ExposureCompensation, -10, 10))

	intervalMillis := wrapConfigProperty(me.CaptureInterval, 1, 5) * 1000

	params = addConfigIntProperty(params, "timeout", wrapConfigProperty(me.CaptureIntervalCount, 1, 10)*intervalMillis+100)
	params = addConfigIntProperty(params, "timelapse", intervalMillis)

	params = addConfigProperty(params, "awb", "auto")
	params = addConfigProperty(params, "exposure", "auto")
	params = addConfigProperty(params, "metering", "average")

	return params
}

func addConfigProperty(params []string, name, value string) []string {
	return append(append(params, fmt.Sprintf("--%s", name)), value)
}

func addConfigIntProperty(params []string, name string, value int) []string {
	return append(append(params, fmt.Sprintf("--%s", name)), fmt.Sprintf("%d", value))
}

func wrapConfigProperty(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}
