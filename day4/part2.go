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
  lines := strings.Split(string(input), "\n\n")

  fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

  // conditions := map[string][]string{
  //   "between": {"1920", "2002"},
  //   "between": {"2010", "2002"},
  //   "between": {"2020", "2030"},
  //   "between": {"if": {}"regex", {".+(cm)"}}
  // }

  for _, val := range lines {
    passValid := 1

    for i := 0; i < len(fields); i++ {

      keyTest := regexp.MustCompile(fields[i])

      re := regexp.MustCompile(fields[i] + ":[^ ]+")
      keyval := re.FindString(val)

      if (len(keyval) == 0) { // Key not found
        passValid = 0
        break
      }

      fmt.Println(keyval)

      fieldValue, err := strconv.Atoi(keyval[len(fields[i]) + 1:])
      fmt.Println(fieldValue)

      check(err)

      if (i == 0) {
        if (fieldValue < 1920 || fieldValue > 2002) {
          passValid = 0
          break
        }
      } else if (i == 1) {
        if (fieldValue < 2010 || fieldValue > 2002) {
          passValid = 0
          break
        }
      } else if (i == 2) {
        if (fieldValue < 2020 || fieldValue > 2030) {
          passValid = 0
          break
        }
      } else if (i == 3) {
        cm := regexp.MustCompile(".+(cm)")
        in := regexp.MustCompile(".+(in)")

        if (cm.MatchString(keyval)) {
          cmVal, err := strconv.Atoi(keyval[len(fields[i]) + 1:len(keyval) - 3])
          check(err)

          if (cmVal < 105 || cmVal > 193) {
            passValid = 0
            break
          }
        } else if (in.MatchString(keyval)) {
          inVal, err := strconv.Atoi(keyval[len(fields[i]) + 1:len(keyval) - 3])
          check(err)

          if (inVal < 59 || inVal > 76) {
            passValid = 0
            break
          }
        }
      } else if (i == 4) {
        height := regexp.MustCompile("#[0-9a-f]*")
        if (!height.MatchString(values[1])) {
          passValid = 0
          break
        }
      } else if (i == 5) {
        eyeCl := regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")
        if (!eyeCl.MatchString(values[1])) {
          passValid = 0
          break
        }
      } else if (i == 6) {
        passpId := regexp.MustCompile("[0-9]{9}$")
        if (!passpId.MatchString(values[1])) {
          passValid = 0
          break
        }
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
