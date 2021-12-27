package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	dataFile, err := os.Open("input.txt")
	checkErr(err)

	defer dataFile.Close()

	// create scanner to read file
	scanner := bufio.NewScanner(dataFile)

	// store the lines as int
	var lines []int
	// for each line, read and add it to int array
	for scanner.Scan() {
		convLine, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		lines = append(lines, convLine)
	}

	answer1(lines)
	answer2(lines)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func answer1(lines []int) {
	// times the value increased
	timesIncreased := 0

	// get the current and next line
	// see if its increased
	for index, line := range lines {
		if index < len(lines)-1 {
			nextLine := lines[index+1]
			// check if LT because EOF has no next line
			if line < nextLine {
				timesIncreased += 1
			}
		}
	}
	fmt.Println("Times without sliding window:")
	fmt.Println(timesIncreased)
}

func answer2(lines []int) {
	timesIncreased := 0

	for index, line := range lines {
		if index < len(lines)-3 {
			slidingWindowOneTotal := line + lines[index+1] + lines[index+2]
			slidingWindowTwoTotal := lines[index+1] + lines[index+2] + lines[index+3]

			// check if LT because EOF has no next line
			if slidingWindowOneTotal < slidingWindowTwoTotal {
				timesIncreased += 1
			}
		}
	}
	fmt.Println("Times with sliding windows")
	fmt.Println(timesIncreased)
}
