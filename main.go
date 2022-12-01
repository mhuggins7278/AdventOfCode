package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mhuggins7278/AdventOfCode/twozerotwotwo"
	"github.com/mhuggins7278/AdventOfCode/utils"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
utils.PrintMemUsage()
	twozerotwotwo.Day1()
  twozerotwotwo.Day2()
utils.PrintMemUsage()

}
