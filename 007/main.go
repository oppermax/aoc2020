package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

type rule struct {
	origin string
	targets map[string]int
	total int
}




func fillTargetsMap(targetRaw []byte) (string, int){
	targetRe := regexp.MustCompile(`([^\d\s]\w+ \w+)`)
	numRe := regexp.MustCompile(`(\d+)`)
	num, err := strconv.Atoi(string(numRe.Find(targetRaw)))
	if err != nil {
		log.WithError(err)
	}
	target := string(targetRe.Find(targetRaw))

	return target, num

}

func makeRuleStruct(rawRule string) rule{
	byteRule := []byte(rawRule)
	targets := make(map[string]int)
	total := 0
	orgiginRe := regexp.MustCompile(`(\w+ \w+)`) // shiny
	targetsRe := regexp.MustCompile(`(\d+ \w+ \w+)`) // 5 shiny golden
	orgigin := string(orgiginRe.Find(byteRule))
	targetsRaw := targetsRe.FindAll(byteRule, -1)
	for _, targetRaw := range targetsRaw {
		target, num := fillTargetsMap(targetRaw)
		targets[target] = num
		total += num
	}
	 out := rule{
		origin:  orgigin,
		targets: targets,
		total:   total,
	}
	return out

}

func makeRuleMap(rules []string) map[string]rule{
	out := make(map[string]rule)

	for _, rule := range rules {
		ruleStruct := makeRuleStruct(rule)
		out[ruleStruct.origin] = ruleStruct
	}
	return out
}

func countSubBags(ruleMap map[string]rule, startRule string) int{
	out := 0
	if ruleMap[startRule].targets == nil { // base case
		return 0
	}
	for nextRule, t := range ruleMap[startRule].targets {
	 	out += t * countSubBags(ruleMap, nextRule)
	}

	return out + ruleMap[startRule].total
}

func countGoldenBags(ruleMap map[string]rule, lookingFor string, checked map[string]bool) (int, map[string]bool) {
	out := 0
	if ruleMap[lookingFor].targets == nil {
		return 0, checked
	}

	for r := range ruleMap {
		if _, ok := ruleMap[r].targets[lookingFor]; ok {
			add := 0
			if !checked[r] { // need this to avoid duplication
				checked[r] = true
				add, checked = countGoldenBags(ruleMap, r, checked)
				out += 1 + add
			} else {
			}

			}
	}

	return out, checked
}



func main(){
	rules := utils.GetLines("007/input")
	ruleMap := makeRuleMap(rules)
	const startRule = "shiny gold"
	checked := make(map[string]bool)
	totalOne, _ := countGoldenBags(ruleMap, startRule, checked)
	totalTwo := countSubBags(ruleMap, startRule)
	log.Infof("The results:\n%v bags can eventually hold at least one shiny golden bag\n%v bags are needed to carry one golden bag", totalOne, totalTwo)
}