package main

import (
	"fmt"
	"github.com/iloveicedgreentea/advent_of_code/common"
)

func main() {
	lines := common.ReadFileInt("input.txt")

	answer1(lines)
	answer2(lines)
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
