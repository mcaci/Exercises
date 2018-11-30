package main

import (
	"bufio"
	"os"
)

// LoadFromFile returns a slice of lines of which the input file is composed
func LoadFromFile(fileName string) (lines []string, err error) {
	directory := "./"
	file, err := os.Open(directory + fileName)
	if err == nil {
		defer file.Close()
		lines, err = readLines(bufio.NewScanner(file))
	}
	return lines, err
}

func readLines(scanner *bufio.Scanner) ([]string, error) {
	var lines []string
	var err error
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, err
}
