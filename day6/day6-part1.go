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

	var numOfFish []int

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, i := range strings.Split(scanner.Text(), ",") {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			numOfFish = append(numOfFish, j)
		}
	}
	
	for i := 0; i < 80; i++ {
		numOfFish = checkFish(numOfFish)
	}
	fmt.Println("Finished! The Answer is:", len(numOfFish))
}

func checkFish(numOfFish []int) []int {
	for index := range numOfFish {
		if numOfFish[index] !=0 {
			numOfFish[index] -= 1
		}else{
			numOfFish[index] += 6
			numOfFish = append(numOfFish, 8)
		}
	}
	return numOfFish
}