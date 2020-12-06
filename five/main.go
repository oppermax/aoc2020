package main

import (
	"bufio"
	"os"
	log "github.com/sirupsen/logrus"
	"sort"
)

func returnHalf(s []int, part string) []int {
	var out []int
	switch part {
	case "upper":
		return append(out, s[len(s)/2:]...)
	case "lower":
		return append(out, s[:len(s)/2]...)
	default:
		return out
	}
}

func findMissingID(s []int) int {
	var a []int
	for i := 89; i <= 888; i++{
		a = append(a, i)
	}
	for i, id := range a {
		if id == s[i] {
			continue
		} else if id != s[i] {
			return id
		}
	}
	return 0
}


func findRow(bp string) int {
	passRows := bp[:7]
	var rows []int
	for i := 0; i <= 127 ; i++ {
		rows = append(rows, i)
	}
		for _, letter := range passRows {
			if string(letter) == "B" {
				rows = returnHalf(rows, "upper")
			} else if string(letter) == "F"{
				rows = returnHalf(rows, "lower")
			}
		}
	return rows[0]
}

func findSeat(bp string) int {
	passSeats := bp[7:]
	var seats []int
	for i := 0; i <= 7 ; i++ {
		seats = append(seats, i)
	}
	for _, letter := range passSeats {
		if string(letter) == "R" {
			seats = returnHalf(seats, "upper")
		} else if string(letter) == "L"{
			seats = returnHalf(seats, "lower")
		}
	}
	return seats[0]
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
	file, err := os.Open("input")
	if err != nil {
		log.WithError(err)
	}
	boardingPasses := readFile(file)
	var seatID []int
	for _, pass := range boardingPasses{
		row := findRow(pass)
		seat := findSeat(pass)
		seatID = append(seatID, row *8 + seat)

	}
	sort.Ints(seatID)
	log.Infof("The highest seatID is %v", seatID[len(seatID)-1])
	missingSeat := findMissingID(seatID)
	log.Info("Your seatID is %v", missingSeat)



}