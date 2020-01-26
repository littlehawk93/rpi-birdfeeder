package ping

import (
	"math"
	"time"

	"periph.io/x/periph/conn/gpio/gpioreg"

	"periph.io/x/periph/conn/gpio"
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
	echo gpio.PinIO
	trig gpio.PinIO
}

// Ping sends a ping signal and attempts to return the measured distance, in meters.
// Returns math.NaN() if no signal is returned within the expected time frame (2cm <= distance <= 500cm) or if this sensor has already been closed.
// Returns an error if there was any communication errors with the ping sensor
func (me *Sensor) Ping() (float64, error) {

	if err := me.trig.Out(gpio.High); err != nil {
		return 0, err
	}

	time.Sleep(5 * time.Microsecond)

	if err := me.trig.Out(gpio.Low); err != nil {
		return 0, err
	}

	start := time.Now()

	if me.trig.WaitForEdge(echoTimeoutMillis * time.Millisecond) {
		return time.Now().Sub(start).Seconds() * speedOfSound / 2.0, nil
	}

	return math.NaN(), nil
}

// Close cleans up any allocated GPIO resources to run this sensor.
// Multiple calls to this method are safe.
func (me *Sensor) Close() {
	me.trig.Halt()
}

// NewSensor creates and initializes a new Ping Sensor.
// Returns the new sensor or any errors from attempting to create it
func NewSensor(echo, trig string) (*Sensor, error) {

	p := &Sensor{
		echo: gpioreg.ByName(echo),
		trig: gpioreg.ByName(trig),
	}

	if err := p.trig.Out(gpio.Low); err != nil {
		return nil, err
	}

	if err := p.echo.In(gpio.PullDown, gpio.RisingEdge); err != nil {
		return nil, err
	}
	time.Sleep(echoTimeoutMillis * time.Millisecond)
	return p, nil
}
