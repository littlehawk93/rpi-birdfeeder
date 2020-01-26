package power

// BusVoltageTolerance defines the tolerated range of voltage values
type BusVoltageTolerance uint16

const (
	// BusVoltageRange16 0-16 Volt range
	BusVoltageRange16 BusVoltageTolerance = 0x0000

	// BusVoltageRange32 0-32 Volt range
	BusVoltageRange32 BusVoltageTolerance = 0x2000
)
