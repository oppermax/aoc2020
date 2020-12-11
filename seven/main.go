package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strconv"
	"time"
)


func findRelevantRules(allRules []string, relRules []string) []string {
	var t []string
	reNewRule := regexp.MustCompile(`\w+ \w+`)
	if len(relRules) == 0 {
		relRules = append(relRules, "shiny gold")
	}
	for _, rule := range allRules{
		for _, relRule := range relRules {

			reRaw := `(\d+ ` + relRule + `)`
			re := regexp.MustCompile(reRaw)
			if re.Find([]byte(rule)) != nil {
				nextRule := reNewRule.Find([]byte(rule))
				t = append(t, string(nextRule))
			}
		}
	}
	return t

}

func findNextLevelRules(allRules []string, relRules map[string]int) map[string]int {
	if len(relRules) == 0 {
		relRules["shiny gold"] = 0
	}
	for _, rule := range allRules{
		for relRule, _ := range relRules {
			// log.Infof("New rule: %v", relRule)
			re2 := regexp.MustCompile(`( ` + relRule + `)`) // eg `(shiny gold)`
			re3 := regexp.MustCompile(`(\d+ \w+ \w+)`) // eg 3 wavy posh
			re3sub1 := regexp.MustCompile(`\d+`) // eg 3
			re3sub2 := regexp.MustCompile(`\w+ \w+`) // eg wavy posh
			if re2.Find([]byte(rule)) != nil { // finds rule for shiny gold
				nextRulez := re3.FindAll([]byte(rule), -1) // finds []of rules under shiny gold
				totalSubBags := 0
				for _, r := range nextRulez {
					num, err := strconv.Atoi(string(re3sub1.Find(r))) // finds 3
					if err != nil {
						log.WithError(err).Panic("Could not convert to int")
					}
					nextRule := re3sub2.Find([]byte(rule)) // finds wavy posh
					relRules[string(nextRule)] = num
					totalSubBags += num
				}
				relRules[relRule] = totalSubBags

			}
		}
	}
	return relRules

}

//func numGoldenBags(rule string) int {
//	out := 0
//	re1 := regexp.MustCompile(`(\d+ shiny gold)`)
//	re2 := regexp.MustCompile(`\d+`)
//	r := re1.Find([]byte(rule))
//	if r != nil {
//		num, err := strconv.Atoi(string(re2.Find(r)))
//		if err != nil {
//			log.WithError(err).Panic("Could not convert to int. Is regex incorrect?")
//		}
//		return num
//	}
//	return out
//}



func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var layers []string
	for scanner.Scan() {
		num := scanner.Text()

		layers = append(layers, num)
	}
	return layers
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func main(){
	start := time.Now()
	inputFile, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	rules := readFile(inputFile)
	var total []string
	relRules := []string{"shiny gold"}
	for len(relRules) > 0 {
		relRules = findRelevantRules(rules, relRules)
		for _, rule := range relRules {
			total = append(total, rule)
		}
	}
	total = removeDuplicates(total)
	end := time.Since(start)
	log.Info(len(total), end)

	topRules := make(map[string]int)
	for len(topRules) < 91 {
		log.Info(len(topRules))
		topRules = findNextLevelRules(rules, topRules)
	}
	totalGold := 0
	for key, value := range topRules{
		log.Info(key, value)
		totalGold += value
	}

	log.Info(totalGold)



}