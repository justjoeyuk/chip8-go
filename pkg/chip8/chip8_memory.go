package chip8

// ProgramStartAddress - The start address of most CHIP8 Programs
const ProgramStartAddress = 0x200

// HexSpriteAddress - The address in RAM where the Character Sprites are saved
const HexSpriteAddress = 0x00

// Sprites 0 to F, to be stored in Memory at location 0x00
func defaultCharacterSet() []byte {
	return []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
}

// Memory - Interpretation of the CHIP8 Memory and Stack
type Memory struct {
	// 4K RAM
	ram [4096]byte

	// Stack - 16 levels of subroutine nesting
	stack [16]uint16

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

/*
	NewMemory - Returns an instance of the Chip8 with the
	Default Character Set at memory location 0x00
*/
func NewMemory() *Memory {
	m := &Memory{}

	copy(m.ram[HexSpriteAddress:], defaultCharacterSet())

	return m
}

// Set a byte at a given memory address location
func (m *Memory) Set(loc int, val byte) {
	m.ram[loc] = val
}

// Get a byte from a given memory address location
func (m *Memory) Get(loc int) byte {
	return m.ram[loc]
}

// GetTwoBytes (a uint16) from memory
func (m *Memory) GetTwoBytes(loc int) uint16 {
	b1 := uint16(m.Get(loc))
	b2 := uint16(m.Get(loc + 1))

	return b1<<8 | b2
}

// GetNBytes from a given location in memory, returned as an array of bytes
func (m *Memory) GetNBytes(loc, numBytes int) []byte {
	bytes := make([]byte, numBytes)

	for i := 0; i < numBytes; i++ {
		bytes[i] = m.ram[loc+i]
	}

	return bytes
}

// PushStack - Push an address on to the stack and increment the Stack Pointer
func (m *Memory) PushStack(val uint16) {
	m.stack[m.sp+1] = val
	m.sp += 1
}

// PushStack - Pop an address from the stack and decrement the Stack Pointer
func (m *Memory) PopStack() uint16 {
	addr := m.stack[m.sp]
	m.sp -= 1

	return addr
}
