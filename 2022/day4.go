package twozerotwotwo

import (
	"log"
	"strconv"
	"strings"

	"github.com/mhuggins7278/AdventOfCode/utils"
)

func Day4() {
	scanner := utils.GetFileScanner("./2022/data/day4.txt")
	input := make([]string, 0)

	for scanner.Scan() {
		var row = scanner.Text()
		input = append(input, row)
	}

	var part1 = Day4Part1(input)
	log.Printf("Part 1: %v", part1)
	var part2 = Day4Part2(input)
	log.Printf("Part 2: %v", part2)
}

type Group = []int

func convertToGroup(input string) Group {
		var group = strings.Split(input, "-")
  // log.Println(group)
  b := make([]int, 2)
   for i, v := range group {
        b[i], _ = strconv.Atoi(v)
    }
  return b

}

func Day4Part1(input []string) int {
	var groupsWithOverLap int = 0
	for _, groupPair := range input {
		var currRow = strings.Split(groupPair, ",")
		var firstPair Group = convertToGroup(currRow[0])
		var secondPair Group = convertToGroup(currRow[1])
		if firstPair[0] <= secondPair[0] && firstPair[1] >= secondPair[1]{ 
      log.Printf("firstPair: %v secondPair: %v", firstPair, secondPair)
				groupsWithOverLap++
				continue
		}
		if secondPair[0] <= firstPair[0] && secondPair[1] >= firstPair[1] {
      log.Printf("secondPair: %v firstPair: %v", secondPair, firstPair)
				groupsWithOverLap++
		}
	}
	return groupsWithOverLap

}
func Day4Part2(input []string) int {
	var groupsWithOverLap int = 0
	for _, groupPair := range input {
		var currRow = strings.Split(groupPair, ",")
		var firstPair Group = convertToGroup(currRow[0])
		var secondPair Group = convertToGroup(currRow[1])
    //second pair is contained in first pair
		if firstPair[0] <= secondPair[0] && firstPair[1] >= secondPair[1]{ 
      // log.Printf("firstPair: %v secondPair: %v", firstPair, secondPair)
				groupsWithOverLap++
				continue
		}
    //first pair is contained in second pair
		if secondPair[0] <= firstPair[0] && secondPair[1] >= firstPair[1] {
      // log.Printf("secondPair: %v firstPair: %v", secondPair, firstPair)
				groupsWithOverLap++
        continue
		}
    //first pair overlaps second pair
		if firstPair[0] <= secondPair[0] && firstPair[1] >= secondPair[0]  { 
      // log.Printf("firstPair: %v secondPair: %v", firstPair, secondPair)
				groupsWithOverLap++
				continue
		}
		if secondPair[0] <= firstPair[0] && secondPair[1] >= firstPair[0] {
      // log.Printf("secondPair: %v firstPair: %v", secondPair, firstPair)
				groupsWithOverLap++
        continue
    }
	}
	return groupsWithOverLap

}
