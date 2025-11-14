package main

type Rect struct {
	pos  Pos
	size Pos
}

func (rect *Rect) topLeft() Pos {
	return rect.pos
}

func (rect *Rect) topRight() Pos {
	return Pos{rect.pos.x+rect.size.x, rect.pos.y}
}

func (rect *Rect) bottomLeft() Pos {
	return Pos{rect.pos.x, rect.pos.y+rect.size.y}
}

func (rect *Rect) bottomRight() Pos {
	return rect.pos.add(rect.size)
}

func (rect *Rect) pointWithin(pos Pos) bool {
	return pos.x >= rect.pos.x && pos.x <= rect.pos.x+rect.size.x &&
		pos.y >= rect.pos.y && pos.y <= rect.pos.y+rect.size.y
}

func (rect *Rect) segmentIntersect(segment Segment) bool {
	// Find intersections between any lines
	_, i := (&Segment{rect.topLeft(), rect.topRight()}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	_, i = (&Segment{rect.topRight(), rect.bottomRight()}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	_, i = (&Segment{rect.bottomRight(), rect.bottomLeft()}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	_, i = (&Segment{rect.bottomLeft(), rect.topLeft()}).intersect(segment)
	if i != LI_NONE {
		return true
	}

	return rect.pointWithin(segment.start)
}

// Positive if in front, negative if behind, 0 is on top
func (rect *Rect) segmentRelation(segment Segment) int {
	b1 := segment.pointInFront(rect.topLeft())
	b2 := segment.pointInFront(rect.topRight())
	b3 := segment.pointInFront(rect.bottomLeft())
	b4 := segment.pointInFront(rect.bottomRight())

	if b1 && b2 && b3 && b4 {
		// Fully in front
		return 1
	} else if !b1 && !b2 && !b3 && !b4 {
		// Fully behind
		return -1
	} else {
		// Both
		return 0
	}
}
