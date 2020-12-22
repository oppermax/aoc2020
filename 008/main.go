package main

import (
	"errors"
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



func forceProgramFinish(rules []string) {
	acc := accumulator{setting: 0}
	iterator := 0

	for i := 0; i<= len(rules); i++ {
		ruleMap := makeRuleMap(rules)
		if ruleMap[i].kind == "jmp" {
			ruleMap[i] = rule{
				kind:    "nop",
				value:   ruleMap[i].value,
				checked: false,
			}
			log.Infof("%v changed from jmp to %v", i, ruleMap[i].kind)
			log.Info("Calling readRules() with custom rulemap")
			out, err := readRules(ruleMap, acc, iterator)
			if out == 920 {
				log.Panic()
			}
			if err == nil{
				log.Infof("success %v", out)
			} else {
				log.Infof("Accumulator still at %v: %v", out, err)
			}
		}
	}
}


func readRules(ruleMap map[int]rule, acc accumulator, iterator int) (int, error) {
	if ruleMap[iterator].checked {
		err := errors.New("condition already checked")
		return acc.setting, err
	} else if !ruleMap[iterator].checked {
		switch ruleMap[iterator].kind {
		case "acc":
			acc.setting += ruleMap[iterator].value // plus + plus = plus ; plus + minus = minus
			iterator  += 1
		case "jmp":
			iterator += ruleMap[iterator].value
			if ruleMap[iterator+1].kind != "jmp" {
			}
		case "nop":
			iterator  += 1
		}
		t := iterator - 1 // as this happens after the iterator is advanced, we need t to check of the current rule and not the next
		ruleMap[t] = rule{
			kind: ruleMap[t].kind,
			value: ruleMap[t].value,
			checked: true,
		} // mark as checked

		acc.setting, _ = readRules(ruleMap, acc, iterator)
	}
	return acc.setting, nil
}

func main(){
	acc := accumulator{setting: 0}
	iterator := 0
	rules := utils.GetLines("008/input")
	ruleMap := makeRuleMap(rules)
	accSetting, _ := readRules(ruleMap, acc, iterator)
	log.Infof("At the time of system repetition, the accumulator is set to %v", accSetting)
	forceProgramFinish(rules)
}