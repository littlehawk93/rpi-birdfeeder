package power

// ShuntResolution the bit resolution for samples taken by the shunt sensor
type ShuntResolution uint16

const (
	// Res9Bit 9 bit ADC resolution
	Res9Bit ShuntResolution = 0x0000

	// Res10Bit 10 bit ADC resolution
	Res10Bit ShuntResolution = 0x0008

	// Res11Bit 11 bit ADC resolution
	Res11Bit ShuntResolution = 0x0010

	// Res12Bit 12 bit ADC resolution
	Res12Bit ShuntResolution = 0x0018

	// Res12BitX2 12 bit ADC resolution sampled 2 times
	Res12BitX2 ShuntResolution = 0x0048

	// Res12BitX4 12 bit ADC resolution sampled 4 times
	Res12BitX4 ShuntResolution = 0x0050

	// Res12BitX8 12 bit ADC resolution sampled 8 times
	Res12BitX8 ShuntResolution = 0x0058

	// Res12BitX16 12 bit ADC resolution sampled 16 times
	Res12BitX16 ShuntResolution = 0x0060

	// Res12BitX32 12 bit ADC resolution sampled 32 times
	Res12BitX32 ShuntResolution = 0x0068

	// Res12BitX64 12 bit ADC resolution sampled 64 times
	Res12BitX64 ShuntResolution = 0x0070

	// Res12BitX128 12 bit ADC resolution sampled 128 times
	Res12BitX128 ShuntResolution = 0x0078
)
