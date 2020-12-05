package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "regexp"
  "strconv"
)

func main() {

  validPasswords := 0

  input, err := ioutil.ReadFile("input")
  check(err)
  lines := strings.Split(string(input), "\n")

  re := regexp.MustCompile(`[^:]*`)

  for _, val := range lines {

    matched := re.FindAllString(val, 2)

    if (len(matched) < 2) {
      continue
    }

    prefix := matched[0]
    suffix := matched[1]

    var targetChar byte = prefix[len(prefix) - 1]
    var countMin, countMax int

    for i := 0; i < len(prefix); i++ {
      if (prefix[i] == '-') {
        countMin, _ = strconv.Atoi(prefix[0:i])
        countMax, _ = strconv.Atoi(prefix[i+1:len(prefix)-2])
      }
    }

    targetCharCount := 0

    for _, character := range suffix {
      if (character == rune(targetChar)) {
        targetCharCount++
      }
    }

    if (targetCharCount >= countMin && targetCharCount <= countMax) {
      validPasswords++
    }

  }

  fmt.Println(validPasswords)

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}
