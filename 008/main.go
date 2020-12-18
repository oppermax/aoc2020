package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

type accumulator struct {
	setting int
}

type rule struct {
	kind string
	value int
	checked bool
}

func (a *accumulator) changeAcc(v int) {
	a.setting += v
}


func makeRuleMap(rules []string) map[int]rule {
	out := make(map[int]rule)
	for i, r := range rules {
		kind, value := getRuleDetails(r)
		out[i] = rule{
			kind:    kind,
			value:   value,
			checked: false,
		}
	}
	return out
}

func getRuleDetails(rule string) (string, int) {
	reKind := regexp.MustCompile(`(\w{3})`)
	reValue := regexp.MustCompile(`([^\w\s]\d+)`)
	kind := string(reKind.Find([]byte(rule)))
	valueStr := string(reValue.Find([]byte(rule)))
	value, err := strconv.Atoi(valueStr)
	if err != nil{
		log.WithError(err)
	}

	return kind, value
}



func readRules(ruleMap map[int]rule, acc accumulator, iterator int) int {
	if ruleMap[iterator].checked {
		return acc.setting
	} else if !ruleMap[iterator].checked {
		switch ruleMap[iterator].kind {
		case "acc":
			acc.setting += ruleMap[iterator].value // plus + plus = plus ; plus + minus = minus
			iterator  += 1
		case "jmp":
			iterator += ruleMap[iterator].value
		case "nop":
			iterator  += 1
		}
		t := iterator - 1 // as this happens after the iterator is advanced, we need t to check of the current rule and not the next
		ruleMap[t] = rule{
			checked: true,
		} // mark as checked

		acc.setting = readRules(ruleMap, acc, iterator)
	}
	return acc.setting
}

func main(){
	acc := accumulator{setting: 0}
	iterator := 0
	rules := utils.GetLines("008/input")
	ruleMap := makeRuleMap(rules)
	accSetting := readRules(ruleMap, acc, iterator)
	log.Infof("At the time of system repetition, the accumulator is set to %v", accSetting)
}