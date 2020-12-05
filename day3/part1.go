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

  fmt.Println(testSlope(lines, 0, 0))
}

func testSlope(land []string, index, trees int) int { // Tail-recursion ;)

  // fmt.Println(len(land))

  if (len(land[0]) == 0) { // Base case
    return trees
  }

  for (index > len(land[0]) - 1) { // Accounts for genetic sequences in trees
    index = index - len(land[0])
  }

  treeExists := 0

  if (land[0][index] == '#') {
    treeExists = 1
  }

  return testSlope(land[1:], index + 3, trees + treeExists)
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}
