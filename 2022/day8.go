package main

import (
	// "fmt"
	"log"
	// "sync"
	// "os"
	// "regexp"
	// "sort"
	// "strconv"
	// "strings"
	// "github.com/k0kubun/pp"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
)

func Day8Part1(trees [][]int) int {
	tallestInCol := make([]int, len(trees[0]))
	tallestInRow := make([]int, len(trees))
	retVal := int(0)

	for r, _ := range trees {
		for c, _ := range trees[r] {
			if trees[r][c] > tallestInCol[c] {
				tallestInCol[c] = trees[r][c]
			}
			if trees[r][c] > tallestInRow[r] {
				tallestInRow[r] = trees[r][c]
			}
		}
	}
	log.Printf("Tallest in Col: %v", tallestInCol)
	log.Printf("Tallest in Row: %v", tallestInRow)

	for r, row := range trees {
		if r == 0 || r == len(trees)-1 {
			for i, tree := range row {
				if i == 0 || i == len(row)-1 && tree == tallestInRow[i] {
					retVal++
				}
			}
		}else{
			for i, tree := range row {
				if i == 0 || i == len(row)-1 && tree == tallestInCol[i] {
					retVal++
				}
			}

    }
	}

	return retVal

}
