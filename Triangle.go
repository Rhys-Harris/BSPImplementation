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
