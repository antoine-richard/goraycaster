package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type position struct {
	x         float64
	y         float64
	direction float64
}

func renderFps(canvas *ebiten.Image) {
	ebitenutil.DebugPrint(canvas, fmt.Sprintf("%2.f FPS", ebiten.CurrentFPS()))
}
