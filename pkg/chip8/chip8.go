package chip8

const (
	EmulatorWidth = 64
	EmulatorHeight = 32
)

type Chip8 struct {
	Memory *Memory
	Screen *Screen
}

func NewChip8(writer PixelWriter) *Chip8 {
	chip8 := &Chip8{}

	chip8.Memory = NewMemory()
	chip8.Screen = NewScreen(writer)

	return chip8
}