package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func executionTimeTracker(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "Timer")

	var inputLines []int
	var increasedCounter int

	input, _ := os.Open("input.txt")
	defer input.Close()
	
	scanner := bufio.NewScanner(input)
	
    for scanner.Scan() {
		scanText := scanner.Text()
		value, _ := strconv.Atoi(scanText)
		inputLines = append(inputLines, value)
    }

	for index := 0; index < len(inputLines)-3; index++ {

		firstWindow := inputLines[index] + inputLines[index+1] + inputLines[index+2]
		secondWindow := inputLines[index+1] + inputLines[index+2] + inputLines[index+3]

		if firstWindow < secondWindow {
			increasedCounter++
		}
	}
	fmt.Println("Finished! Answer is:", increasedCounter)
}