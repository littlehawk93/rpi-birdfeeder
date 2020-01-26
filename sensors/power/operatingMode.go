package power

// OperatingMode defines an operating mode for the INA219 sensor
type OperatingMode uint16

const (
	// ModePowerDown power down the sensor
	ModePowerDown OperatingMode = 0x0000

	// ModeShuntTriggered read shunt sensor values
	ModeShuntTriggered OperatingMode = 0x0001

	// ModeBusTriggered read bus sensor values
	ModeBusTriggered OperatingMode = 0x0002

	// ModeShuntBusTriggered read shunt and bus sensor values
	ModeShuntBusTriggered OperatingMode = 0x0003

	// ModeADCOff turn ADC off
	ModeADCOff OperatingMode = 0x0004

	// ModeShuntContinuous read shunt sensor values continuously
	ModeShuntContinuous OperatingMode = 0x0005

	// ModeBusContinuous read bus sensor values continuously
	ModeBusContinuous OperatingMode = 0x0006

	// ModeShuntBusContinuous read shunt and bus sensor values continuously
	ModeShuntBusContinuous OperatingMode = 0x0007
)
