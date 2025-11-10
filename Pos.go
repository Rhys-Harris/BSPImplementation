package main

import "math"

type Pos struct {
	x, y float64
}

func (pos Pos) add(other Pos) Pos {
	return Pos{
		pos.x + other.x,
		pos.y + other.y,
	}
}

func (pos Pos) scale(scale float64) Pos {
	return Pos{
		pos.x * scale,
		pos.y * scale,
	}
}

func (pos *Pos) angleTo(other Pos) float64 {
	dx := other.x - pos.x
	dy := other.y - pos.y
	dis := math.Sqrt(dx*dx + dy*dy)

	// Avoid division by 0
	if dis == 0 {
		return 0
	}

	a := math.Acos(dx/dis)
	if dy < 0 {
		a = -a
	}
	return a
}
