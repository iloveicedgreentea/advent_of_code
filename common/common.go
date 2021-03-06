package common

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// checks an error
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// RemoveFromArrayStr safely removes an element
func RemoveFromArrayStr(s []string, index int) []string {
	ret := make([]string, 0)
	// up to but excluding index
	ret = append(ret, s[:index]...)

	// return starting from and including index+1
	return append(ret, s[index+1:]...)
}

// RemoveFromArrayInt safely removes an element
func RemoveFromArrayInt(s []int, index int) []int {
	ret := make([]int, 0)
	// up to but excluding index
	ret = append(ret, s[:index]...)

	// return starting from and including index+1
	return append(ret, s[index+1:]...)
}

// Read a file and return []string
func ReadFile(fileName string) []string {
	dataFile, err := os.Open(fileName)
	CheckErr(err)

	defer dataFile.Close()

	// create scanner to read file
	scanner := bufio.NewScanner(dataFile)

	// store the lines as string
	var lines []string

	// for each line, read and add it to string array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Read a file and return []int
func ReadFileInt(fileName string) []int {
	dataFile, err := os.Open(fileName)
	CheckErr(err)

	defer dataFile.Close()

	// create scanner to read file
	scanner := bufio.NewScanner(dataFile)

	// store the lines as string
	var lines []int

	// for each line, read and add it to int array
	for scanner.Scan() {
		convLine, err := strconv.Atoi(scanner.Text())
		CheckErr(err)
		lines = append(lines, convLine)
	}

	return lines
}
