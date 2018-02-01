package main

import "math"

const (
	mapSize = 10
)

var wallGrid = [100]int{
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

func cellHeight(x float64, y float64) int {
	cellX := int(math.Floor(x))
	cellY := int(math.Floor(y))
	if cellX < 0 || cellX > mapSize-1 || cellY < 0 || cellY > mapSize-1 {
		return -1
	}
	return wallGrid[cellY*mapSize+cellX]
}
