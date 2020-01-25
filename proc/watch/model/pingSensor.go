package model

import (
	"math"
	"time"

	"github.com/warthog618/gpio"
)

const (
	speedOfSound      float64 = 343.0
	echoTimeoutMillis         = 30
)

// PingSensor handles interfacing with the SR04 ultrasonic ping sensor
type PingSensor struct {
	echo *gpio.Pin
	trig *gpio.Pin
}

// Ping sends a ping signal and attempts to return the measured distance, in meters.
// Returns math.NaN() if no signal is returned within the expected time frame (2cm <= distance <= 500cm) or if this sensor has already been closed
func (me *PingSensor) Ping() float64 {

	c := make(chan time.Time)
	received := false

	defer me.echo.Unwatch()

	me.echo.Watch(gpio.EdgeRising, func(p *gpio.Pin) {
		t := time.Now()
		if !received {
			received = true
			c <- t
			me.echo.Unwatch()
			close(c)
		}
	})

	me.trig.High()
	time.Sleep(10 * time.Microsecond)
	me.trig.Low()

	start := time.Now()

	select {
	case t := <-c:
		return t.Sub(start).Seconds() * speedOfSound / 2.0
	case <-time.After(echoTimeoutMillis * time.Millisecond):
		return math.NaN()
	}

	return math.NaN()
}

// NewPingSensor creates and initializes a new Ping Sensor
func NewPingSensor(echo, trig int) *PingSensor {
	p := &PingSensor{
		echo: gpio.NewPin(echo),
		trig: gpio.NewPin(trig),
	}
	p.trig.Output()
	p.trig.Low()
	time.Sleep(echoTimeoutMillis * time.Millisecond)
	return p
}
