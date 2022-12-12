package main

import (
	"fmt"
	"log"
	// "sync"
	"os"
	// "regexp"
	"github.com/k0kubun/pp"
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

func (n *Node) AddToTail(nodeNum int) {
	if n.next == nil {
		n.next = &Node{
			nodeNum:       nodeNum,
			x:             0,
			y:             0,
			visitedSpaces: map[string]int{"00": 1},
			prev:          n,
		}
		return
	}
	n.next.AddToTail(nodeNum)
}

func (n *Node) Move(direction string, numTimes int) {
	log.Printf("Moving node %v: %v %v", n.nodeNum, direction, numTimes)
	if numTimes == 0 {
		return
	}
	if n.nodeNum != 0 {
		if math.Abs(float64(n.x-n.prev.x)) <= 1 && math.Abs(float64(n.y-n.prev.y)) <= 1 {
			log.Printf("Skipping move node is within one position of it's parent x node %v prev %v", n.x, n.prev.x)
			return

		}
	}
	switch direction {
	case "U":
		if n.prev == nil || n.prev.x == n.x  {
			n.y = n.y + 1
		} else if n.x > n.prev.x {
			n.y = n.y + 1
			n.x = n.x - 1
		} else if n.x < n.prev.x {
			n.y = n.y + 1
			n.x = n.x + 1
		}
	case "D":
		if n.prev == nil || n.x == n.prev.x {
			n.y = n.y - 1
		} else if n.x > n.prev.x {
			n.y = n.y - 1
			n.x = n.x - 1
		} else if n.x < n.prev.x {
			n.y = n.y - 1
			n.x = n.x + 1
		}
	case "R":
		if n.prev == nil || n.prev.y == n.y {
			n.x = n.x + 1
		} else if n.y > n.prev.y {
			n.y = n.y - 1
			n.x = n.x + 1
		} else if n.y < n.prev.y {
			n.y = n.y + 1
			n.x = n.x + 1
		}
	case "L":
		if n.prev == nil || n.y == n.prev.y  {
			n.x = n.x - 1
		} else if n.y > n.prev.y {
			n.y = n.y - 1
			n.x = n.x - 1
		} else if n.y < n.prev.y {
			n.y = n.y + 1
			n.x = n.x - 1
		}
	}
	if _, exists := n.visitedSpaces[fmt.Sprint(n.x,n.y)]; exists {
		return
	}
	n.visitedSpaces[fmt.Sprint(n.x, n.y)] = 1
  log.Printf("Node %v moved to %v %v", n.nodeNum, n.x, n.y)
	if n.next != nil {
		n.next.Move(direction, numTimes)
	}

}

func Day9() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day9.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	Day9Part1(input, 2)

}

func Day9Part1(moves []string, size int) int {
	head := &Node{nodeNum: 0,
		x:             0,
		y:             0,
		visitedSpaces: map[string]int{"00": 1},
	}
	for i := 1; i < size; i++ {
		head.AddToTail(i)
	}
	for _, move := range moves {
		// head.Move(direction, numTimes)
		log.Printf("Move is %v", move)
		moves := strings.Split(move, " ")
		direction := moves[0]
		numTimes, _ := strconv.Atoi(moves[1])
		for i := 0; i < numTimes; i++ {
			head.Move(direction, numTimes-i)
		}
	}
  pp.Println(head.next)
	retVal := len(head.next.visitedSpaces)
log.Printf("Return value is %v", retVal)


	return retVal
}
