package aoc2020

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		num := scanner.Text()

		lines = append(lines, num)
	}
	return lines
}

func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	out := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			out = append(out, elements[v])
		}
	}
	return out
}

func GetLines(filepath string) []string {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	return readFile(inputFile)
}
