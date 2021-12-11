package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var inputLines []string //Holds the values in the text file
var gammaRate string
var epsilonRate string

func executionTimeTracker(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "Timer")

	input, _ := os.Open("input.txt")
	defer input.Close()
	
	scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        inputLines = append(inputLines, scanner.Text())
    }

	for index := range inputLines[0] {
		bitValue := checkCount(inputLines, index)

		switch bitValue {
		case 0:
			gammaRate += "0"
			epsilonRate += "1"
		case 1:
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	finalGamma, _ := strconv.ParseUint(gammaRate, 2, 64)
	finalEpsilon, _ := strconv.ParseUint(epsilonRate, 2, 64)
	fmt.Println(finalGamma*finalEpsilon)
}

func checkCount(inputLines[] string, index int) (bitValue int){
	zeroVal := 0
	oneVal := 1

	for _, bitString := range inputLines {
		if bitString[index] == '0' {
			zeroVal++
		} else {
			oneVal++
		}
	}
	if zeroVal > oneVal {
		return 0
	}
	return 1
}