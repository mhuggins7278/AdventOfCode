package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mhuggins7278/AdventOfCode/utils"
)

func main() {
	start := time.Now()
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	utils.PrintMemUsage()
	// main.Day1()
	// main.Day2()
	// main.Day3()
  // main.Day4()
  // Day6()
  // Day7()
  Day8()
	utils.PrintMemUsage()
	elapsed := time.Since(start)
	log.Printf("App took %s", elapsed)

}
