package chip8

const PROGRAM_START_ADDRESS = 0x200
const HEX_SPRITE_ADDRESS = 0x00

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

func NewMemory() *Memory {
	m := &Memory{}

	copy(m.ram[HEX_SPRITE_ADDRESS:], defaultCharacterSet())

	return m
}

func (m *Memory) Set(loc int, val byte) {
	m.ram[loc] = val
}

func (m *Memory) Get(loc int) byte {
	return m.ram[loc]
}

func (m *Memory) GetTwoBytes(loc int) uint16 {
	b1 := uint16(m.Get(loc))
	b2 := uint16(m.Get(loc + 1))

	return b1<<8 | b2
}

func (m *Memory) GetNBytes(loc, numBytes int) []byte {
	bytes := make([]byte, numBytes)

	for i := 0; i < numBytes; i++ {
		bytes[i] = m.ram[loc+i]
	}

	return bytes
}

func (m *Memory) PushStack(val uint16) {
	m.stack[m.sp+1] = val
	m.sp += 1
}

func (m *Memory) PopStack() uint16 {
	addr := m.stack[m.sp]
	m.sp -= 1

	return addr
}
