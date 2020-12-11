package chip8

import (
	"math/rand"
	"time"
)

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

func boolToByte(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// NewChip8 - Returns an instance of the Chip8
func NewChip8(romData []uint8) *Chip8 {
	chip8 := &Chip8{}

	chip8.Memory = NewMemory(romData)
	chip8.Screen = NewScreen()
	chip8.Keyboard = NewKeyboard()

	return chip8
}

func (c *Chip8) extractReferenceBits(opcode uint16) (nnn uint16, n, x, y, kk uint8) {
	nnn = opcode & 0x0FFF
	n = uint8(opcode & 0x000F)
	x = uint8((opcode >> 8) & 0x000F)
	y = uint8((opcode >> 4) & 0x000F)
	kk = uint8(opcode & 0x00FF)

	return
}

func (c *Chip8) ExecOpExtended(opcode uint16) {
	nnn, n, x, y, kk := c.extractReferenceBits(opcode)

	switch opcode & 0xF000 {
	case 0x1000:
		c.Memory.registers.pc = nnn
	case 0x2000:
		c.Memory.PushStack(c.Memory.registers.pc)
		c.Memory.registers.pc = nnn
	case 0x3000:
		if c.Memory.registers.v[x] == kk {
			c.Memory.registers.pc += 4
		}
	case 0x4000:
		if c.Memory.registers.v[x] != kk {
			c.Memory.registers.pc += 4
		}
	case 0x5000:
		if c.Memory.registers.v[x] == c.Memory.registers.v[y] {
			c.Memory.registers.pc += 4
		}
	case 0x6000:
		c.Memory.registers.v[x] = kk
	case 0x7000:
		c.Memory.registers.v[x] += kk
	case 0x8000:
		switch n {
		case 0:
			c.Memory.registers.v[x] = c.Memory.registers.v[y]
		case 1:
			c.Memory.registers.v[x] = c.Memory.registers.v[x] | c.Memory.registers.v[y]
		case 2:
			c.Memory.registers.v[x] = c.Memory.registers.v[x] & c.Memory.registers.v[y]
		case 3:
			c.Memory.registers.v[x] = c.Memory.registers.v[x] ^ c.Memory.registers.v[y]
		case 4:
			tmp := c.Memory.registers.v[x] + c.Memory.registers.v[y]
			c.Memory.registers.v[x] = tmp & 0xFF
			c.Memory.registers.v[0xf] = boolToByte(tmp > 255)
		case 5:
			c.Memory.registers.v[0xf] = boolToByte(c.Memory.registers.v[x] > c.Memory.registers.v[y])
			c.Memory.registers.v[x] = c.Memory.registers.v[x] - c.Memory.registers.v[y]
		case 6:
			c.Memory.registers.v[0x0F] = c.Memory.registers.v[x] & 0x01
			c.Memory.registers.v[x] = c.Memory.registers.v[x] >> 1
		case 7:
			c.Memory.registers.v[0xf] = boolToByte(c.Memory.registers.v[y] > c.Memory.registers.v[x])
			c.Memory.registers.v[x] = c.Memory.registers.v[y] - c.Memory.registers.v[x]
		case 0xE:
			c.Memory.registers.v[0x0F] = (c.Memory.registers.v[x] >> 7) & 0x01
			c.Memory.registers.v[x] = c.Memory.registers.v[x] << 1
		}
	case 0x9000:
		if c.Memory.registers.v[x] != c.Memory.registers.v[y] {
			c.Memory.registers.pc += 4
		}
	case 0xA000:
		c.Memory.registers.i = nnn
	case 0xB000:
		c.Memory.registers.pc = nnn + uint16(c.Memory.registers.v[0])
	case 0xC000:
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		c.Memory.registers.v[x] = uint8(r1.Intn(256)) & kk
	case 0xD000:
		sprite := c.Memory.ram[c.Memory.registers.i : c.Memory.registers.i+uint16(n)]
		c.Memory.registers.v[0x0F] = boolToByte(c.Screen.DrawSprite(int(c.Memory.registers.v[x]), int(c.Memory.registers.v[y]), &sprite))
	case 0xE000:
		switch kk {
		case 0x9E:
			if c.Keyboard.IsKeyDown(c.Memory.registers.v[x]) {
				c.Memory.registers.pc += 4
			}
		case 0xA1:
			if !c.Keyboard.IsKeyDown(c.Memory.registers.v[x]) {
				c.Memory.registers.pc += 4
			}
		}
	case 0xF000:
		switch kk {
		case 0x07:
			c.Memory.registers.v[x] = c.Memory.registers.dt
		case 0x0A:
			c.Keyboard.LockKeypressToRegister = int(x)
		case 0x15:
			c.Memory.registers.dt = c.Memory.registers.v[x]
		case 0x18:
			c.Memory.registers.st = c.Memory.registers.v[x]
		case 0x1E:
			c.Memory.registers.i += uint16(c.Memory.registers.v[x])
		case 0x29:
			c.Memory.registers.i = uint16(c.Memory.registers.v[x]) * 5
		case 0x33:
			c.Memory.ram[c.Memory.registers.i] = c.Memory.registers.v[x] / 100
			c.Memory.ram[c.Memory.registers.i+1] = c.Memory.registers.v[x] / 10 % 10
			c.Memory.ram[c.Memory.registers.i+2] = c.Memory.registers.v[x] % 10
		case 0x55:
			var i uint16
			for i = 0; i <= uint16(x); i++ {
				c.Memory.ram[c.Memory.registers.i+i] = c.Memory.registers.v[i]
			}
		case 0x65:
			var i uint16
			for i = 0; i <= uint16(x); i++ {
				c.Memory.registers.v[i] = c.Memory.ram[c.Memory.registers.i+i]
			}
		}
	}
}

// ExecOp - Emulates the given 16bit operation on the Chip8
func (c *Chip8) ExecOp(opcode uint16) {
	switch opcode {
	case 0x00E0:
		c.Screen.ClearPixels()
		break
	case 0x00EE:
		c.Memory.registers.pc = c.Memory.PopStack() + 2
		break
	default:
		c.ExecOpExtended(opcode)
	}
}

// HandleKeyPressed -g tracks if the keyboard is waiting for a key press. If it's waiting for a key, and no key is pressed then return false. Otherwise, return true.
func (c *Chip8) HandleKeyPressed(key *uint8) bool {
	if c.Keyboard.LockKeypressToRegister != -1 {
		if key == nil {
			return false
		}

		c.Memory.registers.v[c.Keyboard.LockKeypressToRegister] = *key
		c.Keyboard.LockKeypressToRegister = -1
	}

	return true
}

// DelayTimer - Returns the value of the delay timer
func (c *Chip8) DelayTimer() uint8 {
	return c.Memory.registers.dt
}

// SetDelayTimer - Sets the value of the delay timer
func (c *Chip8) SetDelayTimer(val uint8) {
	c.Memory.registers.dt = val
}

// SoundTimer - Returns the value of the sound timer
func (c *Chip8) SoundTimer() uint8 {
	return c.Memory.registers.st
}

// SetSoundTimer - Sets the value of the delay timer
func (c *Chip8) SetSoundTimer(val uint8) {
	c.Memory.registers.st = val
}

// GetProgramCounter - from the registers
func (c *Chip8) GetProgramCounter() uint16 {
	return c.Memory.registers.pc
}

// IncrementProgramCounter - by two bytes
func (c *Chip8) IncrementProgramCounter() {
	c.Memory.registers.pc += 2
}
