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
