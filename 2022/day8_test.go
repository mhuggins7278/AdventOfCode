
package main

import (
	// "fmt"
	"log"
	"strings"
	"strconv"
	"testing"
) 

var day8test = `30373
25512
65332
33549
35390`

func TestDay8Part1(t *testing.T) {
  lines := strings.Split(day8test, "\n")
  trees := make([][]int, 0)

	for _, line := range lines {
    heights := strings.Split(line, "")
    row := make([]int, 0)
    for _, height := range heights {
      tree, _ := strconv.Atoi(height)
      row = append(row, tree)
    } 
    trees = append(trees, row)
  }
  log.Printf("trees is %v long", trees)

  answer := Day8Part1(trees)
  log.Println(answer)
  
	if answer != 21 {
		t.Fatalf(`Day 8 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 8 Part 1 answer is correct: %v", answer)
	}
}
func TestDay8Part2(t *testing.T) {
  lines := strings.Split(day8test, "\n")
  trees := make([][]int, 0)

	for _, line := range lines {
    heights := strings.Split(line, "")
    row := make([]int, 0)
    for _, height := range heights {
      tree, _ := strconv.Atoi(height)
      row = append(row, tree)
    } 
    trees = append(trees, row)
  }

  answer := Day8Part2(trees)
  
	if answer != 21 {
		t.Fatalf(`Day 8 Part 2 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 8 Part 2 answer is correct: %v", answer)
	}
}

// func TestDay7Part2(t *testing.T) {
//   lines := strings.Split(day7test, "\n")
// 	fileSystem := BuildFileSystem(lines)
// answer := Day7Part2(fileSystem)
//   log.Println(answer)
// 	if answer != 24933642{
// 		t.Fatalf(`Day 7 Part 1 answer is incorrect: %v`, answer)
// 	} else {
// 		t.Logf("Day 7 Part 1 answer is correct: %v", answer)
// 	}

// }
