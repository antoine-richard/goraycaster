package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	background *ebiten.Image
)

func init() {
	background, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
	background.Fill(color.RGBA{49, 162, 242, 255})
}

func renderBackground(canvas *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	canvas.DrawImage(background, op)
}
