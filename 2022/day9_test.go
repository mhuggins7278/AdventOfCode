package main

import (
	// "fmt"
	// "log"
	"log"
	"strings"
	// "strconv"
	"testing"
) 

var day9test = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var day9test2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

// func TestDay9Part1(t *testing.T) {
//   lines := strings.Split(day9test, "\n")
//   log.Printf("Moves: %v", lines)

//   answer := Day9Part1(lines, 2)

//   // answer := Day8Part1(trees)
//   // log.Println(answer)
//   
// 	if answer != 13 {
// 		t.Fatalf(`Day 9 Part 1 answer is incorrect: %v`, answer)
// 	} else {
// 		t.Logf("Day 9 Part 1 answer is correct: %v", answer)
// 	}
// }
func TestDay9Part2(t *testing.T) {
  lines := strings.Split(day9test2, "\n")
  log.Printf("Moves: %v", lines)

  answer := Day9Part2(lines, 10)

  // answer := Day8Part1(trees)
  // log.Println(answer)
  
	if answer != 36 {
		t.Fatalf(`Day 9 Part 2 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 9 Part 2 answer is correct: %v", answer)
	}
}
