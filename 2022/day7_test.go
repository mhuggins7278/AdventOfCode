package main

import (
	// "fmt"
	"log"
	"strings"
	"testing"
	// "log"
) 

var day7test = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDay7Part1(t *testing.T) {
  lines := strings.Split(day7test, "\n")
	fileSystem := BuildFileSystem(lines)
answer := Day7Part1(fileSystem)
  log.Println(answer)
	if answer != 95437 {
		t.Fatalf(`Day 7 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 7 Part 1 answer is correct: %v", answer)
	}

}

func TestDay7Part2(t *testing.T) {
  lines := strings.Split(day7test, "\n")
	fileSystem := BuildFileSystem(lines)
answer := Day7Part2(fileSystem)
  log.Println(answer)
	if answer != 24933642{
		t.Fatalf(`Day 7 Part 1 answer is incorrect: %v`, answer)
	} else {
		t.Logf("Day 7 Part 1 answer is correct: %v", answer)
	}

}
