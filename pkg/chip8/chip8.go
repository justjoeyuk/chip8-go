package chip8

const (
	// EmulatorWidth - Width of the Emulators Screen
	EmulatorWidth = 64

	// EmulatorHeight - Height of the Emulators Screen
	EmulatorHeight = 32
)

// Chip8 - The emulator for the CHIP8 Interpreter
type Chip8 struct {
	Memory   *Memory
	Screen   *Screen
	Keyboard *Keyboard
}

// NewChip8 - Returns an instance of the Chip8
func NewChip8(writer PixelWriter, romData []byte) *Chip8 {
	chip8 := &Chip8{}

	chip8.Memory = NewMemory(romData)
	chip8.Screen = NewScreen(writer)
	chip8.Keyboard = NewKeyboard()

	return chip8
}

func (c *Chip8) extractReferenceBits(opcode uint16) (nnn, n, x, y, kk uint16) {
	nnn = opcode & 0x0FFF
	n = opcode & 0x000F
	x = (opcode >> 8) & 0x000F
	y = (opcode >> 4) & 0x000F
	kk = opcode & 0x00FF

	return
}

// ExecOp - Emulates the given 16bit operation on the Chip8
func (c *Chip8) ExecOp(opcode uint16) {
	//nnn, n, x, y, kk := c.extractReferenceBits(opcode)
	switch opcode {
	case 0x00E0:
		c.Screen.ClearPixels()
		break
	case 0x00EE:
		c.Memory.registers.pc = c.Memory.stack[c.Memory.registers.sp]
		c.Memory.PopStack()
		break
	}
}

// HandleKeyPressed -g tracks if the keyboard is waiting for a key press. If it's waiting for a key, and no key is pressed then return false. Otherwise, return true.
func (c *Chip8) HandleKeyPressed(key *byte) bool {
	if key != nil {
		c.Keyboard.LastKeyPressed = *key
	}

	if c.Keyboard.WaitingForKeyPress {
		if key == nil {
			return false
		}

		c.Keyboard.WaitingForKeyPress = false
	}

	return true
}

// GetProgramCounter - from the registers
func (c *Chip8) GetProgramCounter() uint16 {
	return c.Memory.registers.pc
}

// IncrementProgramCounter - by two bytes
func (c *Chip8) IncrementProgramCounter() {
	c.Memory.registers.pc += 2
}
