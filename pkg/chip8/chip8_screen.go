package chip8

import (
	"image/color"
)

// PixelWriter - An interface able to Set pixels and fetch pixels At a location
type PixelWriter interface {
	// Set a pixel at a given coordinate (x, y), a given color
	Set(x, y int, color color.Color)

	// At - Returns the color of the pixel at a given coordinate (x, y)
	At(x int, y int) color.Color
}

// Screen - Interpretation of the CHIP8 Screen
type Screen struct {
	writer PixelWriter
}

// NewScreen - Returns an instance of the Screen
func NewScreen(writer PixelWriter) *Screen {
	s := &Screen{
		writer,
	}

	return s
}

// EnablePixel - XOR a pixel onto the PixelWriter at coordinates (x,y) and ensure it wraps if out of bounds
func (s *Screen) EnablePixel(x, y int) {
	pixelColor := color.Black

	existingState := s.GetPixelState(x, y)
	if !existingState {
		pixelColor = color.White
	}

	x = x % EmulatorWidth
	y = y % EmulatorHeight

	s.writer.Set(x, y, pixelColor)
}

// DisablePixel - Disable a pixel on the PixelWriter at coordinates (x, y)
func (s *Screen) DisablePixel(x, y int) {
	s.writer.Set(x, y, color.Black)
}

// GetPixelState - Returns the state of a pixel on the PixelWriter (on/off)
func (s *Screen) GetPixelState(x, y int) bool {
	pixelColor := s.writer.At(x, y)

	if pixelColor == color.White {
		return true
	}

	return false
}

//DrawSprite - Draw a Sprite to the PixelWriter and return if there was a collision with pixels
func (s *Screen) DrawSprite(x, y int, sprite []byte) bool {
	pixelCollision := false

	for i := 0; i < len(sprite); i++ {
		rowByte := sprite[i]

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
