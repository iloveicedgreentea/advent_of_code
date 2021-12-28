package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/iloveicedgreentea/advent_of_code/common"
)

func main() {
	lines := common.ReadFile("input.txt")
	getPositionPart1(lines)
	getPositionPart2(lines)
}

func getPositionPart1(commands []string) {
	position := 0
	depth := 0

	// split the line into command and value
	for _, line := range commands {
		commandLine := strings.Split(line, " ")
		val, err := strconv.Atoi(commandLine[1])
		common.CheckErr(err)

		switch commandLine[0] {
		case "forward":
			position += val

		case "up":
			depth -= val

		case "down":
			depth += val
		}
	}

	fmt.Println(position * depth)
}

func getPositionPart2(commands []string) {
	position := 0
	depth := 0
	aim := 0

	// split the line into command and value
	for _, line := range commands {
		commandLine := strings.Split(line, " ")
		val, err := strconv.Atoi(commandLine[1])
		common.CheckErr(err)

		switch commandLine[0] {
		case "forward":
			position += val
			depth += aim*val

		case "up":
			aim -= val

		case "down":
			aim += val
		}
	}

	fmt.Println(position * depth)
}
