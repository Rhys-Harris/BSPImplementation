package main

import "math"

type Camera struct {
	pos     Pos
	angle   float64
	viewDis float64
	fov     float64
}

// Generates a bounding triangle that
// should encapsulate everything that
// the camera can see
func (camera *Camera) createViewFrustum() Triangle {
	return Triangle{
		p1: camera.pos,
		p2: camera.pos.add(Pos{
			math.Cos(camera.angle-camera.fov/2),
			math.Sin(camera.angle-camera.fov/2),
		}.scale(camera.viewDis)),
		p3: camera.pos.add(Pos{
			math.Cos(camera.angle+camera.fov/2),
			math.Sin(camera.angle+camera.fov/2),
		}.scale(camera.viewDis)),
	}
}

// Queries the BSP tree using the
// view frustum for all viewable entities
func (camera *Camera) getEntitiesInView(world *BSPTree) []*Entity {
	return world.queryEntitiesByTriangle(camera.createViewFrustum())
}


// Queries the BSP tree using the
// view frustum for all viewable walls
func (camera *Camera) getWallsInView(world *BSPTree) []*Segment {
	return camera.backfaceCull(
		world.querySegmentsByTriangle(
			camera.createViewFrustum(),
			),
		)
}

// Deletes any walls that aren't
// facing the player
func (camera *Camera) backfaceCull(walls []*Segment) []*Segment {
	facing := make([]*Segment, 0, len(walls))

	for i := range len(walls) {
		w := walls[i]
		if w.pointInFront(camera.pos) {
			facing = append(facing, w)
		}
	}

	return facing
}
