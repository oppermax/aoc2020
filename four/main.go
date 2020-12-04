package main

import (
	"bufio"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var passport struct {
	byr int // Birth Year
	iyr int // Issue Year
	eyr int // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid int // Passport ID
	cid int // Country ID
}

func createRecords(input []string) []string{
	var out []string
	var t string
	for _, line := range input {
		if line != ""{
			t = t + line
		} else if line == ""{
			out = append(out, t)
			t = ""
		} else {
			log.Panicf("What is this: %v", line)
		}

	}
	out = append(out, t) // make sure the last is also added
	return out
}

func checkFields(passport string) bool {
	reByr := regexp.MustCompile(`(byr)`)
	reIyr := regexp.MustCompile(`(iyr)`)
	reEyr := regexp.MustCompile(`(eyr)`)
	reHgt := regexp.MustCompile(`(hgt)`)
	reHcl := regexp.MustCompile(`(hcl)`)
	reEcl := regexp.MustCompile(`(ecl)`)
	rePid := regexp.MustCompile(`(pid)`)
	// reCid := regexp.MustCompile(`(cid)`) CID is optional
	res := []*regexp.Regexp{reByr, reIyr, reEyr, reHgt, reHcl, reEcl, rePid}
	for _, re := range res {
		if re.Find([]byte(passport)) == nil {
			log.Info(passport)
			return false
		}

	}
	return true
}

func validateRecords(passports []string) int {
	var validRecords int
 	for _, passport := range passports {
			if checkFields(passport) {
				validRecords++
			}

		}
	return validRecords
	}

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var layers []string
	for scanner.Scan() {
		num := scanner.Text()

		layers = append(layers, num)
	}
	return layers
}

func main(){
	inputFile, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	lines := readFile(inputFile)
	records := createRecords(lines)
	validR := validateRecords(records)
	log.Infof("%v valid passports found, Total: %v", validR, len(records))


}