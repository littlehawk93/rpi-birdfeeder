package power

// BusResolution the ADC bit resolution for the bus
type BusResolution uint16

const (
	// BusRes9 9 bit resolution
	BusRes9 BusResolution = 0x0000

	// BusRes10 10 bit resolution
	BusRes10 BusResolution = 0x0080

	// BusRes11 11 bit resolution
	BusRes11 BusResolution = 0x0100

	// BusRes12 12 bit resolution
	BusRes12 BusResolution = 0x0180
)
