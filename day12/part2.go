package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	x := 0
	y := 0

	waypointX := 10
	waypointY := 1

	for _, instruction := range strings.Split(string(input), "\n") {
		if instruction == "" { // Skip empty lines
			continue
		}

		action := instruction[0]
		number, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		switch action {
		case 'N':
			waypointY += number
		case 'S':
			waypointY -= number
		case 'E':
			waypointX += number
		case 'W':
			waypointX -= number
		case 'F':
			for i := 0; i < number; i++ {
				x += waypointX
				y += waypointY
			}
		case 'L', 'R':
			if number%90 != 0 {
				panic("rotation was not a multiple of 90")
			}

			for i := 0; i < int(number/90); i++ {
				if action == 'R' {
					// Clockwise
					waypointX, waypointY = waypointY, -waypointX
				} else {
					// Counterclockwise
					waypointX, waypointY = -waypointY, waypointX
				}
			}
		}
	}

	fmt.Println(abs(x) + abs(y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
