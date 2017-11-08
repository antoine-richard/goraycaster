package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func renderBackground(canvas *ebiten.Image) {
	canvas.Fill(color.RGBA{49, 162, 242, 255})
}
