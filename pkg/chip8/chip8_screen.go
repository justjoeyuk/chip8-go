package chip8

import (
	"image/color"
)

type PixelWriter interface {
	Set(x, y int, color color.Color)
	At(x int, y int) color.Color
}

type Screen struct {
	writer PixelWriter
}

func NewScreen(writer PixelWriter) *Screen {
	s := &Screen{
		writer,
	}

	return s
}

func (s *Screen) SetPixelState(x, y int, enabled bool) {
	pixelColor := color.Black

	if enabled {
		existingState := s.GetPixelState(x, y)
		if !existingState {
			pixelColor = color.White
		}
	}

	x = x % EmulatorWidth
	y = y % EmulatorHeight

	s.writer.Set(x, y, pixelColor)
}

func (s *Screen) GetPixelState(x, y int) bool {
	pixelColor := s.writer.At(x, y)

	if pixelColor == color.White {
		return true
	}

	return false
}

//Draw Sprite - Draw a Sprite, wrap the sides
func (s *Screen) DrawSprite(x, y int, sprite []byte) bool {
	pixelCollision := false

	for i := 0; i < len(sprite); i++ {
		rowByte := sprite[i]

		for l := 0; l <= 8; l++ {
			if (rowByte << l) & 0b10000000 == 0 {
				continue
			}

			if !pixelCollision {
				pixelCollision = s.GetPixelState(x, y)
			}

			s.SetPixelState(x + l, y + i, true)
		}
	}

	return pixelCollision
}