package motion

import (
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// DetectHandler event handler whenever the motion sensor detects motion
type DetectHandler func()

// Sensor handles interfacing with the SR501 PIR motion sensor
type Sensor struct {
	pin     gpio.PinIO
	handler DetectHandler
}

// Begin begin listening for motion using the motion sensor.
// Any subsequent 'Begin' calls are ignored
func (me *Sensor) Begin() {
	go func() {
		for me.pin.WaitForEdge(-1) {
			me.handler()
		}
	}()
}

// NewSensor creates and initializes a new Motion Sensor
func NewSensor(pin string, handler DetectHandler) (*Sensor, error) {
	s := &Sensor{
		pin:     gpioreg.ByName(pin),
		handler: handler,
	}

	if err := s.pin.In(gpio.PullDown, gpio.RisingEdge); err != nil {
		return nil, err
	}

	return s, nil
}
