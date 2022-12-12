package main

import (
	"fmt"
	"log"
	// "sync"
	"os"
	// "regexp"
	// "github.com/k0kubun/pp"
	"math"
	"strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)

type Node struct {
	nodeNum       int
	x             int
	y             int
	visitedSpaces map[string]int
	next          *Node
	prev          *Node
}

func (n *Node) AddToTail(nodeNum int, x int, y int) {
	if n.next == nil {
		n.next = &Node{
			nodeNum:       nodeNum,
			x:             x,
			y:             y,
			visitedSpaces: map[string]int{"1 1": 1},
			prev:          n,
		}
		return
	}
	n.next.AddToTail(nodeNum, x, y)
}

func (n *Node) Move(direction string) {
	// log.Printf("Moving node %v %v from %v:%v", n.nodeNum, direction, n.x, n.y)
	if n.nodeNum != 0 {
		if math.Abs(float64(n.x-n.prev.x)) <= 1 && math.Abs(float64(n.y-n.prev.y)) <= 1 {
			// log.Printf("Skipping move node is within one position of it's parent x node %v prev %v", n.x, n.prev.x)
			return
		}
	}
	switch direction {
	case "U":
		if n.prev == nil {
			n.y = n.y + 1
		} else {
			n.y = n.prev.y - 1
			n.x = n.prev.x
		}
	case "D":
		if n.prev == nil {
			n.y = n.y - 1
		} else {
			n.y = n.prev.y + 1
			n.x = n.prev.x
		}
	case "R":
		if n.prev == nil {
			n.x = n.x + 1
		} else {
			n.x = n.prev.x - 1
			n.y = n.prev.y
		}
	case "L":
		if n.prev == nil {
			n.x = n.x - 1
		} else {
			n.x = n.prev.x + 1
			n.y = n.prev.y
		}
	}
	n.visitedSpaces[fmt.Sprint(n.x, n.y)] = 1
	// log.Printf("Node %v moved to %v %v", n.nodeNum, n.x, n.y)
	if n.next != nil {
		n.next.Move(direction)
	}

}

func Day9() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day9.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	// part1 := Day9Part1(input, 2)
	// log.Printf("Part 1 answer is  %v", part1)
	part2 := Day9Part2(input, 10)
	log.Printf("Part 2 answer is  %v", part2)

}

func Day9Part1(moves []string, size int) int {
	head := &Node{nodeNum: 0,
		x:             1,
		y:             1,
		visitedSpaces: map[string]int{"1 1": 1},
	}
	for i := 1; i < size; i++ {
		head.AddToTail(i, 1, 1)
	}
	for _, move := range moves {
		// head.Move(direction, numTimes)
		// log.Printf("Move is %v", move)
		moves := strings.Split(move, " ")
		direction := moves[0]
		numTimes, _ := strconv.Atoi(moves[1])
		for i := 0; i < numTimes; i++ {
			head.Move(direction)
		}
	}
	// pp.Println(head.next)
	retVal := len(head.next.visitedSpaces)
	// log.Printf("Return value is %v", retVal)

	return retVal
}
func Day9Part2(moves []string, size int) int {
	head := &Node{nodeNum: 0,
		x:             12,
		y:             6,
		visitedSpaces: map[string]int{"1 1": 1},
	}
	for i := 1; i < size; i++ {
		head.AddToTail(i, 12, 6)
	}
	for _, move := range moves {
		// head.Move(direction, numTimes)
		// log.Printf("Move is %v", move)
		moves := strings.Split(move, " ")
		direction := moves[0]
		numTimes, _ := strconv.Atoi(moves[1])
		for i := 0; i < numTimes; i++ {
			head.Move(direction)
		}
		currentNode := head
		for currentNode != nil {
			log.Printf("Node:%v x:%v y:%v", currentNode.nodeNum, currentNode.x, currentNode.y)
			currentNode = currentNode.next
		}
	}
	var tail *Node
	currentNode := head

	// Iterate through the list until we reach the end
	for currentNode != nil {
		// Save the current node as the tail
		tail = currentNode
		// Move to the next node in the list
		currentNode = currentNode.next
	}
	// pp.Println(tail)

	retVal := len(tail.visitedSpaces)

	return retVal
}
