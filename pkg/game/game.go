package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/justjoeyuk/chip8-go/pkg/chip8"
)

var keyMap = []ebiten.Key{
	ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4,
	ebiten.KeyQ, ebiten.KeyW, ebiten.KeyE, ebiten.KeyR,
	ebiten.KeyA, ebiten.KeyS, ebiten.KeyD, ebiten.KeyF,
	ebiten.KeyZ, ebiten.KeyX, ebiten.KeyC, ebiten.KeyV,
}

// Game - The primary game object for the ebiten engine
type Game struct {
	ScaleOptions  *ebiten.DrawImageOptions
	Chip8Screen   *ebiten.Image
	Chip8Emulator *chip8.Chip8
}

/*	Updates the Emulator Keyboard State and returns
	which emulator key is pressed, or nil */
func (g *Game) updateKeys() *byte {
	var keyPressed *byte = nil

	for index, key := range keyMap {
		emuKey := byte(index)

		if ebiten.IsKeyPressed(key) {
			keyPressed = &emuKey
			g.Chip8Emulator.Keyboard.PressKey(emuKey)
		} else {
			g.Chip8Emulator.Keyboard.ReleaseKey(emuKey)
		}
	}

	return keyPressed
}

// Update The Game Loop
func (g *Game) Update() error {
	keyPressed := g.updateKeys()
	shouldContinue := g.Chip8Emulator.HandleKeyPressed(keyPressed)

	if !shouldContinue {
		return nil
	}

	return nil
}

// Draw The Render Loop
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Chip8Screen, g.ScaleOptions)
}

// Layout The Screen
func (g *Game) Layout(int, int) (int, int) {
	return 640, 320
}
