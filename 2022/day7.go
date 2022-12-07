package main

import (
	// "fmt"
	"log"
	// "os"
	// "regexp"
	// "sort"
	// "strconv"
	"strings"
	// "reflect"
	// "github.com/mhuggins7278/AdventOfCode/utils"
	// "golang.org/x/exp/slices"
)

type Folder struct {
	Name    string
	Files   *File
	Folders *Folder
}

type File struct {
	name string
	size int
}

func buildFileSystem(lines []string) Folder {
	root := Folder{}
	for _, line := range lines {
		log.Println(line)
		l := strings.Split(line, " ")
		switch l[0] {
    case "$":
		}
	}
	return root
}

func Day7Part1(lines []string) int {
	fileSystem := buildFileSystem(lines)
	log.Println(fileSystem)

	return 0
}
