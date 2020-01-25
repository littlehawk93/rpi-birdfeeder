package motion

import "github.com/warthog618/gpio"

// DetectHandler event handler whenever the motion sensor detects motion
type DetectHandler func()

// Sensor handles interfacing with the SR501 PIR motion sensor
type Sensor struct {
	closed    bool
	listening bool
	pin       *gpio.Pin
	handler   DetectHandler
}

// Begin begin listening for motion using the motion sensor.
// Any subsequent 'Begin' calls are ignored
func (me *Sensor) Begin() {
	if !me.listening && !me.closed {
		me.pin.Watch(gpio.EdgeRising, func(p *gpio.Pin) {
			me.handler()
		})
	}
}

// Close closes the underlying gpio connections to the referenced pins
func (me *Sensor) Close() {
	if !me.closed {
		me.pin.Unwatch()
		me.closed = true
		me.listening = false
	}
}

// NewSensor creates and initializes a new Motion Sensor
func NewSensor(pin int, handler DetectHandler) *Sensor {
	return &Sensor{
		listening: false,
		closed:    false,
		pin:       gpio.NewPin(pin),
		handler:   handler,
	}
}
