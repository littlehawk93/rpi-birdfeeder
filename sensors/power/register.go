package power

// register defines a register on the INA219 sensor
type register byte

const (
	regConfiguration register = 0x00
	regShuntVoltage  register = 0x01
	regBusVoltage    register = 0x02
	regPower         register = 0x03
	regCurrent       register = 0x04
	regCalibration   register = 0x05
)
