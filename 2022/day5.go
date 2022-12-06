package main

import (
	// "fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
)

func Day5() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./2022/data/day5.txt")
	input = strings.Split(string(inputFile), "\n\n")

	// log.Printf("Part 1 input: %v", input)

	part1 := Day5Part1(input)
	part2 := Day5Part2(input)
	log.Printf("Part 1: %v", part1)
	log.Printf("Part 2: %v", part2)
	// var part2 = Day4Part2(input)
	// log.Printf("Part 2: %v", part2)
}

func buildStacks(input string) map[string][]string {
	space := regexp.MustCompile(`\s{4}`)
	stacks := make(map[string][]string)
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	for _, row := range rows {
		row = space.ReplaceAllString(row, " ")
		row = strings.ReplaceAll(row, "[", "")
		row = strings.ReplaceAll(row, "]", "")
		row := strings.Split(row, " ")
		// log.Printf("row: %v", row)
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
	// space := regexp.MustCompile(`\s+`)
	stackString := input[0]
	// log.Printf("stackString: %v", stackString)
	stacks := buildStacks(stackString)
	// log.Printf("Initial stacks: %v", stacks)
	moves := strings.Split(input[1], "\n")
	// log.Printf("moves: %v", moves)

	for _, move := range moves {
		// log.Println(move)
		// log.Printf("stacks: %v", stacks)
		move := strings.Split(strings.Trim(nonAlphanumericRegex.ReplaceAllString(move, " "), " "), " ")
		// log.Printf("move: %v", move)
		quantity, err := strconv.Atoi(move[0])
		if err != nil {
      log.Println("End of input reached")
			break
		}
		from := move[1]
		to := move[2]
		itemsToMove := make([]string, 0)
		// log.Printf("Move %v from: %v to: %v",quantity, from, to)
		newSource := make([]string, 0)
		itemsToMove = reverseItemsToMove(append(itemsToMove, stacks[from][0:quantity]...))
		// log.Printf("itemsToMove: %v", itemsToMove)
		newSource = append(newSource, stacks[from][quantity:]...)
		newTarget := append(itemsToMove, stacks[to]...)
		stacks[to] = newTarget
		stacks[from] = newSource
	}
	// log.Printf("stacks: %v", stacks)
	answer := make([]string, 0)
	keys := make([]int, 0)
	for key, _ := range stacks {
		key, _ := strconv.Atoi(key)
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		answer = append(answer, stacks[strconv.Itoa(key)][0])
	}
	// log.Printf("answer: %v", answer)
	return strings.Join(answer, "")
}
func Day5Part2(input []string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)
	// space := regexp.MustCompile(`\s+`)
	stackString := input[0]
	// log.Printf("stackString: %v", stackString)
	stacks := buildStacks(stackString)
	// log.Printf("Initial stacks: %v", stacks)
	moves := strings.Split(input[1], "\n")
	// log.Printf("moves: %v", moves)

	for _, move := range moves {
		// log.Println(move)
		// log.Printf("stacks: %v", stacks)
		move := strings.Split(strings.Trim(nonAlphanumericRegex.ReplaceAllString(move, " "), " "), " ")
		// log.Printf("move: %v", move)
		quantity, err := strconv.Atoi(move[0])
		if err != nil {
      log.Println("End of input reached")
			break
		}
		from := move[1]
		to := move[2]
		itemsToMove := make([]string, 0)
		// log.Printf("Move %v from: %v to: %v",quantity, from, to)
		newSource := make([]string, 0)
		itemsToMove = append(itemsToMove, stacks[from][0:quantity]...)
		// log.Printf("itemsToMove: %v", itemsToMove)
		newSource = append(newSource, stacks[from][quantity:]...)
		newTarget := append(itemsToMove, stacks[to]...)
		stacks[to] = newTarget
		stacks[from] = newSource
	}
	// log.Printf("stacks: %v", stacks)
	answer := make([]string, 0)
	keys := make([]int, 0)
	for key, _ := range stacks {
		key, _ := strconv.Atoi(key)
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		answer = append(answer, stacks[strconv.Itoa(key)][0])
	}
	log.Printf("answer: %v", answer)
	return strings.Join(answer, "")
}
