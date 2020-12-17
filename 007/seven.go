package main

import (
	"regexp"
	"strconv"
	log "github.com/sirupsen/logrus"
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
