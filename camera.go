package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	resolution  = 80
	focalLength = .8
	// columnSpacing         = screenWidth / resolution
	viewRange float64 = 8
)

type cell struct {
	x        float64
	y        float64
	distance float64
	height   int
}

func renderColumns(canvas *ebiten.Image) {
	for column := 0; column < resolution; column++ {
		x := float64(column)/float64(resolution) - .5
		angle := math.Atan2(x, focalLength)
		cellsHit := castRay(player, player.direction+angle)
		renderColumn(canvas, column, cellsHit)
	}
}

func renderColumn(canvas *ebiten.Image, column int, cellsHit []cell) {
	//TODO
}

func castRay(player position, rayAngle float64) []cell {

	var sin = math.Sin(rayAngle)
	var cos = math.Cos(rayAngle)

	var rayToClosestCell func(cell) []cell
	rayToClosestCell = func(origin cell) []cell {

		stepX := step(sin, cos, float64(origin.x), float64(origin.y), computeStepX)
		stepY := step(cos, sin, float64(origin.y), float64(origin.x), computeStepY)

		var nextStep cell
		if stepX.distance < stepY.distance {
			nextStep = inspect(stepX, 1, 0, origin.distance, sin, cos)
		} else {
			nextStep = inspect(stepY, 0, 1, origin.distance, sin, cos)
		}

		if nextStep.distance > viewRange {
			return []cell{origin}
		}
		return append([]cell{origin}, rayToClosestCell(nextStep)...)
	}

	return rayToClosestCell(cell{x: player.x, y: player.y, distance: 0})
}

func computeStepX(x float64, y float64, distance float64) cell {
	return cell{x: x, y: y, distance: distance}
}
func computeStepY(x float64, y float64, distance float64) cell {
	return cell{x: y, y: x, distance: distance}
}

func step(rise float64, run float64, x float64, y float64, computeStep func(float64, float64, float64) cell) cell {
	if run == 0 {
		return cell{distance: math.Inf(1)}
	}

	var dx float64
	if run > 0 {
		dx = math.Floor(x+1) - x
	} else {
		dx = math.Ceil(x-1) - x
	}

	dy := dx * (rise / run)

	return computeStep(x+dx, y+dy, dx*dx+dy*dy)
}

func inspect(step cell, shiftX float64, shiftY float64, distance float64, sin float64, cos float64) cell {
	var dx float64
	if cos < 0 {
		dx = shiftX
	}
	var dy float64
	if sin < 0 {
		dy = shiftY
	}
	step.height = cellHeight(step.x-dx, step.y-dy)
	step.distance = distance + math.Sqrt(step.distance)
	return step
}
