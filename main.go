package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	canvas *ebiten.Image
	p      player
)

func init() {
	p = player{position{3, 3, 0}, motion{}}
	canvas, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
}

func main() {
	if err := ebiten.Run(loop, screenWidth, screenHeight, 2, "Go Raycaster"); err != nil {
		log.Fatal(err)
	}
}

func loop(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	input()
	update()
	render(screen)

	return nil
}

func input() {
	p.motion.walkFront = ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)
	p.motion.walkBack = ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)
	p.motion.turnLeft = ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)
	p.motion.turnRight = ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)
}

func update() {
	p.update()
}

func render(screen *ebiten.Image) {
	canvas.Fill(color.Black)

	renderBackground(canvas)
	renderColumns(canvas, p.position)
	renderFps(canvas)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(canvas, op)
}
