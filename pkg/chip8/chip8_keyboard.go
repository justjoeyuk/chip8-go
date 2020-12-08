package chip8

type Keyboard struct {
	keys [16]bool

	// LastKeyPressed on the Keyboard
	LastKeyPressed byte

	// WaitingForKeyPress if enabled, game loop will be locked until key press
	WaitingForKeyPress bool
}

// NewKeyboard - Returns and instance of Keyboard
func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

// IsKeyDown returns if a given key is pressed or not
func (k *Keyboard) IsKeyDown(key byte) bool {
	return k.keys[key]
}

// IsKeyDown set the given key to true (pressed)
func (k *Keyboard) PressKey(key byte) {
	k.keys[key] = true
}

// IsKeyDown set the given key to false (released)
func (k *Keyboard) ReleaseKey(key byte) {
	k.keys[key] = false
}
