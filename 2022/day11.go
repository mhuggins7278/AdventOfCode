package main

import (
	// "fmt"
	"log"

	// "sync"
	"os"
	// "regexp"
	// "math"
	// "github.com/k0kubun/pp"
	// "strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
)

type Operator struct {
	operator string
	value    int
}

type Monkey struct {
	id        int
	items     []int
	operation Operator
	test      string
	pass      int
	fail      int
}

type MonkeyMap map[int]Monkey

func (m *Monkey) Turn() {
	for _, item := range m.items {
		log.Printf("Monkey %v has %v items", m, item)
	}
}

func Day11() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day11.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	log.Printf("Day 11 Part 1 input %v", input)
	// part1 := Day9Part1(input, 2)
	// log.Printf("Part 1 answer is  %v", part1)

}

func Day11Part1(monkeys MonkeyMap) int {
	retValue := 0

	return retValue

}
