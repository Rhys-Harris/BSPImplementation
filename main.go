package main

import "fmt"

func main() {
	// Create a new BSP tree
	bsp := &BSPTree{}
	bsp.init()

	fmt.Println("Attempting to add first line")
	fmt.Println(bsp.addLine(&Segment{
		Pos{0, -100},
		Pos{0, 100},
	}))

	fmt.Println("Testing the 'line in front' case")
	fmt.Println(bsp.addLine(&Segment{
		Pos{10, -100},
		Pos{20, 100},
	}))

	fmt.Println("Testing the 'line behind' case")
	fmt.Println(bsp.addLine(&Segment{
		Pos{-10, -100},
		Pos{-20, 100},
	}))

	fmt.Println("Testing the 'line split' case")
	fmt.Println(bsp.addLine(&Segment{
		Pos{10, 0},
		Pos{100, 0},
	}))

	fmt.Println("Testing the 'line coincident' case")
	fmt.Println(bsp.addLine(&Segment{
		Pos{0, 200},
		Pos{0, 300},
	}))
}
