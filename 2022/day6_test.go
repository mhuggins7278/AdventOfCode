package main

import (
	// "fmt"
	// "log"
	// "strings"
	"testing"
	// "log"
)

var input1 = []byte("bvwbjplbgvbhsrlpgdmjqwftvncz")
var input2 = []byte("nppdvjthqldpwncqszvftbrmjlhg")
var input3 = []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
var input4 = []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

func TestDay6Part1(t *testing.T) {
	answer := Day6Part1(input4, 4)
	if answer != 5 {
		t.Fatalf(`Day 6 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 6 Part 1 answer is correct: %v", answer)
	}
}
