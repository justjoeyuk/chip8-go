package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/justjoeyuk/chip8-go/pkg/chip8"
	"github.com/justjoeyuk/chip8-go/pkg/game"
)

func main() {
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("CHIP8 GO")

	// We should render the game at 64x32 and scale it up to 640x320
	scaleOptions := &ebiten.DrawImageOptions{}
	scaleOptions.GeoM.Scale(10, 10)

	chip8Screen := ebiten.NewImage(64, 32)

	g := &game.Game{
		Chip8Screen:   chip8Screen,
		ScaleOptions:  scaleOptions,
		Chip8Emulator: chip8.NewChip8(chip8Screen),
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
