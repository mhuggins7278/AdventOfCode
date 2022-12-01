package main

import (
	"fmt"
	"log"
	"os"

	day1 "github.com/mhuggins7278/AdventOfCode/twozerotwotwo"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	day1.Part1()
	day1.Part2()
}
