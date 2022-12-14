package main

import (
	// "fmt"
	"log"
	"strings"
	"testing"
	// "log"
)

var day5TestInput = ` [D]   
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestDay5Part1(t *testing.T) {

	sections := strings.Split(day5TestInput, "\n\n")

	answer := Day5Part1(sections)
  log.Printf("Answer: %v", answer)

	if answer != "CMZ" {
		t.Fatalf(`Day 5 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 5 Part 1 answer is correct: %v", answer)
	}

}
