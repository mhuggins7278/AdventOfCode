package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mhuggins7278/AdventOfCode/twozerotwotwo"
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
	// twozerotwotwo.Day1()
	// twozerotwotwo.Day2()
	// twozerotwotwo.Day3()
  // twozerotwotwo.Day4()
  twozerotwotwo.Day5()
	utils.PrintMemUsage()
	elapsed := time.Since(start)
	log.Printf("App took %s", elapsed)

}
