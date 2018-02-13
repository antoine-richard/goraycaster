package main

import "math"

type player struct {
	position
	motion
}

type motion struct {
	walkFront bool
	walkBack  bool
	turnLeft  bool
	turnRight bool
}

func (p *player) walk(distance float64) {
	p.position.x += math.Cos(p.position.direction) * distance
	p.position.y += math.Sin(p.position.direction) * distance
}

func (p *player) rotate(angle float64) {
	p.position.direction = math.Mod(p.position.direction+angle+circle, circle)
}

func (p *player) update() {
	if p.motion.walkFront {
		p.walk(0.04)
	}
	if p.motion.walkBack {
		p.walk(-0.04)
	}
	if p.motion.turnRight {
		p.rotate(0.04)
	}
	if p.motion.turnLeft {
		p.rotate(-0.04)
	}
}
