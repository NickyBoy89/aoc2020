package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "sort"
  "strconv"
)

func main() {

  var inputs []string
  var inputNumbers []int

  reader := bufio.NewReader(os.Stdin)

  for {

    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    inputs = append(inputs, text)

    if strings.Compare("end", text) == 0 {
      inputs = inputs[:len(inputs) - 1]
      inputNumbers = stringArrayToIntArray(inputs)
      sort.Ints(inputNumbers)
      break
    }

  }

  for _, first := range inputNumbers { // Since we are not finding two numbers, a triple-nested for loop seems easiest, although it runs at O(n^3)
    for _, second := range inputNumbers {
      for _, third := range inputNumbers {
        if (first + second + third == 2020) {
          fmt.Println(first * second * third)
        }
      }
    }
  }

}

func stringArrayToIntArray(stringArray []string) []int {

  var intArray []int

  for _, i := range stringArray {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        intArray = append(intArray, j)
    }

  return intArray
}
