package twozerotwotwo

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	// "reflect"
	"github.com/mhuggins7278/AdventOfCode/utils"
)

func Day5() {
	scanner := utils.GetFileScanner("./2022/data/day4.txt")
	input := make([]string, 0)

	for scanner.Scan() {
		var row = scanner.Text()
		input = append(input, row)
	}

	// var part1 = Day5Part1(input)
	// log.Printf("Part 1: %v", part1)
	// var part2 = Day4Part2(input)
	// log.Printf("Part 2: %v", part2)
}

func buildStacks(input string) map[string][]string {
	space := regexp.MustCompile(`\s+`)
	stacks := make(map[string][]string)
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	for _, row := range rows {
		log.Printf("row: %v", row)
		row = space.ReplaceAllString(row, " ")
    row = strings.ReplaceAll(row, "[", "")
    row = strings.ReplaceAll(row, "]", "")
		row := strings.Split(row, " ")
		log.Printf("row: %v", row)
		for i, _ := range row {
			char := row[i]
			var key = strconv.Itoa(i + 1)
			if char == "" {
				continue
			}
			if _, exists := stacks[key]; exists {
				stacks[key] = append(stacks[key], char)
			} else {
				stacks[key] = make([]string, 0)
				stacks[key] = append(stacks[key], char)
			}
		}
	}
	return stacks
}
func reverseItemsToMove(stack []string) []string {
	reversed := make([]string, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		reversed = append(reversed, stack[i])
	}
	return reversed
}

func Day5Part1(input []string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)
	space := regexp.MustCompile(`\s+`)
	stackString := input[0]
	stacks := buildStacks(stackString)
	log.Printf("Initial stacks: %v", stacks)
	moves := strings.Split(input[1], "\n")
	for _, move := range moves {
		fmt.Println(move)
		log.Printf("stacks: %v", stacks)
		move := strings.Split(space.ReplaceAllString(nonAlphanumericRegex.ReplaceAllString(move, ""), ""), "")
		quantity, _ := strconv.Atoi(move[0])
		from := move[1]
		to := move[2]
		itemsToMove := make([]string, 0)
		newSource := make([]string, 0)
		itemsToMove = reverseItemsToMove(append(itemsToMove, stacks[from][0:quantity]...))
		newSource = append(newSource, stacks[from][quantity:]...)
		newTarget := append(itemsToMove, stacks[to]...)
		stacks[to] = newTarget
		stacks[from] = newSource
	}
	answer := make([]string, 0) 
  for _, stack := range stacks {
    answer = append(answer, stack[0])
  } 
	log.Printf("answer: %v", answer)
	return strings.Join(answer, "")
}
