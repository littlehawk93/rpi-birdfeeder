package power

// CalibrationSetting predefined sensor calibration settings
type CalibrationSetting uint16

const (
	// V32A2 calibration for 32 Volts, 2 Amps
	V32A2 CalibrationSetting = 4096

	// V32A1 calibration for 32 Volts, 1 Amp
	V32A1 CalibrationSetting = 10240

	// V16A04 calibration for 16 Volts, 400 mA
	V16A04 CalibrationSetting = 8192
)
