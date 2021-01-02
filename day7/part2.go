package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "regexp"
  "strconv"
)

type Bag struct {
  name string
  number int
  contents []Bag
}

func (b *Bag) sumContents() int {
  var total int

  for _, i := range b.contents {
    total += i.number
  }

  return total
}

func main() {

  var bags []Bag

  input, err := ioutil.ReadFile("input")
  check(err)

  rules := strings.Split(string(input), "\n")

  rules = rules[:len(rules) - 1] // Strip out the last whitespace

  // Parse every line
  for _, rule := range rules {

    // Matches the two words before the word "bag"
    bagMatch := regexp.MustCompile(`[^ ]+\s+[^ ]+.(bag)`)

    // Matches all the numbers in the string
    numberMatches := regexp.MustCompile(`[0-9]+`)

    // Get matched values
    matchedBags := bagMatch.FindAllString(rule, -1)
    matchedNumbers := numberMatches.FindAllString(rule, -1)

    matchedBag := strings.Trim(matchedBags[0][:len(matchedBags[0]) - 4], " ")

    bags = append(bags, Bag{
      name: matchedBag,
      number: 0,
      contents: populateContents(matchedBags[1:], matchedNumbers),
    })

  }

  // Find the shiny gold bag
  for _, i := range bags {
    if (i.name == "shiny gold") {
      // NOTE: This assumes that there are no duplicates in bag definitions
      fmt.Println(sumNestedBags(i, bags))
    }
  }

}

func sumNestedBags(target Bag, bags []Bag) int {
  var total int

  for _, i := range target.contents {
    total += i.number * sumNestedBags(findBag(i.name, bags), bags)
  }

  return total + target.sumContents()
}

func findBag(bagName string, bags []Bag) Bag {
  for _, i := range bags {
    if (i.name == bagName) {
      return i
    }
  }

  panic("Bag not found")
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

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func populateContents(contentBags []string, numbers []string) []Bag {

  var contents []Bag

  for j, i := range contentBags {

    bagName := strings.Trim(i[:len(i) - 4], " ")

    // Handles case where there is no other bag contained in the bag
    if (bagName == "no other") {
      continue
    }

    bagContentsNumber, err := strconv.Atoi(numbers[j])
    check(err)

    contents = append(contents, Bag{
      name: bagName,
      number: bagContentsNumber,
      contents: nil,
    })
  }

  return contents
}
