package main

import (
	// "fmt"
	"log"

	// "sync"
	"os"
	// "regexp"
	// "github.com/k0kubun/pp"
	// "math"
	// "github.com/k0kubun/pp"
	"strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)

type QueueItem struct {
	tickAdded, value int
}

type Queue []QueueItem

func (q *Queue) Push(n QueueItem) {
	*q = append(*q, n)
}
func (q *Queue) Pop() QueueItem {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}
func (q *Queue) Peek() QueueItem {
	var n QueueItem
	if len(*q) > 0 {
		n = (*q)[0]
	}
	return n

}

func Day10() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day9.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	// part1 := Day9Part1(input, 2)
	// log.Printf("Part 1 answer is  %v", part1)
	part1 := Day10Part1(input)
	log.Printf("Part 2 answer is  %v", part1)

}

func Day10Part1(moves []string) int {
	registerValues := Queue{}
	registerValue := 1
	computedValues := []int{}
	cycleToCheck := 20
	retVal := 0
	noops := 0

	for i, move := range moves {
		tick := i + 1
		cycle := (tick * 2) - noops
		log.Printf("noops is %v", noops)
		log.Printf("Register value is %v at the start of cycle %v", registerValue, cycle)
		moveParts := strings.Split(move, " ")
		// log.Printf("Move: %v", moveParts)
		if len(moveParts) == 2 {
			inc, _ := strconv.Atoi(moveParts[1])
			// log.Printf("Adding value: %v to queue on tick %v", inc, tick)
			registerValues.Push(QueueItem{tick, inc})
		} else {
			noops = noops + 1
      continue
		}
		if cycle == cycleToCheck {
			log.Printf("Computing new value: Register value at cycle %v is %v", cycle, registerValue)
			computedValues = append(computedValues, registerValue*cycle)
			cycleToCheck = cycleToCheck + 40

		}
		//this has to happen at the end of the cycle nothing should be after this
		if len(registerValues) > 0 {
			curValue := registerValues.Peek()
			if tick-curValue.tickAdded == 1 {

				// log.Printf("Poppping %v off the queue at tick %v",curValue, tick)
				registerValue += curValue.value
				registerValues.Pop()
			}
		}
		log.Printf("Register value is %v at the end of cycle %v", registerValue, cycle)

	}
	log.Printf("computedValues: %v", computedValues)

	for _, a := range computedValues {
		retVal = retVal + a
	}
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
