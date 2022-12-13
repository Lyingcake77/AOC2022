package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("day1.q.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	maxAmount := 0
	currentAmount := 0
	for _, line := range lines {
		if line == "" {
			if currentAmount > maxAmount {
				maxAmount = currentAmount
			}
			currentAmount = 0
		} else {
			tempResult, _ := strconv.ParseInt(line, 10, 64)
			currentAmount = currentAmount + int(tempResult)
		}

	}
	fmt.Println("max:", maxAmount)
}
