package chip8

// Registers - of the Chip8 hardware
type Registers struct {
	// 16 8-bit General Purpose Registers
	v [16]uint8

	// Memory Address Register
	i uint16

	// Delay Timer Register
	dt uint8

	// Sound Timer Register
	st uint8

	// Program Counter Register
	pc uint16

	// Stack Pointer Register
	sp uint8
}
