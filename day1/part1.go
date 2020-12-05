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
    text = strings.Replace(text, "\n", "", -1) // Reads input from STDIN, not the most convenient way, but works

    inputs = append(inputs, text)

    if strings.Compare("end", text) == 0 {
      inputs = inputs[:len(inputs) - 1]
      inputNumbers = stringArrayToIntArray(inputs)
      sort.Ints(inputNumbers) // Uses quicksort, means a O(n log(n)) average
      break
    }

  }

  for _, element := range inputNumbers {
    if (inputNumbers[sort.SearchInts(inputNumbers, getComplement(element))] == getComplement(element)) { // Binary search, O(log(n))
      fmt.Println(element * getComplement(element))
    }
  }
}

func getComplement(number int) int {
  return 2020 - number
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
