package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	currentTimestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	_ = currentTimestamp
	schedule := lines[1]

	busLanes := strings.Split(schedule, ",")

	firstLaneNumber, _ := strconv.Atoi(busLanes[0])

	//if ticks % 7 == 0 && ticks % 13 == 1 && ticks % 59 == 4

	ticks := 0
	for ; ; ticks += firstLaneNumber {
		for ind, lane := range busLanes {
			if lane == "x" {
				continue
			}
			laneNumber, _ := strconv.Atoi(lane)
			if (ticks+ind)%laneNumber != 0 {
				break
			}
			if ind == len(busLanes)-1 {
				panic(ticks)
			}
		}
	}
}
