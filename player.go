package main

import "math"

var player = position{3, 3, 0}

func rotatePlayer(angle float64) {
	player.direction = math.Mod(player.direction+angle+circle, circle)
}
