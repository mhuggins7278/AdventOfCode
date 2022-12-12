package main

import (
	// "fmt"
	"log"
	// "sync"
	"os"
	// "regexp"
	"strconv"
	"strings"
	// "github.com/k0kubun/pp"
	// "reflect"
	"github.com/mhuggins7278/AdventOfCode/utils"
	"golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)


func Day8() {

	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day8.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	trees := make([][]int, 0)
	for _, line := range input {
		heights := strings.Split(line, "")
		row := make([]int, 0)
		for _, height := range heights {
			tree, _ := strconv.Atoi(height)
			row = append(row, tree)
		}
		trees = append(trees, row)

	}
	// log.Printf("trees is %v long", trees)

	part1Answer := Day8Part1(trees)
	part2Answer := Day8Part2(trees)
	log.Printf("Day8Part1 Answer %v", part1Answer)
	log.Printf("Day8Part2 Answer %v", part2Answer)

}

func Day8Part1(trees [][]int) int {
	retVal := int(0)

	for r, row := range trees {
		if r == 0 || r == len(trees)-1 {
			//if the tree is in the first or last row it's visible no matter what
			retVal += len(row)
			// log.Printf("row %v is visible because it's the first or last row", r)
			// log.Printf("retVal is now %v", retVal)
			continue
		}
		for treeColumnPosition, treeHeight := range row {
			// log.Printf("row: %v", row)
			//if this is the first or last column just add it
			if treeColumnPosition == 0 || treeColumnPosition == len(row)-1 {
				retVal++
				// log.Printf("adding row:%v column: %v is visible because it's the first or last column in the row", r, treeColumnPosition)
				// log.Printf("retVal is now %v", retVal)
				continue
			}
			//check if there are any trees taller than this tree in to the left in the same row
			treesToLeftEdge := make([]int, 0)
			treesToLeftEdge = append(treesToLeftEdge, row[0:treeColumnPosition]...)
			slices.Sort(treesToLeftEdge)
			// log.Printf("Trees with height %v to the Left Edge %v for row: %v col: %v", treeHeight, treesToLeftEdge, r, treeColumnPosition)
			if treeHeight > treesToLeftEdge[len(treesToLeftEdge)-1] {
				// log.Printf("adding row:%v column: %v is visible because there are no taller trees to the left of it", r, treeColumnPosition)
				retVal++
				continue
			}
			//check if there are any trees taller than this tree in to the right in the same row
			treesToRightEdge := make([]int, 0)
			treesToRightEdge = append(treesToRightEdge, row[treeColumnPosition+1:]...)
			slices.Sort(treesToRightEdge)
			// log.Printf("Trees with height %v to the Right Edge %v row: %v", treeHeight, treesToRightEdge, row)
			if treeHeight > treesToRightEdge[len(treesToRightEdge)-1] {
				// log.Printf("adding row:%v column: %v is visible because there are no taller trees to the right of it", r, treeColumnPosition)
				retVal++
				continue
			}
			//check if there are any trees taller than this tree in to the top in the same row
			treesToTopEdge := make([]int, 0)
			// log.Printf("currRow: %v currCol: %v top", r, treeColumnPosition)
			for i := r; i > 0; i-- {
				treesToTopEdge = append(treesToTopEdge, trees[i-1][treeColumnPosition])
			}
			slices.Sort(treesToTopEdge)
			// log.Printf("Trees with height %v to the top Edge %v", treeHeight, treesToTopEdge)
			if treeHeight > treesToTopEdge[len(treesToTopEdge)-1] {
				// log.Printf("adding row:%v column: %v is visible because there are no taller trees above it %v", r, treeColumnPosition, treesToTopEdge)
				retVal++
				continue
			}
			treesToBottomEdge := make([]int, 0)
			// log.Printf("currRow: %v currCol: %v bottom", r, treeColumnPosition)
			for i := r; i < len(trees)-1; i++ {
				treesToBottomEdge = append(treesToBottomEdge, trees[i+1][treeColumnPosition])
			}
			slices.Sort(treesToBottomEdge)
			// log.Printf("Trees with height %v to the bottom Edge %v", treeHeight, treesToBottomEdge)
			if treeHeight > treesToBottomEdge[len(treesToBottomEdge)-1] {
				// log.Printf("adding row:%v column: %v is visible because there are no taller trees below it %v", r, treeColumnPosition, treesToBottomEdge)
				retVal++
				continue
			}

		}
	}
	return retVal
}
func Day8Part2(trees [][]int) int {
	scores := make([]int, 0)

	for r, row := range trees {
		if r == 0 || r == len(trees)-1 {
			// log.Println("Skipping first  and last row")
			continue
		}
		for treeColumnPosition, treeHeight := range row {

			score := 0
			distances := make([]int, 0)
			//if this is the first or last column just add it
			if treeColumnPosition == 0 || treeColumnPosition == len(row)-1 {
				// log.Println("Skipping first  and last column")
				continue
			}
			// log.Printf("Checking row: %v col: %v", r, treeColumnPosition)
			//check if there are any trees taller than this tree in to the left in the same row
			treesToLeftEdge := make([]int, 0)
			treesToLeftEdge = append(treesToLeftEdge, row[0:treeColumnPosition]...)
			utils.ReverseSlice(treesToLeftEdge)
      // log.Printf("to left edge %v", treesToLeftEdge)
			for i, t := range treesToLeftEdge {
				if t >= treeHeight || i == len(treesToLeftEdge) -1 {
					distances = append(distances, i+1)
          break
				}
			}
			//check if there are any trees taller than this tree in to the right in the same row
			treesToRightEdge := make([]int, 0)
			treesToRightEdge = append(treesToRightEdge, row[treeColumnPosition+1:]...)
      // log.Printf("to right edge %v", treesToRightEdge)
			for i, t := range treesToRightEdge {
				if t >= treeHeight || i == len(treesToRightEdge) -1  {
					distances = append(distances, i+1)
          break
				}

			}
			//check if there are any trees taller than this tree in to the top in the same row
			treesToTopEdge := make([]int, 0)
			// log.Printf("currRow: %v currCol: %v top", r, treeColumnPosition)
			for i := r; i > 0; i-- {
				treesToTopEdge = append(treesToTopEdge, trees[i-1][treeColumnPosition])
			}
      // log.Printf("trees to bottom %v", treesToTopEdge)
			// ReverseSlice(treesToTopEdge)
			for i, t := range treesToTopEdge {
				if t >= treeHeight || i == len(treesToTopEdge) -1  {
					distances = append(distances, i+1)
          break
				}

			}
			treesToBottomEdge := make([]int, 0)
			// log.Printf("currRow: %v currCol: %v bottom", r, treeColumnPosition)
			for i := r; i < len(trees)-1; i++ {
				treesToBottomEdge = append(treesToBottomEdge, trees[i+1][treeColumnPosition])
      }
      // ReverseSlice(treesToBottomEdge)
      // log.Printf("trees to bottom %v", treesToBottomEdge)
			for i, t := range treesToBottomEdge {
        // log.Printf("index:%v, currHeight:%v, treeHeight:%v", i,t, treeHeight)
				if t >= treeHeight || i == len(treesToBottomEdge) -1  {
				distances = append(distances, i+1)
          break
				}

			}
			// log.Printf("distances: %v", distances)
			for _, distance := range distances {
				if score == 0 {
					score = distance
					continue
				}
				score = score * distance
			}
			scores = append(scores, score)
		}
	}
	// log.Printf("Scores %v", scores)
	slices.Sort(scores)
	bestScore := scores[len(scores)-1]
	return bestScore

}
