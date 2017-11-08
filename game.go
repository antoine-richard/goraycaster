package main

import "math"

type position struct {
	x         float64
	y         float64
	direction float64
}

const (
	mapSize = 10
)

var (
	player   = position{3, 3, 0}
	wallGrid = [100]int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 1, 1, 0, 0, 0, 1,
		1, 0, 0, 0, 1, 1, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

func cellHeight(x float64, y float64) int {
	cellX := int(math.Floor(x))
	cellY := int(math.Floor(y))
	if cellX < 0 || cellX > mapSize-1 || cellY < 0 || cellY > mapSize-1 {
		return -1
	}
	return wallGrid[cellY*mapSize+cellX]
}
