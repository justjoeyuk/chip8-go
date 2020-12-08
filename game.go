package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/justjoeyuk/chip8-go/pkg/chip8"
)

var KeyMap = []ebiten.Key{
	ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4,
	ebiten.KeyQ, ebiten.KeyW, ebiten.KeyE, ebiten.KeyR,
	ebiten.KeyA, ebiten.KeyS, ebiten.KeyD, ebiten.KeyF,
	ebiten.KeyZ, ebiten.KeyX, ebiten.KeyC, ebiten.KeyV,
}

type Game struct {
	scaleOptions  *ebiten.DrawImageOptions
	chip8Screen   *ebiten.Image
	chip8Emulator *chip8.Chip8
}

// UpdateKeys - Update the Emulator Keyboard State
// Returns the key if a key has been pressed
func (g *Game) UpdateKeys() *byte {
	var keyPressed *byte = nil

	for index, key := range KeyMap {
		emuKey := byte(index)

		if ebiten.IsKeyPressed(key) {
			keyPressed = &emuKey
			g.chip8Emulator.Keyboard.PressKey(emuKey)
		} else {
			g.chip8Emulator.Keyboard.ReleaseKey(emuKey)
		}
	}

	return keyPressed
}

func (g *Game) Update() error {
	keyPressed := g.UpdateKeys()
	shouldContinue := g.chip8Emulator.HandleKeyPressed(keyPressed)

	if !shouldContinue {
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.chip8Screen, g.scaleOptions)
}

func (g *Game) Layout(ow int, oh int) (int, int) {
	return 640, 320
}
