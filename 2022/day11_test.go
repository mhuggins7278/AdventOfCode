package main

import (
	// "fmt"
	"log"
	// "strconv"
	"strings"

	// "strconv"
	// "github.com/k0kubun/pp"
	"golang.org/x/exp/slices"
	"testing"
)

var day11test = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func TestDay11Part1(t *testing.T) {
	monkeys := strings.Split(day11test, "\n\n")
	answer := 0
	// pp.Printf("Monkeys is %v", monkeys)
	monkeyMap := make(MonkeyMap)
	for _, monkey := range monkeys {
		newMonkey := NewMonkey(monkey)
		// pp.Printf("Monkeyis %v", newMonkey)
		monkeyMap[newMonkey.id] = newMonkey
	}
	for i := 0; i < 20; i++ {
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
  answer = inspectedItems[2] * inspectedItems[3] 

	// answer := Day10Part1(lines)

	// log.Printf("Day10 part 1 test %v",answer)

	if answer != 10605 {
		t.Fatalf(`Day 10 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 10 Part 1 answer is correct: %v", answer)
	}
}
