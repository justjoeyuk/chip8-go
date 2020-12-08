package chip8

const (
	// EmulatorWidth - Width of the Emulators Screen
	EmulatorWidth = 64

	// EmulatorWidth - Height of the Emulators Screen
	EmulatorHeight = 32
)

// Chip8 - The emulator for the CHIP8 Interpreter
type Chip8 struct {
	Memory   *Memory
	Screen   *Screen
	Keyboard *Keyboard
}

// NewChip8 - Returns an instance of the Chip8
func NewChip8(writer PixelWriter) *Chip8 {
	chip8 := &Chip8{}

	chip8.Memory = NewMemory()
	chip8.Screen = NewScreen(writer)
	chip8.Keyboard = NewKeyboard()

	return chip8
}

/*	HandleKeyPressed tracks if the keyboard is waiting for
	a key press. If it's waiting for a key, and no key is pressed
	then return false. Otherwise, return true.
*/
func (c *Chip8) HandleKeyPressed(key *byte) bool {
	if key != nil {
		c.Keyboard.LastKeyPressed = *key
	}

	if c.Keyboard.WaitingForKeyPress {
		if key == nil {
			return false
		} else {
			c.Keyboard.WaitingForKeyPress = false
		}
	}

	return true
}
