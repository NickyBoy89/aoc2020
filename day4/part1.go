package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "regexp"
)

func main() {

  var valid int = 0

  input, err := ioutil.ReadFile("input")
  check(err)
  lines := strings.Split(string(input), "\n\n")

  fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

  for _, val := range lines {
    passValid := 1

    for i := 0; i < len(fields); i++ {
      re := regexp.MustCompile(fields[i])
      if (!re.MatchString(val)) {
        passValid = 0
        break
      }
    }

    valid += passValid
  }

  fmt.Println(valid)
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}
