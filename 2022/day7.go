package main

import (
	// "fmt"
	"log"
	"os"

	// "regexp"
	// "sort"
	"strconv"
	"strings"

	// "github.com/k0kubun/pp"
	// "github.com/k0kubun/pp"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	"golang.org/x/exp/slices"
)

type node struct {
	Size int
	Type string
}

func newNode(size int, nodeType string) node {
	n := &node{
		Size: size,
		Type: nodeType,
	}
	return *n
}

type Nodes map[string]node

type File struct {
	name string
	size int
}

func BuildFileSystem(lines []string) Nodes {
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
      }

      currPath = append(currPath, l[2])
			fullPath := strings.Join(currPath, "/")
			// log.Printf("Current Path %v", currPath)
			if _, exists := nodes[l[2]]; exists {
				// log.Printf("Node Exists %v", l)
			} else {
        n := node{Size:0, Type:"dir"}
				nodes[fullPath] = n
				// log.Printf("Node Doesn't Exists %v", l)
			}
		}

		size, err := strconv.Atoi(l[0])
		if err == nil {
			fullPath := strings.Join(append(currPath, l[1]), "/")
      n := node{Size:size, Type:"file"}
			if _, exists := nodes[fullPath]; exists {
			} else {
				nodes[fullPath] = n
				for i, _ := range currPath {
					// log.Printf("Index %v", i)
					nodePath := strings.Join(currPath[0:i+1], "/")
					// log.Printf("nodePath %v ", nodePath)
					un := nodes[nodePath]
					// log.Printf("Adding size %v to %v", size, un)
					un.Size += size
					nodes[nodePath] = un
				}
			}
		}
	}
	return nodes
}
func Day7() {
	input := make([]string, 0)
	inputFile, _ := os.ReadFile("./data/day7.txt")
	input = strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")
	fileSystem := BuildFileSystem(input)


	part1 := Day7Part1(fileSystem)
	part2 := Day7Part2(fileSystem)
	// part2 := Day6Part1(input, 14)
	// part2 := Day6Part2(input)
	log.Printf("Day 7 Part 1: %v", part1)
	log.Printf("Day 7 Part 2: %v", part2)
	// log.Printf("Part 2: %v", part2)
	// log.Printf("Part 2: %v", part2)
}

func Day7Part1(fileSystem Nodes) int {
	totalSize := 0
	for _, node := range fileSystem {
		if node.Type != "file" && node.Size <= 100_000 {
			totalSize += node.Size
		}
	}
	return totalSize
}

func Day7Part2(fileSystem Nodes) int   {
  potentialDirs := make([]int, 0)
  totalSize := 70_000_000
  requiredFreeSpace := 30_000_000
  currentUsedSpace := fileSystem["/"].Size
  freeSpace := totalSize - currentUsedSpace 
  minSizeToDelete := requiredFreeSpace - freeSpace
  
  for k, node := range fileSystem {
    if node.Type != "file" && node.Size >= minSizeToDelete {
      newFreeSpace := freeSpace + node.Size 
      log.Printf("Node %v:%v would result in %v free space", k,node.Size, newFreeSpace)
      potentialDirs = append(potentialDirs, node.Size)


  }
    }
  slices.Sort(potentialDirs) 
  return potentialDirs[0]
}

