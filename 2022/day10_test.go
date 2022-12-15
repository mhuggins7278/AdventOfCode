package main

import (
	"fmt"
	// "log"
	"log"
	"strings"
	// "strconv"
	"testing"
) 

var day10test = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

var day10simple = `noop
addx 3
addx -5
noop`



func TestDay10Part1(t *testing.T) {
  lines := strings.Split(day10test, "\n")
  log.Printf("Moves: %v", lines)

  answer := Day10Part1(lines)

  log.Printf("Day10 part 1 test %v",answer)
  
	if answer !=  13140 {
		t.Fatalf(`Day 10 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 10 Part 1 answer is correct: %v", answer)
	}
}
func TestDay10Part2(t *testing.T) {
  lines := strings.Split(day10test, "\n")
  log.Printf("Moves: %v", lines)

  answer := Day10Part2(lines)


  for _, line := range answer {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
  
	if true != false   {
		t.Fatalf(`Day 10 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 10 Part 1 answer is correct: %v", answer)
	}
}
// func TestDay9Part2(t *testing.T) {
//   lines := strings.Split(day9test2, "\n")
//   log.Printf("Moves: %v", lines)

//   answer := Day9Part2(lines, 10)

//   // answer := Day8Part1(trees)
//   // log.Println(answer)
//   
// 	if answer != 36 {
// 		t.Fatalf(`Day 9 Part 2 answer is incorrect: %v`, answer)
// 	} else {
// 		t.Logf("Day 9 Part 2 answer is correct: %v", answer)
// 	}
// }
