package main

import "os"

type BSPTree struct {
	root *BSPNode
}

func (tree *BSPTree) init() {
	tree.root = &BSPNode{
		front:    nil,
		back:     nil,
		splitter: nil,
		entities: nil,
	}
}

func (tree *BSPTree) addLine(line *Segment) bool {
	return tree.root.addLine(line)
}

func (tree *BSPTree) entitiesNearby(pos Pos) []*Entity {
	// Find the node that encompasses this area
	node := tree.root.nodeAtPos(pos)

	// No node found (should practically never happen),
	// simply give back no entities
	if node == nil {
		return nil
	}

	return node.entities
}

// Uses generic shape so that anything
// can be used for the query
func (tree *BSPTree) querySegments(shape Shape) []*Segment {
	return tree.root.querySegments(shape)
}

func (tree *BSPTree) queryEntities(shape Shape) []*Entity {
	return tree.root.queryEntities(shape)
}

func (tree *BSPTree) queryEntitiesByCircle(circle Circle) []*Entity {
	return tree.root.queryEntitiesByCircle(circle)
}

func (tree *BSPTree) queryEntitiesByHitscan(scan Segment) []*Entity {
	return tree.root.queryEntitiesByHitscan(scan)
}

func (tree *BSPTree) querySegmentsByHitscan(scan Segment) []*Segment {
	return tree.root.querySegmentsByHitscan(scan)
}

// Dumps the BSP tree to a file
func (tree *BSPTree) dump(fileName string) {
	// Segments that are referenced by the nodes
	segments := []*Segment{}
	nodes := []*BSPNodeDump{}

	tree.root.dump(&segments, &nodes, -1)

	out := ""

	for i := range len(segments) {
		out += segments[i].String() + "\n"
	}

	for i := range len(nodes) {
		out += nodes[i].String() + "\n"
	}

	os.WriteFile(fileName, []byte(out), 0664)
}

func (tree *BSPTree) addEntity(e *Entity) {
	tree.root.addEntity(e)
}
