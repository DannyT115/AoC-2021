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

func main() {
	defer executionTimeTracker(time.Now(), "Timer")

	var numOfFish = make([]int, 9)
	sumOfAllFish := 0

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, i := range strings.Split(scanner.Text(), ",") {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			numOfFish[j]++
		}
	}
	
	for i := 0; i < 256; i++ {
		numOfFish = checkFish(numOfFish)
	}
	for _, fishCount := range(numOfFish){
		sumOfAllFish += fishCount
	}
	fmt.Println("Finished! The Answer is:", (sumOfAllFish))
}

func checkFish(numOfFish []int) []int {
	var allFish = make([]int, 9)
	for index := 1; index < 9; index++ {
		allFish[index-1] = numOfFish[index]
	}
	
	allFish[6] += numOfFish[0]
	allFish[8] += numOfFish[0]
	return allFish
}