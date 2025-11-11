package main

type Circle struct {
	origin Pos
	radius float64
}

func (circle *Circle) pointWithin(pos Pos) bool {
	return circle.origin.distance(pos) <= circle.radius 
}
