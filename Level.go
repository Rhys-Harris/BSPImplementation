package main

import (
	"os"
	"strconv"
	"strings"
)

type Level struct {
	walls []*Segment
}

func (level *Level) init(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	if len(lines) == 1 {
		lines = strings.Split(string(data), "\n")
	}

	for i := range len(lines) {
		line := strings.Trim(lines[i], " \t")

		if len(line) == 0 {
			continue
		}

		args := strings.Split(line, " ")
		cmd := args[0]
		args = args[1:]

		switch cmd {
		case "SEGMENT":
			if len(args) != 4 {
				panic("Invalid number of args for SEGMENT")
			}

			x1, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				panic(err)
			}

			y1, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				panic(err)
			}

			x2, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				panic(err)
			}

			y2, err := strconv.ParseFloat(args[3], 64)
			if err != nil {
				panic(err)
			}
			
			level.walls = append(level.walls, &Segment{
				Pos{x1, y1},
				Pos{x2, y2},
			})

		case "//":
	
		default:
			panic("Invalid comman")
		}
	}
}
