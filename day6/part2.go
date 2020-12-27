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

    groupAnswers := make(map[int]int)

    groupMembers := strings.Split(group, "\n")

    blankMembers := 0 // Avoids the costly effort of removing the item from the list

    for _, person := range groupMembers {

      fmt.Printf("Person: %v\n", person)

      if (person == "") {
        blankMembers++
        break
      }

      var answers []int

      for _, answer := range person {
        if (!containedIn(answers, int(answer))) {
          answers = append(answers, int(answer))
        }
      }

      for _, groupAnswer := range answers {
        groupAnswers[groupAnswer] += 1
      }

    }

    fmt.Println(groupAnswers, len(groupMembers) - blankMembers)

    start := totalUnique

    for _, unique := range groupAnswers {
      if (unique == len(groupMembers) - blankMembers) {
        totalUnique++
      }
    }

    fmt.Println(totalUnique - start)

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

func remove(array []int, element int) []int {
  for i, e := range array {
    if (e == element) {
      return array[i - 1:i]
    }
  }
  return array // Not in the array, no change needed
}
