package chip8

const (
	EmulatorWidth = 64
	EmulatorHeight = 32
)

type Chip8 struct {
	Memory *Memory
	Screen *Screen
	Keyboard *Keyboard
}

func NewChip8(writer PixelWriter) *Chip8 {
	chip8 := &Chip8{}

	chip8.Memory = NewMemory()
	chip8.Screen = NewScreen(writer)
	chip8.Keyboard = NewKeyboard()

	return chip8
}

// HandleKeyPressed - Returns true if we are able to process the key stroke
func (c *Chip8) HandleKeyPressed(key *byte) bool {
	if key != nil {
		c.Keyboard.LastKeyPressed = *key
	}

	if c.Keyboard.WaitingForKeyPress {
		if key == nil  {
			return false
		} else {
			c.Keyboard.WaitingForKeyPress = false
		}
	}

	return true
}