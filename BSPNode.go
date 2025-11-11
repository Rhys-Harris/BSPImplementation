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

func (node *BSPNode) queryEntitiesByTriangle(triangle Triangle) []*Entity {
	if node.isLeaf() {
		// Create the list of entities to
		// give back to query
		chosen := []*Entity{}

		// Find all entities within
		for i := range len(node.entities) {
			e := node.entities[i]
			if triangle.pointWithin(e.pos) {
				chosen = append(chosen, e)
			}
		}

		return chosen
	}

	relation := node.splitter.triangleRelation(triangle)

	switch relation {
	case 1:
		return node.front.queryEntitiesByTriangle(triangle)
	case -1:
		return node.back.queryEntitiesByTriangle(triangle)
	default:
		return append(
			node.front.queryEntitiesByTriangle(triangle),
			node.back.queryEntitiesByTriangle(triangle)...,
		)
	}
}

func (node *BSPNode) queryEntitiesByCircle(circle Circle) []*Entity {
	if node.isLeaf() {
		// Create the list of entities to
		// give back to query
		chosen := []*Entity{}

		// Find all entities within
		for i := range len(node.entities) {
			e := node.entities[i]
			if circle.pointWithin(e.pos) {
				chosen = append(chosen, e)
			}
		}

		return chosen
	}

	relation := node.splitter.circleRelation(circle)

	switch relation {
	case 1:
		return node.front.queryEntitiesByCircle(circle)
	case -1:
		return node.back.queryEntitiesByCircle(circle)
	default:
		return append(
			node.front.queryEntitiesByCircle(circle),
			node.back.queryEntitiesByCircle(circle)...,
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
