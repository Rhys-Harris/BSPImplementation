package main

import "math"

type Circle struct {
	origin Pos
	radius float64
}

func (circle *Circle) pointWithin(pos Pos) bool {
	return circle.origin.distance(pos) <= circle.radius 
}

// Positive if in front, negative if behind, 0 is on top
func (circle *Circle) segmentRelation(segment Segment) int {
	// Angle start -> origin
	a1 := segment.start.angleTo(circle.origin)

	// Angle start -> end
	a2 := segment.start.angleTo(segment.end)

	// Angle between 2 angles
	a := a2 - a1

	// Distance from start to origin
	sdis := segment.start.distance(circle.origin)

	// Distance from closest point on line to origin
	dis := sdis * math.Cos(a)

	// Intersection?
	if dis <= circle.radius {
		return 0
	}

	// Fully on one side case
	if segment.pointInFront(circle.origin) {
		return 1
	} else {
		return -1
	}
}

func (circle *Circle) segmentIntersect(segment Segment) bool {
	// Angle start -> origin
	a1 := segment.start.angleTo(circle.origin)

	// Angle start -> end
	a2 := segment.start.angleTo(segment.end)

	// Angle between 2 angles
	a := a2 - a1

	// Distance from start to origin
	sdis := segment.start.distance(circle.origin)

	// Distance from closest point on line to origin
	dis := sdis * math.Cos(a)

	// Intersection?
	return dis <= circle.radius
}
