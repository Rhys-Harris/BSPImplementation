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
		if segment.pointWithin(other.start) {
			return Pos{}, LI_COINCIDENT
		}
		return Pos{}, LI_NONE
	}

	t := ((segment.start.x-other.start.x)*
		(other.start.y-other.end.y) -
		(segment.start.y-other.start.y)*
			(other.start.x-other.end.x)) / d

	u := -((segment.start.x-segment.end.x)*
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
		if segment.pointWithin(other.start) {
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

func (segment *Segment) pointWithin(pos Pos) bool {
	a1 := segment.start.angleTo(segment.end)
	a2 := segment.start.angleTo(pos)
	s := math.Sin(a2-a1)
	return s == 0
}

// Positive if in front, negative if behind, 0 is on top
func (segment *Segment) segmentRelation(other Segment) int {
	_, i := segment.intersectAsInfinite(other)

	if i != LI_NONE {
		return 0
	}

	if segment.pointInFront(other.start) {
		return 1
	}

	return -1
}

// Simply returns true if there was
// an intersection, rather than giving
// a `LineIntersection`
func (segment *Segment) segmentIntersect(other Segment) bool {
	_, li := segment.intersect(other)
	return li != LI_NONE
}
