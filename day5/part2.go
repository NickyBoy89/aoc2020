package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
)

func main() {

  var ids []int

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

    ids = append(ids, seatId)

  }

  sort.Ints(ids)

  fmt.Println(ids)

  for _, i := range ids {
    if (!containedIn(ids, i + 1) && containedIn(ids, i + 2)) {
      fmt.Println(i + 1)
    }
  }

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func containedIn(array []int, element int) bool {
  if (element > len(array)) {
    return false
  }
  if (array[sort.SearchInts(array, element)] == element) { // Binary search
    return true
  } else {
    return false
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
