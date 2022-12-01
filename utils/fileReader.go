package utils

import (
	"bufio"
	"log"
	"os"
)

func GetFileScanner(filePath string) *bufio.Scanner {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

  scanner := bufio.NewScanner(file)
  return scanner 
}

