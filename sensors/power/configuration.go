package power

// Configuration INA219 configuration bit flags
type Configuration struct {
	Tolerance       BusVoltageTolerance
	ShuntResolution ShuntResolution
	BusResolution   BusResolution
	OperatingMode   OperatingMode
	Gain            Gain
}

// Flatten combines all the bit flags in this configuration into a single 16 bit value
func (me Configuration) Flatten() uint16 {
	return uint16(me.Tolerance) | uint16(me.ShuntResolution) | uint16(me.BusResolution) | uint16(me.OperatingMode) | uint16(me.Gain)
}

// DefaultConfiguration returns the default configuration for the INA219 sensor
func DefaultConfiguration() Configuration {
	return Configuration{
		Tolerance:       BusVoltageRange32,
		Gain:            Gain320,
		ShuntResolution: Res12Bit,
		BusResolution:   BusRes12,
		OperatingMode:   ModeShuntBusContinuous,
	}
}
