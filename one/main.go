package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func twoMake2020(c int, p int, input []int) (int, int) {
	for i, number := range input {
		p := c + number
		if i == p {
			continue
		}
		if p == 2020 {
			return c, number
		}
		//log.Infof("Possibly %v", p)
		fmt.Println(p)
	}
	return 0, 0
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
	// for i, num := range numbers {
	// 	a, b := twoMake2020(num, i, numbers)
	// 	if a != 0 {
	// 		log.Infof("Solution: %v ", a*b)
	// 		return
	// }
	// }
	s := threeMake2020(numbers)
	if s != 0 {
		log.Infof("Solution: %v", s)
	} else {
		log.Info("No solution found")
	}

}
