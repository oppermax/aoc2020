package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type passport struct {
	byr int // Birth Year
	iyr int // Issue Year
	eyr int // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid int // Passport ID
	cid int // Country ID
}

func byteToInt(in []byte) int {
	out, err := strconv.Atoi(string(in))
	if err != nil{
		log.WithError(err).Panicf("Could not convert byte to int")
	}
	return out
}

func strToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil{
		log.WithError(err).Panicf("Could not convert string to int")
	}
	return out
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
func checkEcl(field []byte) bool{
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	docEcl := string(field)[len(string(field))-3:]
	for _, color := range colors {
		if docEcl == color{
			return true
		}
	}
	log.Infof("Invalid. Reason: %v: %v", string(field), docEcl)
	return false
}

func checkFields(input string) bool {
	reByr := regexp.MustCompile(`(byr:\d{4})`) // birth year
	reByrSub := regexp.MustCompile(`(\d{4})`)
	reIyr := regexp.MustCompile(`(iyr:\d{4})`) // issue year
	reIyrSub := regexp.MustCompile(`(\d{4})`)
	reEyr := regexp.MustCompile(`(eyr:\d{4})`) // expiration year
	reEyrSub := regexp.MustCompile(`(\d{4})`)
	reHgt := regexp.MustCompile(`(hgt:\d+\w{2})`) // height
	reHgtSub := regexp.MustCompile(`(\d+\w{2})`)
	reHcl := regexp.MustCompile(`(hcl:#\w{6})`) // hair color
	reHclSub := regexp.MustCompile(`(#\w{6})`) // hair color
	reEcl := regexp.MustCompile(`(ecl:\w{3})`) // eye color
	reEclSub := regexp.MustCompile(`(\w{3})`) // eye color
	rePid := regexp.MustCompile(`(pid:\d{9})`) // passport ID
	rePidSub := regexp.MustCompile(`(\d{9})`) // passport ID
	// reCid := regexp.MustCompile(`(cid)`) CID is optional
	res := []*regexp.Regexp{reByr, reIyr, reEyr, reHgt, reHcl, reEcl, rePid}
	subRes := []*regexp.Regexp{reByrSub, reIyrSub, reEyrSub, reHgtSub, reHclSub, reEclSub, rePidSub}


	for i, re := range res {
		field := re.Find([]byte(input))
		if field == nil {
			log.Info("Invalid. Reason: A field missing")
			return false
		}
		val := subRes[i].Find(field)
		switch i{
		case 0:
			intVal := byteToInt(val)
			if  intVal >= 1920 && intVal <= 2002 {
			} else {
				log.Infof("Invalid. Reason: %v: %v incorrect", string(field), intVal)
				return false
			}
		case 1:
			intVal := byteToInt(val)
			if  intVal >= 2010 && intVal <= 2020 {
			} else {
				log.Infof("Invalid. Reason: %v: %v incorrect", string(field), intVal)
				return false
			}
		case 2:
			intVal := byteToInt(val)
			if  intVal >= 2020 && intVal <= 2030 {
			} else {
				log.Infof("Invalid. Reason: %v: %v incorrect", string(field), intVal)
				return false
			}
		case 3:
			t := string(val)
			v := strToInt(t[:len(t)-2])
			unit := t[len(t)-2:]
			switch unit{
			case "cm":
				if  v >= 150 && v <= 193 {
				} else {
					log.Infof("Invalid. Reason: %v: %v cm is incorrect", string(field), v)
					return false
				}
			case "in":
				if  v >= 59 && v <= 76 {
				} else {
					log.Infof("Invalid. Reason: %v: %v in is incorrect", string(field), v)
					return false
				}
			}
		case 4:
			// should be covered by regex
			log.Info(string(val))
			if val == nil{

				return false
			}
		case 5:
			if checkEcl(field) != true {
				return false
			}
		case 6:
			// should be covered by regex

		}

	}
	log.Infof("Valid.")
	return true
}

func validateRecords(passports []string) int {
	var validRecords int
 	for i, passport := range passports {
 		log.Infof("Passport %v", i)
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

