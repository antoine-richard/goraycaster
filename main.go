package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	canvas *ebiten.Image
)

func update() {}

func render(canvas *ebiten.Image) {
	renderBackground(canvas)
}

func init() {
	canvas, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
}

func loop(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	update()
	render(canvas)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(canvas, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%2.f FPS", ebiten.CurrentFPS()))
	return nil
}

func main() {
	if err := ebiten.Run(loop, screenWidth, screenHeight, 2, "Go Raycaster"); err != nil {
		log.Fatal(err)
	}
}
