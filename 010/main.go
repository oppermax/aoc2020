package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
	"sort"
)

func assembleAdapters(adapters []int) (int,int,int){
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3) // add the joltage of the last device
	var oneJolt int
	var twoJolt int
	var threeJolt int
	b := 0
	for _, adapter := range adapters {
		d := adapter - b
		switch d {
		case 1:
			oneJolt += 1
		case 2:
			twoJolt += 1
		case 3:
			threeJolt += 1
		default:
			log.Panicf("No compatible adapter for difference: %v", d)
		}
		b = adapter
	}
	return oneJolt, twoJolt, threeJolt
}

func findSub(adapters []int, adapterMap map[int]int, index int) int{
	out := 0
	if adapterMap[adapters[index]] != -1 {
		if adapterMap[adapters[index]] != 0 {
			return adapterMap[adapters[index]]
		}
	}
	if index >= len(adapters) - 1{
		return 1
	}
	if len(adapters) > index + 1 && adapters[index + 1 ] - adapters[index] <= 3{
		out += findSub(adapters, adapterMap, index + 1)
	}
	if len(adapters) > index + 2 && adapters[index + 2 ] - adapters[index] <= 3{
		out += findSub(adapters, adapterMap, index + 2)
	}
	if len(adapters) > index + 3 && adapters[index + 3 ] - adapters[index] <= 3{
		out += findSub(adapters, adapterMap, index + 3)
	}
	adapterMap[adapters[index]] = out
	return out
}

func makeMap(adapters []int) (map[int]int, []int) {
	outMap := make(map[int]int)
	outSlice := []int{0} // first device
	outMap[0] = -1 // first device
	for _, a := range adapters {
		outMap[a] = -1
		outSlice = append(outSlice, a)
	}
	outMap[adapters[len(adapters)-1]+3] = -1 // last device
	outSlice = append(outSlice, adapters[len(adapters)-1]+3) // last device
	log.Warn(outSlice)
	log.Info(outMap)
	return outMap, outSlice
}

func main() {
	nums := utils.GetLinesInt("010/input")
	x, _, z := assembleAdapters(nums)
	adapterMap, adapters := makeMap(nums)
	log.Infof("Result for part one: %v",x*z)
	o := findSub(adapters, adapterMap, 0)
	log.Infof("Result for part two: %v",o)
}