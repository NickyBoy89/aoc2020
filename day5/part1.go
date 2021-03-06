package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {

  var maxId int = 0

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

    if (seatId > maxId) {
      maxId = seatId
    }

  }

  fmt.Println(maxId)

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
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
