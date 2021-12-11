package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func executionTimeTracker(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "Timer")

	var inputLines []string //Holds the values in the text file
	var increasedCounter int //Counter for increases

	input, _ := os.Open("input.txt")
	defer input.Close()
	
	scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        inputLines = append(inputLines, scanner.Text())
    }

	for index, eachLine := range inputLines {
		if index+1 < len(inputLines){
			if eachLine < inputLines[index+1] {
				increasedCounter++
			}
		}else{
			fmt.Println("Finished", increasedCounter)
		}
	}
}