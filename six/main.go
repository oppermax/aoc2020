package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func countAnswers(group string) int{
	c := ""
	for _, answer := range group {
		if strings.Contains(c, string(answer)){
			continue
		} else {
			c = c + string(answer)
		}

	}
return len(c)
}

func countAnswersEveryone(group []string) int{
	log.Info(len(group))
	var c string
	var out int
	for _, person := range group {
		log.Info(person)
		for _, answer := range person {
			c = c + string(answer)

		}
	}
	var checked string
	for _, a := range c{
		if strings.Contains(checked, string(a)) != true {
			if strings.Count(c, string(a)) == len(group) {
				log.Infof("%v is %v times in a group of %v people", string(a), strings.Count(c, string(a)), len(group))

				out++
				checked = checked + string(a)
			}
		}
	}
	return out
}

func splitByIndividuals(input []string) [][]string {
	var out [][]string
	var t []string
	for _, i := range input {
		if i != "" {
			t = append(t, i)
		} else {
			out = append(out, t)
			t = []string{}
		}
	}
	return out
}

func splitByGroups(input []string) []string{
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
func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var layers []string
	for scanner.Scan() {
		layers = append(layers, scanner.Text())
	}
	return layers
}

func main(){
	inputFile, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	file := readFile(inputFile)

	groups := splitByGroups(file)

	var yesAnswers int
	for _, group := range groups {
		yesAnswers += countAnswers(group)
	}
	log.Infof("All questions answered with yes summed up: %v", yesAnswers)

	yesAnswers = 0
	ppl := splitByIndividuals(file)
	for _, group := range ppl {
		yesAnswers += countAnswersEveryone(group)
	}

	log.Infof("All questions everyone answered with yes summed up: %v", yesAnswers)



}