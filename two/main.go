package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func checkPw1(line string) bool {
	nums := findNum(line)
	char := findChar(line)
	pw := findPw(line)
	lo := nums[0]
	hi := nums[1]
	var count int

	log.Infof("This should contain between %v and %v instances of %v in %v", lo, hi, char, pw)

	for _, c := range pw {
		if string(c) == char {
			count++
		}
	}
	if lo > count || count > hi {
		return true
	} else {
		return false
	}
}

func checkPw2(line string) bool {
	nums := findNum(line)
	char := findChar(line)
	pw := findPw(line)
	lo := nums[0] -1
	hi := nums[1] -1

	log.Infof("\nThis should have a %v at %v or %v in %v len: %v", char, lo, hi, pw, len(pw))

	if char == string(pw[lo]) && char == string(pw[hi]) {
		log.Infof("This is false. %v is at %v and %v is at %v", string(pw[lo]), lo, string(pw[hi]), hi)
		return false
	} else if char == string(pw[lo]) || char == string(pw[hi]) {
		log.Infof("This is true. %v is at %v and %v is at %v", string(pw[lo]), lo, string(pw[hi]), hi)
		return true
	} else if char != string(pw[lo]) && char != string(pw[hi]){
		log.Infof("Something is wrong: %v is at %v and %v is at %v", string(pw[lo]), lo, string(pw[hi]), hi)
		return false
	} else {
		log.Panic()
		return false
	}
}

func findNum(input string) []int {
	re := regexp.MustCompile(`(\w+-\w+)`) // find number range
	re2 := regexp.MustCompile(`\d+`)      // find individual numbers
	t := string(re.Find([]byte(input)))
	numByte := re2.FindAll([]byte(t), -1)
	var nums []int
	for _, num := range numByte {
		num, err := strconv.Atoi(string(num))
		if err != nil {
			log.Panicf("Could not create int from line string: ", err)
		}
		nums = append(nums, num)

	}
	return nums // return a slice of two numbers e.g. [2,14]

}

func findChar(input string) string {
	re := regexp.MustCompile(`(\w+:)`)
	re2 := regexp.MustCompile(`\w+`)
	t := string(re.Find([]byte(input)))
	return string(re2.Find([]byte(t)))
}

func findPw(input string) string {
	re := regexp.MustCompile(`(\s\w+\w)`)
	re2 := regexp.MustCompile(`(\w+\w)`)
	t := string(re.Find([]byte(input)))
	return string(re2.Find([]byte(t)))

}

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var passwords []string
	for scanner.Scan() {
		num := scanner.Text()

		passwords = append(passwords, num)
	}
	return passwords
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	defer input.Close()
	passwords := readFile(input)

	var wPw1, wPw2 int

	for _, line := range passwords {
		// log.Infof("%v: %v", i, line)
		t := checkPw1(line)
		if t == false {
			wPw1++
		}

	}
	for _, line := range passwords {
		t := checkPw2(line)
		if t == true {
			wPw2++
		}
	}
	log.Infof("%v incorrect passwords for part one and %v for part two", wPw1, wPw2)

}
