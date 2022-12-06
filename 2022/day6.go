package twozerotwotwo

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
)

func Day6(string) {
	input, _ := os.ReadFile("./2022/data/day6.txt")

	// log.Printf("Part 1 input: %v", input)

	part1 := Day6Part1(input)
	// part2 := Day6Part2(input)
	log.Printf("Part 1: %v", part1)
	// log.Printf("Part 2: %v", part2)

}

func Day6Part1(str []byte) int {
	values := make(map[byte]int)
	log.Printf("Part 1 input: %v", str)
	for i, char := range str {
      log.Printf("Part1 values %v", values)
    log.Printf("char: %v, index: %v", char, i)
		if _, exists := values[char]; exists {
      values = make(map[byte]int)
      values[char] = i
		  continue	
		}else {
      values[char] = i
    }
    if len(values) == 4 {
      log.Printf("Part1 answer %v", i+1)
      log.Printf("Part1 values %v", values)
      return i + 1
    }
	}
	return 0
}
