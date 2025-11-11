package main

import (
	"fmt"
	"math"
)

func main() {
	// Load in level data
	level := &Level{}
	level.init("input.level")

	// Create a new BSP tree
	bsp := &BSPTree{}
	bsp.init()

	fmt.Println("Attempting to add lines")
	for i := range len(level.walls) {
		fmt.Println(bsp.addLine(level.walls[i]))
	}

	bsp.dump("output.rbsp")

	// Get some entities?
	camera := Camera{
		pos:     Pos{0, 0},
		angle:   0,
		viewDis: 200,
		fov:     math.Pi/2, // 90deg
	}

	bsp.addEntity(&Entity{
		name: "John",
		pos:  Pos{50, 0},
	})

	bsp.addEntity(&Entity{
		name: "Mark",
		pos:  Pos{0, 50},
	})

	bsp.addEntity(&Entity{
		name: "Luke",
		pos:  Pos{-50, 0},
	})

	bsp.addEntity(&Entity{
		name: "Matthew",
		pos:  Pos{0, -50},
	})

	entities := camera.getEntitiesInView(bsp)
	fmt.Println("Found these entities from triangle")
	for i := range len(entities) {
		fmt.Println(entities[i].name)
	}
	fmt.Println()

	walls := camera.getWallsInView(bsp)
	fmt.Println("Found these walls from triangle")
	for i := range len(walls) {
		fmt.Println(walls[i])
	}
	fmt.Println()

	entities = bsp.queryEntitiesByCircle(Circle{
		Pos{0, 0},
		100,
	})	
	fmt.Println("Found these entities from circle")
	for i := range len(entities) {
		fmt.Println(entities[i].name)
	}
	fmt.Println()
}
