package main

type Triangle struct {
	p1, p2, p3 Pos
}

func (triangle *Triangle) pointWithin(pos Pos) bool {
	return (
		(&Segment{triangle.p1, triangle.p2}).pointInFront(pos) &&
		(&Segment{triangle.p2, triangle.p3}).pointInFront(pos) &&
		(&Segment{triangle.p3, triangle.p1}).pointInFront(pos))
}

func (triangle *Triangle) segmentIntersect(segment Segment) bool {
	// Find intersections between any lines
	_, i := (&Segment{triangle.p1, triangle.p2}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	_, i = (&Segment{triangle.p2, triangle.p3}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	_, i = (&Segment{triangle.p3, triangle.p1}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	return triangle.pointWithin(segment.start)
}

func (triangle *Triangle) segmentRelation(segment Segment) int {
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
