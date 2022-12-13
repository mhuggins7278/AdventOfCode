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

	"github.com/k0kubun/pp"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)

type Delta struct {
	x, y int
}

type Node struct {
	nodeNum       int
	x             int
	y             int
	visitedSpaces map[string]int
	next          *Node
	prev          *Node
}

func (n *Node) AddToTail(nodeNum int, x int, y int) {
	if n.prev == nil {
		n.prev = &Node{
			nodeNum:       nodeNum,
			x:             x,
			y:             y,
			visitedSpaces: map[string]int{"12 6": 1},
			next:          n,
		}
		return
	}
	n.prev.AddToTail(nodeNum, x, y)
}

func (n *Node) Move(direction string) {
	log.Printf("trying to move node:%v %v from %v:%v", n.nodeNum, direction, n.x, n.y)
	//if this is the head node just move it and try to move the next node
	if n.nodeNum == 0 {
		// log.Printf("Moving head node")
		switch direction {
		case "U":
			n.y++
		case "D":
			n.y--
		case "R":
			n.x++
		case "L":
			n.x--
		}
		log.Printf("Node %v moved to %v %v", n.nodeNum, n.x, n.y)
		n.visitedSpaces[fmt.Sprint(n.x, n.y)] = 1
		n.prev.Move(direction)
		return
	}
	// log.Printf("Moving non-head node %v", n.nodeNum)
	delta := Delta{n.x - n.next.x, n.y - n.next.y}
	log.Printf("Delta is %v", delta)
	// if  the node is within one spot of the node before it nothing to do
	if math.Abs(float64(delta.x)) <= 1 && math.Abs(float64(delta.y)) <= 1 {
		// log.Printf("Skipping move node %v is within one position of it's parent node %v prev %v", n.nodeNum, n.x, n.prev.x)
		return
	}

	if delta.y < 0 {
		n.y = n.y + 1
	} else if delta.y > 0 {
		n.y = n.y - 1
	} 
  if delta.x > 0 {
		n.x = n.x - 1
	} else if delta.x < 0 {
		n.x = n.x + 1
	}

	n.visitedSpaces[fmt.Sprint(n.x, n.y)] = 1

	log.Printf("Node %v moved to %v %v", n.nodeNum, n.x, n.y)
	if n.prev != nil {
		n.prev.Move(direction)
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
		x:             0,
		y:             0,
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
		visitedSpaces: map[string]int{"12 6": 1},
	}
	for i := 1; i < size; i++ {
		head.AddToTail(i, 12, 6)
	}
	pp.Print(head)
	for _, move := range moves {
		log.Printf("Move is %v", move)
		moves := strings.Split(move, " ")
		direction := moves[0]
		numTimes, _ := strconv.Atoi(moves[1])
		for i := 0; i < numTimes; i++ {
			head.Move(direction)
		}
		currentNode := head
		for currentNode != nil {
			log.Printf("Node:%v x:%v y:%v", currentNode.nodeNum, currentNode.x, currentNode.y)
			currentNode = currentNode.prev
		}
	}
	var tail *Node
	currentNode := head

	// Iterate through the list until we reach the end
	for currentNode != nil {
		// Save the current node as the tail
		tail = currentNode
		// Move to the next node in the list
		currentNode = currentNode.prev
	}
	// pp.Println(tail.visitedSpaces)

	retVal := len(tail.visitedSpaces)

	return retVal
}
