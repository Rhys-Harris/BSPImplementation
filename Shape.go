package main

type Shape interface {
	// Whether this shaps is fully behind,
	// fully in front, or on top of the given segment
	segmentRelation(seg Segment) int

	pointWithin(pos Pos) bool

	// Does a segment intersect with any area of this shape?
	segmentIntersect(seg Segment) bool
}
