package chip8

// Keyboard - Interpretation of the CHIP8 Keyboard
type Keyboard struct {
	keys [16]bool

	// LastKeyPressed on the Keyboard
	LastKeyPressed byte

	// LockKeypressToRegister if -1, game loop will be locked until key press. The key is then stored in register x
	LockKeypressToRegister int
}

// NewKeyboard - Returns and instance of Keyboard
func NewKeyboard() *Keyboard {
	return &Keyboard{LockKeypressToRegister: -1}
}

// IsKeyDown - returns if a given key is pressed or not
func (k *Keyboard) IsKeyDown(key uint8) bool {
	return k.keys[key]
}

// PressKey - set the given key to true (pressed)
func (k *Keyboard) PressKey(key uint8) {
	k.keys[key] = true
}

// ReleaseKey - set the given key to false (released)
func (k *Keyboard) ReleaseKey(key uint8) {
	k.keys[key] = false
}
