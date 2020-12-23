package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
	"sort"
)


func checkInvalidNum(i int, nums []int) (bool, int){
	r := nums[i-25:i]
	for x, n := range r {
		for y, nu := range r {
			if x != y {
				if n + nu == nums[i] {
					return true, nums[i]
				}
			}
		}
	}
	return false, nums[i]
}


func parseNums(nums []int) int {
	for i := range nums {
		if i > 25 { // skip first 25 numbers
			b, n := checkInvalidNum(i, nums)
			if !b {
				return n
			}
		}
	}
	return 0
}

func findContNums(nums []int, n int) []int {
	for i, num := range nums {
		t := 0 // temp to count the sum
		ts := []int{num} // temp slice to add numbers
		t += num // add first contestant to temp
		for e, num2 := range nums {
			if e > i{ // start adding numbers after the first number
				if t < n{ // if t is smaller to n
					t += num2 // add the next number in the list
					ts = append(ts, num2)
				} else {
					break
				}
			}
		}
		if t == n {
			log.Infof("The Encription Weakness is %v",findEncWeakness(ts))
			return ts
		}
	}
	return []int{}
}

func findEncWeakness(in []int) int{
	sort.Ints(in)
	return in[0] + in[len(in)-1]
}

func main(){
	nums := utils.GetLinesInt("009/input")
	n := parseNums(nums)
	log.Infof("the first invalid number is %v", n)
	findContNums(nums, n)
}