package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func executionTimeTracker(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main(){
	defer executionTimeTracker(time.Now(), "Timer")

	horizontalVal := 0
	depthVal := 0

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

    for scanner.Scan() {
		direction := strings.Split(scanner.Text(), " ")
		directionValue, _ := strconv.Atoi(direction[1])
		switch direction[0] {
		case "forward":
			horizontalVal += directionValue
		case "down":
			depthVal += directionValue
		case "up":
			depthVal -= directionValue
		}
	}
	fmt.Println(horizontalVal * depthVal)
}