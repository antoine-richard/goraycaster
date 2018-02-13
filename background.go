package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func renderBackground(canvas *ebiten.Image) {
	ebitenutil.DrawRect(canvas, 0, 0, screenWidth, screenHeight/2, color.RGBA{49, 162, 242, 255})
	ebitenutil.DrawRect(canvas, 0, screenHeight/2, screenWidth, screenHeight/2, color.RGBA{64, 64, 64, 255})
}
