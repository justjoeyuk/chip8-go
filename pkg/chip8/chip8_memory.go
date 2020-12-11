package chip8

// ProgramStartAddress - The start address of most CHIP8 Programs
const ProgramStartAddress = 0x200

// HexSpriteAddress - The address in RAM where the Character Sprites are saved
const HexSpriteAddress = 0x00

// Sprites 0 to F, to be stored in Memory at location 0x00
func defaultCharacterSet() []uint8 {
	return []uint8{
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

// Memory - Interpretation of the Chip8 Memory and Stack
type Memory struct {
	// 4K RAM
	ram [4096]uint8

	// Stack - 16 levels of subroutine nesting
	stack [16]uint16

	// Registers - The collection of registers (and meta registers) on the Chip8
	registers Registers
}

// NewMemory - Returns an instance of the Chip8 with theDefault Character Set at memory location 0x00
func NewMemory(romData []uint8) *Memory {
	m := &Memory{}

	if len(romData) > len(m.ram)-ProgramStartAddress {
		panic("Could not load ROM into emulator memory. Too big.")
	}

	copy(m.ram[HexSpriteAddress:], defaultCharacterSet())
	copy(m.ram[ProgramStartAddress:], romData)

	return m
}

// Set a byte at a given memory address location
func (m *Memory) Set(loc uint16, val uint8) {
	m.ram[loc] = val
}

// Get a byte from a given memory address location
func (m *Memory) Get(loc uint16) uint8 {
	return m.ram[loc]
}

// GetTwoBytes (a uint16) from memory
func (m *Memory) GetTwoBytes(loc uint16) uint16 {
	b1 := m.Get(loc)
	b2 := m.Get(loc + 1)

	return uint16(b1)<<8 | uint16(b2)
}

// GetNBytes from a given location in memory, returned as an array of bytes
func (m *Memory) GetNBytes(loc uint16, numBytes int) []uint8 {
	bytes := make([]byte, numBytes)

	for i := 0; i < numBytes; i++ {
		bytes[i] = m.ram[int(loc)+i]
	}

	return bytes
}

// PushStack - Push an address on to the stack and increment the Stack Pointer
func (m *Memory) PushStack(val uint16) {
	m.stack[m.registers.sp+1] = val
	m.registers.sp += 1
}

// PopStack - Pop an address from the stack and decrement the Stack Pointer
func (m *Memory) PopStack() uint16 {
	addr := m.stack[m.registers.sp]
	m.registers.sp -= 1

	return addr
}
