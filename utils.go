package utils

import (
	"bufio"
	"os"
	log "github.com/sirupsen/logrus"
)

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		num := scanner.Text()

		lines = append(layers, num)
	}
	return lines
}

func getLines(filepath string) []string {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	return readFile(inputFile)
}
