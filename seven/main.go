package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
)






func main(){
	rules := utils.GetLines("seven/input")
	var total []string
	relRules := []string{"shiny gold"}
	for len(relRules) > 0 {
		relRules = FindRelevantRules(rules, relRules)
		for _, rule := range relRules {
			total = append(total, rule)
		}
	}
	total = utils.RemoveDuplicates(total)

	topRules := make(map[string]int)
	for len(topRules) < 91 {
		log.Info(len(topRules))
		topRules = FindNextLevelRules(rules, topRules)
	}
	totalGold := 0
	for key, value := range topRules{
		log.Info(key, value)
		totalGold += value
	}

	log.Info(totalGold)

}