package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {

  var ids []int

  seatMap := make([][]int, 127)

  for i, _ := range seatMap {
    seatMap[i] = []int{0, 0, 0, 0, 0, 0, 0, 0}
  }

  input, err := ioutil.ReadFile("input")
  check(err)

  lines := strings.Split(string(input), "\n")

  for _, i := range lines {

    if (i == "") {
      continue
    }

    row := binaryGuess(i[:8], 'F', 0, 127) // First 7 characters

    column := binaryGuess(i[len(i) - 3:], 'L', 0, 7) // Last three characters

    seatId := row * 8 + column

    seatMap[row][column] = 1

    ids = append(ids, seatId)

  }

  // fmt.Println(ids)

  for i, seat := range seatMap {
    for k, j := range seat {
      if (j == 0) {
        fmt.Println(k)
      }
    }
    fmt.Println(seat, i)
  }

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func abs(num int) int {
  if (num < 0) {
    return num * -1
  } else {
    return num
  }
}

func contains(slice []int, value int) bool {
  for _, i := range slice {
    if (value == i) {
      return true
    }
  }

  return false
}

func binaryGuess(directions string, lower rune, low, high int) int {

  for _, i := range directions[:len(directions)] {

    mid := (high - low) / 2 + 1

    if (i == lower) {
      high -= mid
    } else {
      low += mid
    }
  }

  if (rune(directions[len(directions) - 1]) == lower) {
    return low
  } else {
    return high
  }

}
