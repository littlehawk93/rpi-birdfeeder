package power

import (
	"math"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
)

const (
	busVoltageMultiplier  float64 = 0.001
	shuntVoltageMultipler float64 = 0.01
	powerDividend         float64 = 8192.0
	currentDivisor        float64 = 409.6
)

// Sensor handles interfacing with the INA219 power sensor via I2C
type Sensor struct {
	i2c                *i2c.Dev
	currentConfig      Configuration
	currentCalibration CalibrationSetting
}

// Close closes the underlying I2C connection for this sensor
func (me *Sensor) Close() error {
	return me.i2c.Bus.(i2c.BusCloser).Close()
}

// SetConfig set the configuration options for this sensor
func (me *Sensor) SetConfig(config Configuration) error {
	me.currentConfig = config
	return me.writeRegU16BE(regConfiguration, config.Flatten())
}

// GetBusVoltage returns the current voltage from the bus sensor in V, or any errors encountered while reading the sensor
func (me *Sensor) GetBusVoltage() (float64, error) {

	rawValue, err := me.readRegU16BE(regBusVoltage)

	if err != nil {
		return math.NaN(), err
	}

	signedRawValue := int16((rawValue >> 3) * 4)
	return float64(signedRawValue) * busVoltageMultiplier, nil
}

// GetShuntVoltage returns the current voltage from the shunt sensor in mV, or any errors encountered while reading the sensor
func (me *Sensor) GetShuntVoltage() (float64, error) {

	rawValue, err := me.readRegS16BE(regShuntVoltage)

	if err != nil {
		return math.NaN(), err
	}

	return float64(rawValue) * shuntVoltageMultipler, nil
}

// GetCurrent returns the current amp draw from the sensor in mA, or any errors encountered while reading the sensor
func (me *Sensor) GetCurrent() (float64, error) {

	rawValue, err := me.readRegS16BE(regCurrent)

	if err != nil {
		return math.NaN(), err
	}

	divider := float64(me.currentCalibration) / currentDivisor
	return float64(rawValue) / divider, nil
}

// GetPower returns the current power from the sensor in mW, or any errors encountered while reading the sensor
func (me *Sensor) GetPower() (float64, error) {

	rawValue, err := me.readRegS16BE(regPower)

	if err != nil {
		return math.NaN(), err
	}

	multiplier := powerDividend / float64(me.currentCalibration)
	return float64(rawValue) * multiplier, nil
}

// SetPowerSavingMode enables or disables the power saving operating mode on this sensor.
// Returns any errors that occur during operating mode change.
func (me *Sensor) SetPowerSavingMode(enabled bool) error {

	if enabled {
		me.currentConfig.OperatingMode = ModePowerDown
	} else {
		me.currentConfig.OperatingMode = ModeShuntBusContinuous
	}

	return me.SetConfig(me.currentConfig)
}

// SetCalibrationRegister set the calibration value for this sensor.
// Returns any errors that occur during calibration writing
func (me *Sensor) SetCalibrationRegister(val CalibrationSetting) error {
	me.currentCalibration = val

	return me.writeRegU16BE(regCalibration, uint16(val))
}

func (me *Sensor) writeRegU16BE(reg register, val uint16) error {
	_, err := me.i2c.Write([]byte{byte(reg), byte((val & 0xFF00) >> 8), byte(val & 0x00FF)})
	return err
}

func (me *Sensor) readRegS16BE(reg register) (int16, error) {

	res := make([]byte, 2)

	if err := me.i2c.Tx([]byte{byte(reg)}, res); err != nil {
		return 0, err
	}

	result := int16(res[0]) << 8
	result |= int16(res[1]) & 0x00FF
	return result, nil
}

func (me *Sensor) readRegU16BE(reg register) (uint16, error) {

	res := make([]byte, 2)

	if err := me.i2c.Tx([]byte{byte(reg)}, res); err != nil {
		return 0, err
	}

	result := uint16(res[0]) << 8
	result |= uint16(res[1]) & 0x00FF
	return result, nil
}

// NewSensor creates and initializes a new Power sensor
func NewSensor(address uint16, bus string) (*Sensor, error) {

	b, err := i2creg.Open("")

	if err != nil {
		return nil, err
	}

	s := &Sensor{
		i2c: &i2c.Dev{
			Bus:  b,
			Addr: address,
		},
	}

	if err = s.SetConfig(DefaultConfiguration()); err != nil {
		s.Close()
		return nil, err
	}

	if err = s.SetCalibrationRegister(V32A2); err != nil {
		s.Close()
		return nil, err
	}

	return s, nil
}
