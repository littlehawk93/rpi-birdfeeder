package ping

import (
	"fmt"
	"math"
	"time"

	"github.com/warthog618/gpio"
)

const (
	// speedOfSound approximation of the speed of sound, in meters per second (m/s)
	speedOfSound float64 = 343.0

	// echoTimeoutMillis the maximum period to wait for a response from the echo pin before timing out (in milliseconds)
	// the effective range of the SR04 sensor is 500 cm or 5 meters.
	// The time it takes for the ping to go 5 meters out and back is roughly 29 milliseconds. (5 * 2 / 343 * 1000)
	echoTimeoutMillis = 30
)

// Sensor handles interfacing with the SR04 ultrasonic ping sensor
type Sensor struct {
	echo        *gpio.Pin
	trig        *gpio.Pin
	closed      bool
	listening   bool
	trigChannel chan time.Time
}

// Ping sends a ping signal and attempts to return the measured distance, in meters.
// Returns math.NaN() if no signal is returned within the expected time frame (2cm <= distance <= 500cm) or if this sensor has already been closed
func (me *Sensor) Ping() float64 {

	if me.closed {
		return math.NaN()
	}

	me.listening = true
	me.trig.High()
	time.Sleep(5 * time.Microsecond)
	me.trig.Low()

	start := time.Now()

	select {
	case t := <-me.trigChannel:
		me.listening = false
		return t.Sub(start).Seconds() * speedOfSound / 2.0
	case <-time.After(echoTimeoutMillis * time.Millisecond):
		me.listening = false
		return math.NaN()
	}

	me.listening = false
	return math.NaN()
}

// Close cleans up any allocated GPIO resources to run this sensor.
// Multiple calls to this method are safe.
func (me *Sensor) Close() {
	if !me.closed {
		me.closed = true
		close(me.trigChannel)
		me.echo.Unwatch()
	}
}

func (me *Sensor) onPingReceived(p *gpio.Pin) {

	if me.listening && !me.closed {
		me.trigChannel <- time.Now()
		fmt.Printf("[%s] PING DETECTED", time.Now().Format("2006-01-02 3:04:05.9999 PM"))
	}
}

// NewSensor creates and initializes a new Ping Sensor
func NewSensor(echo, trig int) *Sensor {
	p := &Sensor{
		echo:        gpio.NewPin(echo),
		trig:        gpio.NewPin(trig),
		closed:      false,
		listening:   false,
		trigChannel: make(chan time.Time),
	}
	p.trig.Output()
	p.trig.Low()

	p.echo.Watch(gpio.EdgeRising, p.onPingReceived)
	time.Sleep(echoTimeoutMillis * time.Millisecond)
	return p
}
