package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Instruction struct {
  code string
  val int
}

func main() {

  input, err := ioutil.ReadFile("input")
  check(err)

  instr := strings.Split(string(input), "\n")

  instr = instr[:len(instr) - 1] // Strip out the last whitespace

  var instructions []Instruction

  for _, i := range instr {
    instr := strings.Split(i, " ")

    op := instr[0]
    val, err := strconv.Atoi(instr[1])
    check(err)

    instructions = append(instructions, Instruction{
      code: op,
      val: val,
    })

  }

  for j, i := range instructions {

    testInstructions := make([]Instruction, len(instructions))

    copy(testInstructions, instructions)

    if (i.code == "jmp") {
      testInstructions[j].code = "nop" // Try replacing the instruction

      status, result := didExecute(testInstructions)

      if (status) {
        fmt.Println(result)
      }
    } else if (i.code == "nop") {
      testInstructions[j].code = "jmp"

      status, result := didExecute(testInstructions)

      if (status) {
        fmt.Println(result)
      }
    }
  }

}

func didExecute(instructions []Instruction) (bool, int) {

  accumulator := 0

  currInst := 0

  visited := make([]bool, len(instructions))

  for ; !visited[currInst]; {

    visited[currInst] = true // Update visited

    if (instructions[currInst].code == "acc") {
      accumulator += instructions[currInst].val // Increment the accumulator
      currInst++

    } else if (instructions[currInst].code == "jmp") {
      currInst += instructions[currInst].val

    } else if (instructions[currInst].code == "nop") {
      currInst++

    }

    if (currInst == len(instructions)) {
      return true, accumulator
    }
  }

  return false, accumulator
}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}
