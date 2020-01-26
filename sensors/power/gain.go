package power

// Gain the voltage gain
type Gain uint16

const (
	// Gain40 40mV range
	Gain40 Gain = 0x0000

	// Gain80 80mV range
	Gain80 Gain = 0x0800

	// Gain160 160mV range
	Gain160 Gain = 0x1000

	// Gain320 320mV range
	Gain320 Gain = 0x1800
)
