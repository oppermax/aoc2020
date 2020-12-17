package main

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func twoMake2020(input []int) int {
	for _, numberOne := range input {
		for _, numberTwo := range input {
			if numberOne+numberTwo == 2020 {
				return numberOne * numberTwo
			}
		}

	}
	return 0
}

func threeMake2020(input []int) int {
	for _, numOne := range input {
		for _, numTwo := range input {
			for _, numThree := range input {
				// log.Info(numOne + numTwo + numThree)
				if numOne+numTwo+numThree == 2020 {
					return numOne * numTwo * numThree
				}
			}
		}
	}
	return 0
}

func readFile(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("Could not create int from line string: ", err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {

	input, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	defer input.Close()
	numbers := readFile(input)

	a := twoMake2020(numbers)
	if a != 0 {
		log.Infof("Solution of two that make 2020: %v ", a)
	}
	s := threeMake2020(numbers)
	if s != 0 {
		log.Infof("Solution of three that make 2020: %v", s)
	}

}
