package main

type BSPNode struct {
	// Children
	front, back *BSPNode

	// If a branch node
	splitter *Segment

	// Any coincident lines that
	// may lay on the splitter
	lines []*Segment

	// If a leaf node
	entities []*Entity
}

func (node *BSPNode) getLinesWithin(shape Shape) []*Segment {
	// Start with some capacity rather than
	// many grow calls
	marked := make([]*Segment, 0, len(node.lines))

	// Find all intersecting lines
	for i := range len(node.lines) {
		s := node.lines[i]
		if shape.segmentIntersect(*s) {
			marked = append(marked, s)
		}
	}

	return marked
}

// Finds all segments that are within the given shape
func (node *BSPNode) querySegments(shape Shape) []*Segment {
	if node.isLeaf() {
		return nil
	}

	relation := shape.segmentRelation(*node.splitter)

	switch relation {
	case 1:
		return node.front.querySegments(shape)
	case -1:
		return node.back.querySegments(shape)
	default:
		return append(
			node.getLinesWithin(shape),
			append(
				node.front.querySegments(shape),
				node.back.querySegments(shape)...,
			)...
		)
	}
}

func (node *BSPNode) queryEntities(shape Shape) []*Entity {
	if node.isLeaf() {
		// Create the list of entities to
		// give back to query
		chosen := []*Entity{}

		// Find all entities within
		for i := range len(node.entities) {
			e := node.entities[i]
			if shape.pointWithin(e.pos) {
				chosen = append(chosen, e)
			}
		}

		return chosen
	}

	relation := shape.segmentRelation(*node.splitter)

	switch relation {
	case 1:
		return node.front.queryEntities(shape)
	case -1:
		return node.back.queryEntities(shape)
	default:
		return append(
			node.front.queryEntities(shape),
			node.back.queryEntities(shape)...,
		)
	}
}

func (node *BSPNode) dump(segments *[]*Segment, nodes *[]*BSPNodeDump, parentIndex int) {
	nodeDump := &BSPNodeDump{
		parent: parentIndex,
		lines:  []int{},
	}

	for i := range len(node.lines) {
		lineIndex := len(*segments)
		*segments = append(*segments, node.lines[i])
		nodeDump.lines = append(nodeDump.lines, lineIndex)
	}

	thisIndex := len(*nodes)
	*nodes = append(*nodes, nodeDump)

	if node.isLeaf() {
		return
	}

	node.front.dump(segments, nodes, thisIndex)
	node.back.dump(segments, nodes, thisIndex)
}

func (node *BSPNode) isLeaf() bool {
	return node.front == nil && node.back == nil
}

func (node *BSPNode) addEntity(e *Entity) {
	// Leaf case
	if node.isLeaf() {
		node.entities = append(node.entities, e)
		return
	}

	// Branch case
	if node.splitter.pointInFront(e.pos) {
		node.front.addEntity(e)
	} else {
		node.back.addEntity(e)
	}
}

func (node *BSPNode) propogateChildren() {
	// Go through each entity, attempting
	// to pass it on to a child
	for i := range len(node.entities) {
		e := node.entities[i]

		if node.splitter.pointInFront(e.pos) {
			node.front.addEntity(e)
		} else {
			node.back.addEntity(e)
		}
	}

	// Clear our entities just in case
	node.entities = []*Entity{}
}

func (node *BSPNode) addLine(line *Segment) bool {
	// Leaf case
	if node.isLeaf() {
		node.splitter = line
		node.lines = append(node.lines, line)
		node.front = &BSPNode{}
		node.back = &BSPNode{}
		node.propogateChildren()
		return true
	}

	// Branch case

	// Point of intesection
	poi, intersection := node.splitter.intersectAsInfinite(*line)

	switch intersection {
	case LI_INTERSECT:
		// Intersection case
		l1 := &Segment{
			line.start,
			poi,
		}

		l2 := &Segment{
			poi,
			line.end,
		}

		var frontLine, backLine *Segment
		if node.splitter.pointInFront(line.start) {
			frontLine = l1
			backLine = l2
		} else {
			frontLine = l2
			backLine = l1
		}

		added := true
		added = added && node.front.addLine(frontLine)
		added = added && node.back.addLine(backLine)
		return added

	case LI_NONE:
		// Entirely on one side case
		if node.splitter.pointInFront(line.start) {
			return node.front.addLine(line)
		} else {
			return node.back.addLine(line)
		}

	case LI_COINCIDENT:
		// Keep a reference, but don't use
		// as a splitter
		node.lines = append(node.lines, line)
		return true

	default:
		// Invalid intersection code?
		return false
	}
}

func (node *BSPNode) nodeAtPos(pos Pos) *BSPNode {
	// Leaf case
	if node.isLeaf() {
		return node
	}

	// Branch case
	if node.splitter.pointInFront(pos) {
		return node.front.nodeAtPos(pos)
	} else {
		return node.back.nodeAtPos(pos)
	}
}
