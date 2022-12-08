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
	"github.com/adam-lavrik/go-imath/ix"
)

func Day8Part1(trees [][]int) int {
	posTallestInCol := make([][]int, len(trees[0]))
	posTallestInRow := make([][]int, len(trees))
	retVal := int(0)

	for r, row := range trees {
		tallestInRow := int(0)
		for c, height := range row {
			tallestInCol := int(0)
			// log.Printf("r: %d, c: %d, height: %d", r, c, height)
			//find the tallest trees in the current row
			if height > tallestInRow {
				// log.Printf("making new slice for row %v height %v is greater than %v", r, height, tallestInRow)
				newslice := make([]int, 0)
				posTallestInRow[r] = append(newslice, c)
				tallestInRow = height
			} else if height == tallestInRow {
				// log.Printf("appending to slice row %v  height %v is equal to %v", r, height, tallestInRow)
				posTallestInRow[r] = append(posTallestInRow[r], c)
			}
			//find the tallest tree in the current column
			//we'll only need to to do this once per column
			if r == 0 {
				for r2, _ := range trees {
					if trees[r2][c] > tallestInCol {
						// log.Printf("making new slice for column %v height %v is greater than %v", r2, trees[r2][c], tallestInCol)
						newslice := make([]int, 0)
						posTallestInCol[c] = append(newslice, r2)
						tallestInCol = trees[r2][c]
					} else if trees[r2][c] == tallestInCol {
						// log.Printf("appending to slice for column %v  height %v is equal to %v", r2, trees[r2][c], tallestInCol)
						posTallestInCol[c] = append(posTallestInCol[c], r2)
					}
				}
			}
		}
	}
	log.Printf("Tallest in Col: %v", posTallestInCol)
	log.Printf("Tallest in Row: %v", posTallestInRow)

	for r, row := range trees {
		if r == 0 || r == len(trees)-1 {
			//if the tree is in the first or last row it's visible no matter what
			retVal += len(row)
			log.Printf("row %v is visible because it's the first or last row", r)
			log.Printf("retVal is now %v", retVal)
			continue
		}
		//now we need to loop over the middle rows
		//and find the visible trees
		//if the tree is in the first or last column it's visible
		for c, _ := range row {
			if c == 0 || c == len(row)-1 {
				retVal++
				log.Printf("adding row:%v column: %v is visible because it's the first or last column in the row", r, c)
				log.Printf("retVal is now %v", retVal)
				continue
			}
			//check if the tree is outside of the tallestTrees in the row
			//if it's not the first or last item in the row
			minColVal := ix.MinSlice(posTallestInCol[r])
			maxColVal := ix.MaxSlice(posTallestInCol[r])
			minRowVal := ix.MinSlice(posTallestInRow[r])
			maxRowVal := ix.MaxSlice(posTallestInRow[r])
			if c <= minRowVal {
				retVal++
				log.Printf("adding row:%v column: %v is visible because it's before the first tallest tree", r, c)
				log.Printf("retVal is now %v", retVal)
        continue
			} 
      if c >= maxRowVal {
					retVal++
				log.Printf("adding row:%v column: %v is visible because it's after the last tallest tree", r, c)
				log.Printf("retVal is now %v", retVal)
          continue

			}
			if c <= minColVal {
				retVal++
				log.Printf("adding column: %v row:%v is visible because it's before the first tallest tree", r, c)
				log.Printf("retVal is now %v", retVal)
        continue
			} 
      if c >= maxColVal {
					retVal++
				log.Printf("adding column: %v row:%v  is visible because it's before the first tallest tree", r, c)
				log.Printf("retVal is now %v", retVal)
          continue

			}


			if c <= minRowVal {
				retVal++
				log.Printf("adding row:%v column: %v is visible because it's before the first tallest tree", r, c)
				log.Printf("retVal is now %v", retVal)
			} else if c >= maxRowVal {
				{
					retVal++
				}
			}
			if r == 0 {
				for r2, _ := range row {
					for _, index := range posTallestInCol[r2] {
						if r2 <= index && r2 != 0 {
							retVal++
						} else if r2 > index && r2 != len(row)-1 {
							retVal++
						}
					}
				}
			}
		}

	}
	return retVal
}
