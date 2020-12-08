package chip8

type Keyboard struct {
	keys [16]bool

	LastKeyPressed byte
	WaitingForKeyPress bool
}

func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) IsKeyDown(key byte) bool {
	return k.keys[key]
}

func (k *Keyboard) PressKey(key byte) {
	k.keys[key] = true
}

func (k *Keyboard) ReleaseKey(key byte) {
	k.keys[key] = false
}