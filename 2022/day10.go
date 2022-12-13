package main

import (
	// "fmt"
	"log"

	// "sync"
	"os"
	// "regexp"
	// "github.com/k0kubun/pp"
	// "math"
	"strconv"
	"strings"

	// "github.com/k0kubun/pp"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)


func Day10() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day9.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	// part1 := Day9Part1(input, 2)
	// log.Printf("Part 1 answer is  %v", part1)
	part2 := Day9Part2(input, 10)
	log.Printf("Part 2 answer is  %v", part2)

}

func Day10Part1(moves []string) int {
  registerValues := make([]int, 0)
  registerValue := 1
  var currIncrementer int
	for _, move := range moves {
		// head.Move(direction, numTimes)
		// log.Printf("Move is %v", move)
		moveParts := strings.Split(move, " ")
		inc, err := strconv.Atoi(moveParts[1])
      currIncrementer = inc
    if err != nil {
      continue
    } else {
      currIncrementer = inc
    }
	}
	// pp.Println(head.next)
	retVal := len(head.next.visitedSpaces)
	// log.Printf("Return value is %v", retVal)

	return retVal
}

// func Day9Part2(moves []string, size int) int {
// 	head := &Node{nodeNum: 0,
// 		x:             12,
// 		y:             6,
// 		visitedSpaces: map[string]int{"12 6": 1},
// 	}
// 	for i := 1; i < size; i++ {
// 		head.AddToTail(i, 12, 6)
// 	}
// 	pp.Print(head)
// 	for _, move := range moves {
// 		log.Printf("Move is %v", move)
// 		moves := strings.Split(move, " ")
// 		direction := moves[0]
// 		numTimes, _ := strconv.Atoi(moves[1])
// 		for i := 0; i < numTimes; i++ {
// 			head.Move(direction)
// 		}
// 		currentNode := head
// 		for currentNode != nil {
// 			log.Printf("Node:%v x:%v y:%v", currentNode.nodeNum, currentNode.x, currentNode.y)
// 			currentNode = currentNode.prev
// 		}
// 	}
// 	var tail *Node
// 	currentNode := head

// 	// Iterate through the list until we reach the end
// 	for currentNode != nil {
// 		// Save the current node as the tail
// 		tail = currentNode
// 		// Move to the next node in the list
// 		currentNode = currentNode.prev
// 	}
// 	// pp.Println(tail.visitedSpaces)

// 	retVal := len(tail.visitedSpaces)

// 	return retVal
// }
