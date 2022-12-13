package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	listOShit := []int{}
	for _, line := range lines {
		if line == "" {
			if currentAmount > maxAmount {
				maxAmount = currentAmount
			}
			listOShit = append(listOShit, currentAmount)
			currentAmount = 0

		} else {
			tempResult, _ := strconv.ParseInt(line, 10, 64)
			currentAmount = currentAmount + int(tempResult)
		}

	}
	sort.Ints(listOShit)
	fmt.Println("max:", maxAmount)
	fmt.Println("max 3:", listOShit[len(listOShit)-1]+listOShit[len(listOShit)-2]+listOShit[len(listOShit)-3])
}
