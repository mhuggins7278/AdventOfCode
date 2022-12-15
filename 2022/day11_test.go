package main

import (
	// "fmt"
	"log"
	"strconv"
	"strings"

	// "strconv"
	"github.com/k0kubun/pp"
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
		newMonkey := Monkey{}
		for _, line := range strings.Split(monkey, "\n") {
			log.Printf("line is %v", line)
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
        operation := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
        newMonkey.operation.operator = operation[3]
        newMonkey.operation.value, _ = strconv.Atoi(operation[4])
      }
		}
		monkeyMap[newMonkey.id] = newMonkey
	}
	pp.Printf("Monkey map is %v", monkeyMap)

	// answer := Day10Part1(lines)

	// log.Printf("Day10 part 1 test %v",answer)

	if true != false {
		t.Fatalf(`Day 10 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 10 Part 1 answer is correct: %v", answer)
	}
}
