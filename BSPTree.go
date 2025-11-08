package main

type BSPTree struct {
	root *BSPNode
}

func (tree *BSPTree) init() {
	tree.root = &BSPNode{
		front: nil,
		back: nil,
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
