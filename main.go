package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mhuggins7278/AdventOfCode/twozerotwotwo"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	twozerotwotwo.Day1()
}
