package main

import (
	"fmt"
	"log"

	// "sync"
	"os"
	// "regexp"
	"math"
	// "github.com/k0kubun/pp"
	"strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)

func Day10() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day10.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	// part1 := Day9Part1(input, 2)
	// log.Printf("Part 1 answer is  %v", part1)
	part1 := Day10Part1(input)
	log.Printf("Part 1 answer is  %v", part1)
	part2 := Day10Part2(input)
  for _, line := range part2 {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}

}

func Day10Part1(moves []string) int {
	registerValue := 1
	computedValues := []int{}
	cycleToCheck := 20
	currCycle := 0
	retVal := 0

	for _, move := range moves {
		moveParts := strings.Split(move, " ")
		// If this is a noop just increment the cycle
		currCycle++
		// pp.Printf("Start Itteration: %v, Cycle: %v Register Value: %v \n", i, currCycle, registerValue)
		if currCycle == cycleToCheck {
			// log.Printf("Computing new value: Register value at cycle %v is %v", currCycle, registerValue)
			computedValues = append(computedValues, registerValue*currCycle)
			cycleToCheck = cycleToCheck + 40

		}
		if len(moveParts) == 2 {
			currCycle++
			if currCycle == cycleToCheck {
				// log.Printf("Computing new value: Register value at cycle %v is %v", currCycle, registerValue)
				computedValues = append(computedValues, registerValue*currCycle)
				cycleToCheck = cycleToCheck + 40

			}
			inc, _ := strconv.Atoi(moveParts[1])
			registerValue += inc
			// pp.Printf("In If Itteration: %v, Cycle: %v Register Value: %v \n", i, currCycle, registerValue)
		}
		// pp.Printf("End Itteration: %v, Cycle: %v Register Value: %v \n", i, currCycle, registerValue)

	}
	// log.Printf("computedValues: %v", computedValues)

	for _, a := range computedValues {
		retVal = retVal + a
	}
	return retVal
}

func getPixel(currCol int, spritPosition int, spriteWidth int) string {
	if math.Abs(float64(currCol-spritPosition)) <= math.Abs(float64(spriteWidth/2)) {
		return "#"
	} else {
		return " "
	}
}

func Day10Part2(input []string) [6][40]string {
	spriteWidth := 3
	spritPosition := 1
	currCycle := 1
	var output [6][40]string
	for _, instruction := range input {
		currCol := (currCycle - 1) % 40
		currRow := (currCycle - 1) / 40
		output[currRow][currCol] = getPixel(currCol, spritPosition, spriteWidth)
		currCycle++
		instParts := strings.Split(instruction, " ")
		if len(instParts) == 2 {
			currCol = (currCycle - 1) % 40
			currRow = (currCycle - 1) / 40
			output[currRow][currCol] = getPixel(currCol, spritPosition, spriteWidth)
			val, _ := strconv.Atoi(instParts[1])
			spritPosition += val
			currCycle++
		}

	}

	return output

}
