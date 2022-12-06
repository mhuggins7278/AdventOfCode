package main

import (
	"testing"
  "strings"
)

var testInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay4Part1(t *testing.T) {

  lines := strings.Split(testInput, "\n")

	answer := Day4Part1(lines)
	if answer != 2 {
    t.Fatalf(`Day 4 Part 1 answer is incorrect: %v`, answer)
	} else {
    t.Logf("Day 4 Part 1 anser is correct: %v", answer)
  }
}

func TestDay4Part2(t *testing.T) {

  lines := strings.Split(testInput, "\n")

	answer := Day4Part2(lines)
	if answer != 3 {
    t.Fatalf(`Day 4 Part 2 answer is incorrect: %v`, answer)
	} else {
    t.Logf("Day 4 Part 2 anser is correct: %v", answer)
  }
}


