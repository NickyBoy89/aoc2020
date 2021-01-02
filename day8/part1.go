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

  accumulator := 0

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
  }

  fmt.Println(accumulator)

}

func check(err error) {
  if (err != nil) {
    panic(err)
  }
}
