package twozerotwotwo

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/mhuggins7278/AdventOfCode/utils"
)

func Day1() {
  utils.PrintMemUsage()
	scanner := utils.GetFileScanner("./2022/data/day1.txt")
	currCalories := 0
	elves := make([]int, 0)

	for scanner.Scan() {
		currValue, err := strconv.ParseInt(scanner.Text(), 10, 0)
    //parse int will error on an empty string using this as
    // a signal that we've hit the last entry for an elf
    // lets add the currCalories to the array and reset the value
		if err != nil {
			elves = append(elves, currCalories)
			currCalories = 0
		}
		currCalories += int(currValue)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Printf("Part 1 %v \n", elves[0])
	fmt.Printf("Part 2 %v \n", elves[0]+elves[1]+elves[2])
  utils.PrintMemUsage()
}


