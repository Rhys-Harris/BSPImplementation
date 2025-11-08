package main

import "fmt"

func main() {
	// Load in level data
	level := &Level{}
	level.init("input.level")

	// Create a new BSP tree
	bsp := &BSPTree{}
	bsp.init()

	fmt.Println("Attempting to add first line")
	fmt.Println(bsp.addLine(level.walls[0]))

	fmt.Println("Testing the 'line in front' case")
	fmt.Println(bsp.addLine(level.walls[1]))

	fmt.Println("Testing the 'line behind' case")
	fmt.Println(bsp.addLine(level.walls[2]))

	fmt.Println("Testing the 'line split' case")
	fmt.Println(bsp.addLine(level.walls[3]))

	fmt.Println("Testing the 'line coincident' case")
	fmt.Println(bsp.addLine(level.walls[4]))

	bsp.dump("output.rbsp")
}
