package main

import (
	// "fmt"
	"log"
	"os"
	// "regexp"
	// "sort"
	// "strconv"
	// "strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	"golang.org/x/exp/slices"
)

func Day6() {
	input, _ := os.ReadFile("./data/day6.txt")

	// log.Printf("Part 1 input: %v", input)

	part1 := Day6Part1(input, 4)
	part2 := Day6Part1(input, 14)
	// part2 := Day6Part2(input)
	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: %v", part2)
	// log.Printf("Part 2: %v", part2)

}

func Day6Part1(str []uint8, numChars int) int {
	values := make([]uint8, 0)
	// log.Printf("Part 1 input: %v", str)
	for i, char := range str {
		// log.Printf("Part1 values %v %v", values, len(values))
		// log.Printf("char: %v, index: %v", reflect.TypeOf(char), i)
		// log.Printf("slice contains %v", slices.Contains(values, char))
		if slices.Contains(values, char) == true {
			indexToRemoveUpTo := slices.Index(values, char)
			// log.Printf("removing up to %v", indexToRemoveUpTo)
			values = slices.Delete(values, 0, indexToRemoveUpTo + 1)
			values = append(values, char)
		} else {
			values = append(values, char)
		}
		if len(values) == numChars {
			// log.Printf("Part1 answer %v", i+1)
			// log.Printf("Part1 values %v", values)
			return i + 1
		}
	}
	return 0
}
