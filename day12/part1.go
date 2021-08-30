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

	facing := EAST

	x := 0
	y := 0

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
			y += number
		case 'S':
			y -= number
		case 'E':
			x += number
		case 'W':
			x -= number
		case 'L':
			if number%90 != 0 {
				panic("rotation was not a multiple of 90")
			}
			for i := 0; i < int(number/90); i++ {
				facing -= 1
				if facing < 0 {
					facing = WEST
				}
			}
		case 'R':
			if number%90 != 0 {
				panic("rotation was not a multiple of 90")
			}
			for i := 0; i < int(number/90); i++ {
				facing += 1
				if facing > 3 {
					facing = NORTH
				}
			}
		case 'F':
			switch facing {
			case NORTH:
				y += number
			case EAST:
				x += number
			case SOUTH:
				y -= number
			case WEST:
				x -= number
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
