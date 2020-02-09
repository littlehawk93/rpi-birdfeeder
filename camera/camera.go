package camera

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os/exec"
)

const (
	raspistillCmd = "raspistill"
)

// CaptureTimelapse captures a series of images using the provided camera configuration
// and then saves the images using the provided file name template. Returns any errors that occurred during execution
func CaptureTimelapse(fileName string, cfg *Config) error {

	params := cfg.params()

	params = addConfigProperty(params, "output", fileName)

	cmd := exec.Command(raspistillCmd, params...)

	cmd.Stdout = ioutil.Discard
	b := &bytes.Buffer{}
	cmd.Stderr = b

	if err := cmd.Run(); err != nil {
		return errors.New(string(b.Bytes()))
	}

	return nil
}
