package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {

  input, err := ioutil.ReadFile("input")
  check(err)
  lines := strings.Split(string(input), "\n")
  lines = lines[:len(lines) - 1] // Strip out the whitespace at the end of the line

  // fmt.Println(testSlope(lines, 0, 0, 1, 1))
  // fmt.Println(testSlope(lines, 0, 0, 3, 1))
  // fmt.Println(testSlope(lines, 0, 0, 5, 1))
  // fmt.Println(testSlope(lines, 0, 0, 7, 1))
  // fmt.Println(testSlope(lines, 0, 0, 1, 2))

  fmt.Println(testSlope(lines, 0, 0, 1, 1) *
  testSlope(lines, 0, 0, 3, 1) *
  testSlope(lines, 0, 0, 5, 1) *
  testSlope(lines, 0, 0, 7, 1) *
  testSlope(lines, 0, 0, 1, 2))
}

func testSlope(land []string, index, trees, right, down int) int { // Tail-recursion ;)

  // fmt.Println(len(land))

  if (len(land) < down) { // Base case
    return trees;
  }

  return testSlope(land[down:], index + right, trees + Btoi(land[0][index % len(land[0])] == '#'), right, down)
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func Btoi(o bool) int { // For some branchless
  if (o) {
    return 1
  }
  return 0
}
