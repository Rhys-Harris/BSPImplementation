package main

import (
	"fmt"
	"math"
)

type Segment struct {
	start, end Pos
}

func (segment *Segment) String() string {
	return fmt.Sprintf(
		"SEGMENT %f %f %f %f",
		segment.start.x,
		segment.start.y,
		segment.end.x,
		segment.end.y,
	)
}

func (segment *Segment) intersect(other Segment) (Pos, LineIntersection) {
	d := ((segment.start.x-segment.end.x)*
		(other.start.y-other.end.y) -
		(segment.start.y-segment.end.y)*
			(other.start.x-other.end.x))

	// Parallel
	if d == 0 {
		if segment.pointOnLine(other.start) {
			return Pos{}, LI_COINCIDENT
		}
		return Pos{}, LI_NONE
	}

	t := ((segment.start.x-other.start.x)*
		(other.start.y-other.end.y) -
		(segment.start.y-other.start.y)*
			(other.start.x-other.end.x)) / d

	u := ((segment.start.x-segment.end.x)*
		(segment.start.y-other.start.y) -
		(segment.start.y-segment.end.y)*
			(segment.start.x-other.start.x)) / d

	hasIntersected := 0 <= t && t <= 1 && 0 <= u && u <= 1
	if !hasIntersected {
		return Pos{}, LI_NONE
	}

	intersection := Pos{
		segment.start.x + t*(segment.end.x-segment.start.x),
		segment.start.y + t*(segment.end.y-segment.start.y),
	}

	return intersection, LI_INTERSECT
}

func (segment *Segment) intersectAsInfinite(other Segment) (Pos, LineIntersection) {
	d := ((segment.start.x-segment.end.x)*
		(other.start.y-other.end.y) -
		(segment.start.y-segment.end.y)*
			(other.start.x-other.end.x))

	// Parallel
	if d == 0 {
		if segment.pointOnLine(other.start) {
			return Pos{}, LI_COINCIDENT
		}
		return Pos{}, LI_NONE
	}

	u := - ((segment.start.x-segment.end.x)*
		(segment.start.y-other.start.y) -
		(segment.start.y-segment.end.y)*
			(segment.start.x-other.start.x))/d

	hasIntersected := 0 <= u && u <= 1
	if !hasIntersected {
		return Pos{}, LI_NONE
	}

	intersection := Pos{
		other.start.x + u*(other.end.x-other.start.x),
		other.start.y + u*(other.end.y-other.start.y),
	}

	return intersection, LI_INTERSECT
}

func (segment *Segment) pointInFront(pos Pos) bool {
	a1 := segment.start.angleTo(segment.end)
	a2 := segment.start.angleTo(pos)
	s := math.Sin(a2-a1)
	return s >= 0
}

func (segment *Segment) pointOnLine(pos Pos) bool {
	a1 := segment.start.angleTo(segment.end)
	a2 := segment.start.angleTo(pos)
	s := math.Sin(a2-a1)
	return s == 0
}

// Positive if in front, negative if behind, 0 is on top
func (segment *Segment) triangleRelation(triangle Triangle) int {
	b1 := segment.pointInFront(triangle.p1)
	b2 := segment.pointInFront(triangle.p2)
	b3 := segment.pointInFront(triangle.p3)

	if b1 && b2 && b3 {
		// Fully in front
		return 1
	} else if !b1 && !b2 && !b3 {
		// Fully behind
		return -1
	} else {
		// Both
		return 0
	}
}

// Positive if in front, negative if behind, 0 is on top
func (segment *Segment) circleRelation(circle Circle) int {
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
