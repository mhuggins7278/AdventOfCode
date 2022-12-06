package main

import (
	// "fmt"
	"log"
	// "strconv"
	"strings"

	"github.com/mhuggins7278/AdventOfCode/utils"
)

type Choice struct {
	value  int
	wins   string
	looses string
}
type Round struct {
	me   string
	them string
}

var m = map[string]Choice{
	"Rock":     {value: 1, wins: "Scissors", looses: "Paper"},
	"Paper":    {value: 2, wins: "Rock", looses: "Scissors"},
	"Scissors": {value: 3, wins: "Paper", looses: "Rock"},
}

func Day2() {
	scanner := utils.GetFileScanner("./2022/data/day2.txt")
	rounds := make([]Round, 0)
	for scanner.Scan() {
		var currRow = strings.Split(scanner.Text(), " ")
    me := currRow[1]
		them:= currRow[0]
    rounds = append(rounds, Round{me: me, them: them})
		//   currRound := Round{
		//     "me": me,
		//     "them": them,

		// }
		// rounds = append(rounds,currRound)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	part1(rounds)
	part2(rounds)

}

func part1(rounds []Round) {
	// A = Rock
	// B = Paper
	//C = Scissors
	// X = Rock
	// Y = Paper
	//Z = Scissors

	var choices = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	scores := make([]int, 0)

	for _, round := range rounds {
    
    them := choices[round.them]
    me := choices[round.me]
		//the round is a tie
		if them == me {
			scores = append(scores, m[me].value+3)
			continue
		}
		if me == "Rock" {
			if them == "Paper" {
				scores = append(scores, m[me].value)
			} else {
				scores = append(scores, m[me].value+6)
			}
			continue
		}
		if me == "Paper" {
			if them == "Scissors" {
				scores = append(scores, m[me].value)
			} else {
				scores = append(scores, m[me].value+6)
			}
			continue
		}
		if me == "Scissors" {
			if them == "Rock" {
				scores = append(scores, m[me].value)
			} else {
				scores = append(scores, m[me].value+6)
			}
			continue
		}

	}
	arrSum := 0

	for _, a := range scores {
		arrSum = arrSum + a
	}
	log.Println("Part 1: ", arrSum)

}

func part2(rounds []Round) {

	var scores = make([]int, 0)
	var choices = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	var outcomes = map[string]string{
		"X": "looses",
		"Y": "draw",
		"Z": "wins",
	}

	for _, round := range rounds {
    them := m[choices[round.them]]
    desiredOutcome := outcomes[round.me]
		//the round is a tie
		if desiredOutcome == "draw" {
			scores = append(scores, 3+them.value)
			continue
		}
	//     // i need to loose
			if desiredOutcome == "looses" {
        scores = append(scores, m[them.wins].value)
				continue
			}
			if desiredOutcome == "wins" {
        scores = append(scores, m[them.looses].value + 6)
				continue
			}
    }

		arrSum := 0

		for _, a := range scores {
			arrSum = arrSum + a
		}
		log.Println("Part 2: ", arrSum)

}
