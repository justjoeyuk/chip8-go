package chip8

// Screen - Interpretation of the CHIP8 Screen
type Screen struct {
	Pixels [32][64]bool
}

// NewScreen - Returns an instance of the Screen
func NewScreen() *Screen {
	s := &Screen{}
	return s
}

// EnablePixel - XOR a pixel at coordinates (x,y) and ensure it wraps if out of bounds
func (s *Screen) EnablePixel(x, y int) {
	x = x % EmulatorWidth
	y = y % EmulatorHeight

	s.Pixels[y][x] = !s.GetPixelState(x, y)
}

// DisablePixel - Disable a pixel at coordinates (x, y)
func (s *Screen) DisablePixel(x, y int) {
	s.Pixels[y][x] = false
}

// ClearPixels - Disable all pixels
func (s *Screen) ClearPixels() {
	s.Pixels = [32][64]bool{}
}

// GetPixelState - Returns the state of a pixel (on/off)
func (s *Screen) GetPixelState(x, y int) bool {
	return s.Pixels[y][x]
}

//DrawSprite - Draw a Sprite and return if there was a collision with pixels
func (s *Screen) DrawSprite(x, y int, sprite *[]uint8) bool {
	pixelCollision := false
	spriteData := *sprite

	for i := 0; i < len(spriteData); i++ {
		rowByte := spriteData[i]

		for l := 0; l <= 8; l++ {
			if (rowByte<<l)&0b10000000 == 0 {
				continue
			}

			if !pixelCollision {
				pixelCollision = s.GetPixelState(x, y)
			}

			s.EnablePixel(x+l, y+i)
		}
	}

	return pixelCollision
}
