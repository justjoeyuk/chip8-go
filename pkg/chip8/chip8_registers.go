package chip8

type Registers struct {
	// 16 8-bit General Purpose Registers
	v [16]byte

	// Memory Address Register
	i uint16

	// Delay Timer Register
	dt byte

	// Sound Timer Register
	st byte

	// Program Counter Register
	pc uint16

	// Stack Pointer Register
	sp byte
}
