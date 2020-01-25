package model

import "github.com/warthog618/gpio"

// MotionDetectedHandler event handler whenever the motion sensor detects motion
type MotionDetectedHandler func()

// MotionSensor handles interfacing with the SR501 PIR motion sensor
type MotionSensor struct {
	closed    bool
	listening bool
	pin       *gpio.Pin
	handler   MotionDetectedHandler
}

// Begin begin listening for motion using the motion sensor.
// Any subsequent 'Begin' calls are ignored
func (me *MotionSensor) Begin() {
	if !me.listening && !me.closed {
		me.pin.Watch(gpio.EdgeRising, func(p *gpio.Pin) {
			me.handler()
		})
	}
}

// Close closes the underlying gpio connections to the referenced pins
func (me *MotionSensor) Close() {
	if !me.closed {
		me.pin.Unwatch()
		me.closed = true
		me.listening = false
	}
}

// NewMotionSensor creates and initializes a new Motion Sensor
func NewMotionSensor(pin int, handler MotionDetectedHandler) *MotionSensor {
	return &MotionSensor{
		listening: false,
		closed:    false,
		pin:       gpio.NewPin(pin),
		handler:   handler,
	}
}
