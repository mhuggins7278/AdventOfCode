package day1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mhuggins7278/AdventOfCode/utils"
)

func Part1() {
	scanner := utils.GetFileScanner("./2022/data/day1.txt")
	mostCalories := 0
	currCalories := 0

	for scanner.Scan() {
		currValue, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			if mostCalories < currCalories {
				mostCalories = currCalories
			}
			currCalories = 0

		}
		currCalories += int(currValue)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mostCalories)
}

func Part2() {
	scanner := utils.GetFileScanner("./2022/data/day1.txt")
	mostCalories := 0
	secondMostCalories := 0
	thirdMostCalories := 0
	currCalories := 0

	for scanner.Scan() {
		currValue, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			if mostCalories < currCalories {
				mostCalories = currCalories
			} else if secondMostCalories < currCalories {
				secondMostCalories = currCalories
			} else if thirdMostCalories < currCalories {
				thirdMostCalories = currCalories
			}
			currCalories = 0
		}

		currCalories += int(currValue)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mostCalories + secondMostCalories + thirdMostCalories)
}
