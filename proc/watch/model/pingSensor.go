package model

import (
	"math"
	"time"

	"github.com/warthog618/gpio"
)

const (
	speedOfSound      float64 = 343.0
	echoTimeoutMillis         = 3
)

// PingSensor handles interfacing with the SR04 ultrasonic ping sensor
type PingSensor struct {
	echo      *gpio.Pin
	trig      *gpio.Pin
	closed    bool
	listening bool
}

// Ping sends a ping signal and attempts to return the measured distance, in meters.
// Returns math.NaN() if no signal is returned within the expected time frame (2cm <= distance <= 500cm) or if this sensor has already been closed
func (me *PingSensor) Ping() float64 {

	if me.closed {
		return math.NaN()
	}

	me.listening = true
	defer me.stopListening()

	me.trig.High()
	time.Sleep(10 * time.Microsecond)
	me.trig.Low()

	start := time.Now()

	select {
	case t := <-me.triggerChannel:
		return t.Sub(start).Seconds() * speedOfSound / 2.0
	case <-time.After(echoTimeoutMillis * time.Millisecond):
		return math.NaN()
	}

	return math.NaN()
}

// Close closes the underlying gpio connections to the referenced pins
func (me *PingSensor) Close() {
	if !me.closed {
		close(me.triggerChannel)
		me.closed = true
		me.listening = false
	}
}

func (me *PingSensor) stopListening() {
	me.listening = false
}

func (me *PingSensor) handlePingEcho(p *gpio.Pin) {
	if me.listening {
		me.triggerChannel <- time.Now()
	}
}

// NewPingSensor creates and initializes a new Ping Sensor
func NewPingSensor(echo, trig int) *PingSensor {
	p := &PingSensor{
		echo:           gpio.NewPin(echo),
		trig:           gpio.NewPin(trig),
		closed:         false,
		listening:      false,
		triggerChannel: make(chan time.Time),
	}
	p.trig.Output()
	p.echo.Watch(gpio.EdgeRising, p.handlePingEcho)
	return p
}
