package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "regexp"
  "strconv"
)

func main() {

  var valid int = 0

  input, err := ioutil.ReadFile("input")
  check(err)
  lines := strings.Split(string(input), "\n\n") // Separate by blank lines (two newlines in a row)

  var passports []map[string]string

  for _, passport := range lines {

    values := make(map[string]string)

    passport := strings.ReplaceAll(passport, "\n", " ")

    keyGroups := strings.Split(passport, " ")

    for _, group := range keyGroups {

      keyval := strings.Split(group, ":")

      if (len(keyval) < 2) { // No value
        break
      }

      values[keyval[0]] = keyval[1]

    }

    passports = append(passports, values)

  }

  for _, pass := range passports {
    if (isValid(pass)) {
      valid++
    }
    fmt.Println(pass)
  }

  fmt.Println(valid)
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}

func isValid(field map[string]string) bool {

  fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

  for i := 0; i < len(fields); i++ {

    if (field[fields[i]] == "") {
      return false
    }

    content := field[fields[i]]

    contentNum, err := strconv.Atoi(content)
    if (err != nil) {
      contentNum = -1
    }

    if (i == 0) { // byr
      if (contentNum < 1920 || contentNum > 2002) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }

      fmt.Printf("Passed test %d with value %v\n", i, content)

    } else if (i == 1) { // iyr
      if (contentNum < 2010 || contentNum > 2020) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    } else if (i == 2) { // eyr
      if (contentNum < 2020 || contentNum > 2030) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    } else if (i == 3) { // hgt

      cm := regexp.MustCompile(".+(cm)")
      in := regexp.MustCompile(".+(in)")

      if (cm.MatchString(content)) { // Found string "cm" in content

        heightValue, err := strconv.Atoi(content[:len(content) - 2]) // Strip the last 2 letters
        check(err)

        if (heightValue < 105 || heightValue > 193) {
          fmt.Printf("Failed test %d with value %v\n", i, content)
          return false
        }
      } else if (in.MatchString(content)) {  // Found "in"

        heightValue, err := strconv.Atoi(content[:len(content) - 2]) // Strip the last 2 letters
        check(err)

        if (heightValue < 59 || heightValue > 76) {
          fmt.Printf("Failed test %d with value %v\n", i, content)
          return false
        }

      } else {
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    } else if (i == 4) { // hcl

      height := regexp.MustCompile("#[0-9a-f]{6}")

      if (len(content) != 7 || !height.MatchString(content)) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    } else if (i == 5) { // ecl

      eyeCl := regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")

      if (!eyeCl.MatchString(content)) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    } else if (i == 6) { // pid
      passpId := regexp.MustCompile("[0-9]{9}$")

      if (len(content) != 9 || !passpId.MatchString(content)) {
        fmt.Printf("Failed test %d with value %v\n", i, content)
        return false
      }
      fmt.Printf("Passed test %d with value %v\n", i, content)
    }

  }

  return true
}
