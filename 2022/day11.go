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
	"golang.org/x/exp/slices"
	// "github.com/adam-lavrik/go-imath/ix"
	"github.com/Knetic/govaluate"
)

type Operator struct {
	operator string
	value    int
}

type Monkey struct {
	id             int
	items          []int
	expression     string
	test           int
	pass           int
	fail           int
	inspectedItems int
}

type MonkeyMap map[int]*Monkey

func (m *Monkey) Turn(monkeyMap MonkeyMap) {
	for _, item := range m.items {
    m.inspectedItems++
		// log.Printf("monkey is %v", m.id)
		exp, _ := govaluate.NewEvaluableExpression(m.expression)
		parameters := make(map[string]interface{}, 8)
		parameters["old"] = item
		// log.Printf("exp is %v", exp)
		worryLevel, _ := exp.Evaluate(parameters)
    worryLevelfloat := float64(worryLevel.(float64))
    // log.Printf("worryLevel is %v", worryLevelfloat)
    worryLevelInt := int(math.Floor(worryLevelfloat))
		// log.Printf("worryLevel is %v", worryLevelInt)
		if worryLevelInt%m.test == 0 {
			transferMonkey := monkeyMap[m.pass]
      transferItems := make([]int, 0)
      transferItems = append(transferItems, transferMonkey.items...)
			// log.Printf("transfering item %v to monkey %v", worryLevelInt, transferMonkey.id)
			transferMonkey.items = append(transferItems, worryLevelInt)
		} else {
			transferMonkey := monkeyMap[m.fail]
      transferItems := make([]int, 0)
      transferItems = append(transferItems, transferMonkey.items...)
			// log.Printf("transfering item %v to monkey %v", worryLevelInt, transferMonkey.id)
			transferMonkey.items = append(transferItems, worryLevelInt) 
		}
		// log.Printf("items before removing %v", m.items)
		foo := make([]int, 0)
		m.items = append(foo, m.items[1:]...)
		// log.Printf("items after removing %v", m.items)
		// pp.Printf("monkey after turn %v", m)
	}
}

func NewMonkey(data string) *Monkey {
	newMonkey := Monkey{}
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(strings.Trim(line, " "), "Monkey") {
			newMonkey.id, _ = strconv.Atoi(strings.Trim(strings.Split(line, " ")[1], ":"))
		}
		if strings.HasPrefix(strings.Trim(line, " "), "Starting items:") {
			newMonkey.items = make([]int, 0)
			items := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), ", ")
			for _, item := range items {
				i, _ := strconv.Atoi(item)
				newMonkey.items = append(newMonkey.items, i)
			}
		}
		if strings.HasPrefix(strings.Trim(line, " "), "Operation:") {
			operation := strings.Trim(strings.Split(line, "=")[1], " ")
			// log.Printf("operation is %v", operation)
			newMonkey.expression = fmt.Sprintf("(%v)", operation)
		}
		if strings.HasPrefix(strings.Trim(line, " "), "Test:") {
			operation := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			newMonkey.test, _ = strconv.Atoi(operation[2])
		}
		if strings.HasPrefix(strings.Trim(line, " "), "If true:") {
			operation := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			newMonkey.pass, _ = strconv.Atoi(operation[3])
		}
		if strings.HasPrefix(strings.Trim(line, " "), "If false:") {
			operation := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			newMonkey.fail, _ = strconv.Atoi(operation[3])
		}
	}

	return &newMonkey
}

func Day11() {
	inputFile, _ := os.ReadFile("./data/day11.txt")
	monkeys := strings.Split(string(inputFile), "\n\n")
  // log.Printf("monkeys is %v", monkeys)
	monkeyMap := make(MonkeyMap)
	for _, monkey := range monkeys {
		newMonkey := NewMonkey(monkey)
		// pp.Printf("Monkeyis %v", newMonkey)
		monkeyMap[newMonkey.id] = newMonkey
	}

	part1 := Day11Part1(monkeyMap)
	log.Printf("Part 1 answer is  %v", part1)

}

func Day11Part1(monkeyMap MonkeyMap) int {
	retValue := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < len(monkeyMap); j++ {
			monkey := monkeyMap[j]
			// log.Printf("looping monkey is %v", monkey.id)
			monkey.Turn(monkeyMap)
		}
	}
	// pp.Printf("monkeyMap is %v", monkeyMap)
	inspectedItems := make([]int, 0)
	for _, monkey := range monkeyMap {
		inspectedItems = append(inspectedItems, monkey.inspectedItems)
	}
  slices.Sort(inspectedItems) 
  log.Printf("inspectedItems is %v", inspectedItems)
  retValue = inspectedItems[len(inspectedItems) -1] * inspectedItems[len(inspectedItems) - 2] 

	return retValue

}
