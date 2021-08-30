package main

import (
	"fmt"
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
	schedule := lines[1]

	busLanes := strings.Split(schedule, ",")
	fmt.Println(currentTimestamp, schedule, busLanes)

	ticks := 0
	for ; ; ticks++ {
		for _, lane := range busLanes {
			if lane != "x" {
				laneNumber, err := strconv.Atoi(lane)
				if err != nil {
					panic(err)
				}
				if ticks%laneNumber != 0 {
					fmt.Printf("[%v]\t", strings.Repeat(" ", len(strconv.Itoa(ticks))))
				} else {
					fmt.Printf("[%v]\t", ticks)
					if ticks >= currentTimestamp {
						panic(laneNumber * (ticks - currentTimestamp))
					}
				}
			}
		}
		fmt.Println("")
	}
}
