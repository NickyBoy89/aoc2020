package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {

  var totalUnique int

  input, err := ioutil.ReadFile("input")
  check(err)

  groups := strings.Split(string(input), "\n\n")

  for _, group := range groups {

    var answers []int

    group = strings.ReplaceAll(group, "\n", "")

    for _, answer := range group {
      if (!containedIn(answers, int(answer))) {
        answers = append(answers, int(answer))
      }
    }

    fmt.Println(len(answers))

    totalUnique += len(answers)

  }

  fmt.Println(totalUnique)

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func containedIn(array []int, element int) bool { // Not the best, runs in O(n) time
  for _, i := range array {
    if (i == element) {
      return true
    }
  }
  return false
}
