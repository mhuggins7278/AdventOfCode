package twozerotwotwo

import (
	"log"

	"github.com/mhuggins7278/AdventOfCode/utils"
	"golang.org/x/exp/maps"
)

type ruckSack struct {
	side1 map[string]int
	side2 map[string]int
}

func getItemValue(item byte) int {
	if item >= 97 {
		return int(item % 32)

	}
	return int(item%32) + 26
}

func splitSackIntoSections(s string) ruckSack {
	var stringLen = len(s)
	var side1 = map[string]int{}
	var side2 = map[string]int{}
	for i, c := range s {
		if i == 0 || i < stringLen/2 {
			side1[string(c)] = getItemValue(byte(c))
			continue

		}
		side2[string(c)] = getItemValue(byte(c))

	}
	return ruckSack{side1: side1, side2: side2}
}

func chunkSack(slice []ruckSack) [][]ruckSack {
	var chunks [][]ruckSack
	for i := 0; i < len(slice); i += 3 {
		end := i + 3

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func Day3() {
	scanner := utils.GetFileScanner("./2022/data/day3.txt")
	ruckSacks := make([]ruckSack, 0)

	for scanner.Scan() {
		var sack = scanner.Text()
		ruckSacks = append(ruckSacks, splitSackIntoSections(sack))
	}

	three1(ruckSacks)
	three2(ruckSacks)
}

func three1(ruckSacks []ruckSack) {
	var dupes []int
	arrSum := 0
	for _, sack := range ruckSacks {
		for k, v := range sack.side1 {
			if _, exists := sack.side2[k]; exists {
				dupes = append(dupes, v)
				break
			}
		}
	}
	for _, a := range dupes {
		arrSum = arrSum + a
	}
	log.Printf("Part 1 %v \n", arrSum)
}

func three2(ruckSacks []ruckSack) {
	var dupes []int
	arrSum := 0
	//break the ruckSacks array int groups of 3 items
	var chunks = chunkSack(ruckSacks)
	for _, chunk := range chunks {
		for _, sack := range chunk {
			maps.Copy(sack.side1, sack.side2)
		}
		for k, v := range chunk[0].side1 {
			if _, exists := chunk[1].side1[k]; exists {

				if _, exists := chunk[2].side1[k]; exists {
					dupes = append(dupes, v)
					break

				}
			}

		}
	}

	//find the item that is common in each group of 3
	for _, a := range dupes {
		arrSum = arrSum + a
	}
	log.Printf("Part 2 %v \n", arrSum)

}
