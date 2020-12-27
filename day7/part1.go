package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "regexp"
)

type Bag struct {
  name string
  number int
  contents []Bag
}

func main() {

  var bags []Bag

  input, err := ioutil.ReadFile("input")
  check(err)

  rules := strings.Split(string(input), "\n")

  rules = rules[:len(rules) - 1] // Strip out the last whitespace

  for _, rule := range rules {

    bagMatch, err := regexp.Compile(`[^ ]+\s+[^ ]+.(bag)`)
    check(err)

    matchedBags := bagMatch.FindAllString(rule, -1)

    bags = append(bags, Bag{name: strings.Trim(matchedBags[0][:len(matchedBags[0]) - 4], " "), number: 0, contents: populateContents(matchedBags[1:])})

  }

  var shinyGold []Bag

  for _, i := range bags {
    if (len(searchBags(i.contents, "shiny gold")) > 0) {
      shinyGold = append(shinyGold, i)
    }
  }

  validContainers := shinyGold

  for _, j := range shinyGold {
    validContainers = append(validContainers, findAllNested(bags, j.name)...)
  }

  fmt.Println(shinyGold)

  fmt.Println(validContainers)

  fmt.Println(len(removeDuplicates(validContainers)))

}

func removeDuplicates(bags []Bag) []Bag {

  var result []Bag

  for _, i := range bags {
    if (!containedIn(result, i.name)) {
      result = append(result, i)
    }
  }

  return result
}

func containedIn(array []Bag, element string) bool {
  for _, i := range array {
    if (i.name == element) {
      return true
    }
  }
  return false
}

func findAllNested(bags []Bag, element string) []Bag {

  fmt.Println(element)

  var validContainers []Bag

  for _, i := range bags {
    if (len(searchBags(i.contents, element)) > 0) {
      validContainers = append(validContainers, i)
    }
  }

  for _, j := range validContainers {
    validContainers = append(validContainers, findAllNested(bags, j.name)...)
  }

  return validContainers
}

func searchBags(bags []Bag, element string) []Bag {

  if (bags == nil) {
    return nil
  }

  var valid []Bag

  for _, i := range bags {
    if (searchContents(i, element)) {
      valid = append(valid, i)
    }
  }

  return valid
}

func searchContents(contents Bag, element string) bool {

  if (contents.name == element) {
    return true
  }

  for _, i := range contents.contents {
    if (searchContents(i, element)) {
      return true
    }
  }

  // No match found
  return false
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func populateContents(contentBags []string) []Bag {
  var contents []Bag

  for _, i := range contentBags {
    contents = append(contents, Bag{name: strings.Trim(i[:len(i) - 4], " "), number: 0, contents: nil})
  }

  return contents
}
