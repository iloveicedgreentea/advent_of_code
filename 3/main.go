package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/iloveicedgreentea/advent_of_code/common"
)

func main() {
	lines := common.ReadFile("input.txt")
	gamma, epsilon := calcInputs(lines)

	length := len(lines)
	// gam and eps tell us which bit was the majority per row
	gam := processNum(gamma, length)
	eps := processNum(epsilon, length)

	gamDecoded := binToDec(gam)
	epsDecoded := binToDec(eps)

	fmt.Println(gamDecoded, epsDecoded)
	fmt.Println(gamDecoded * epsDecoded)
}

// convert binary to decimal
func binToDec(binary string) int64 {
	dec, _ := strconv.ParseInt(binary, 2, 64)

	return dec
}

func processNum(input [12]int, length int) string {
	finalArr := make([]string, 12)

	// loop over input and check if the majority was 1 or 0
	for index, val := range input {
		if float64(val)/float64(length) < 0.5 {
			finalArr[index] = "0"
		} else {
			finalArr[index] = "1"
		}
	}

	// convert the final array into one binary string
	return strings.Join(finalArr, "")
}

// Each bit in the gamma rate can be determined by finding the most common bit in the
// corresponding position of all numbers in the diagnostic report.
func calcInputs(input []string) ([12]int, [12]int) {
	// input is 12 bits

	// index | Bits
	// ---------------------------------
	// 1     | 0 1 1 0 1 0 0 1 0 1 1 0
	// 2     | 1 0 1 1 1 0 1 0 0 1 1 0

	gamma := [12]int{}
	epsilon := [12]int{}

	for _, line := range input {
		// itter over each bit in line, add them to arr
		for index, bit := range line {
			// bit is a rune, 49 is ascii for 1
			if bit == 49 {
				gamma[index] += 1
			} else {
				epsilon[index] += 1
			}
		}
	}

	return gamma, epsilon
}
