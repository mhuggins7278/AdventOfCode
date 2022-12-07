package main

import (
	// "fmt"
	"log"
	"os"
	// "regexp"
	// "sort"
	"strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
)

type node struct {
	Size int
	Type string
}

func newNode(size int, nodeType string) node {
	n := node{
		Size: size,
		Type: nodeType,
	}
	return n
}

func (n node) incrSize(val int) node {
	n.Size += val
	return n
}

type Nodes map[string]node

type File struct {
	name string
	size int
}

func buildFileSystem(lines []string) Nodes {
	nodes := map[string]node{}
	currPath := make([]string, 0)
	for _, line := range lines {
		// log.Println(line)
		l := strings.Split(line, " ")
		if l[1] == "cd" {
			if l[2] == ".." {
				// log.Printf("Removing last item from currPath%v", currPath)
				currPath = currPath[:len(currPath)-1]
				continue
			} else {
				currPath = append(currPath, l[2])
			}
			// log.Printf("Current Path %v", currPath)
			if _, exists := nodes[l[2]]; exists {
				// log.Printf("Node Exists %v", l)
			} else {
				node := newNode(0, "dir")
				nodes[l[2]] = node
				// log.Printf("Node Doesn't Exists %v", l)
			}
		}

		size, err := strconv.Atoi(l[0])
		if err == nil {
			node := newNode(size, "file")
			nodes[l[1]] = node
			for _, node := range currPath {
				nodes[node] = nodes[node].incrSize(size)
			}
		}
	}

	log.Printf("nodes %v", nodes)
	return nodes
}
func Day7 (){
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day7.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")


	part1 := Day7Part1(input)
	// part2 := Day6Part1(input, 14)
	// part2 := Day6Part2(input)
log.Printf("Day 7 Part 1: %v", part1)
	// log.Printf("Part 2: %v", part2)
	// log.Printf("Part 2: %v", part2)
} 

func Day7Part1(lines []string) int {
	totalSize := 0
	fileSystem := buildFileSystem(lines)
	for _, node := range fileSystem {
		if node.Type == "dir" && node.Size <= 100000 {
			totalSize += node.Size
		}
	}
	return totalSize
}
