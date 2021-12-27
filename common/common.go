package common

import (
	log "github.com/sirupsen/logrus"
	"os"
	"bufio"
	"strconv"
)

// checks an error 
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFile(fileName string) ([]string) {
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

func ReadFileInt(fileName string) ([]int) {
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