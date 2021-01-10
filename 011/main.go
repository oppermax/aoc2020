package main

import (
	"fmt"
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
)

type seat struct {
	x        int
	y        int
	occupied bool
	isFloor  bool
}

func realInt(in int) int {
	if in >= 0 {
		return in
	}
	return in * -1
}

func occupiedSeats(a []seat) int {
	out := 0
	for _, s := range a {
		if s.occupied {
			out++
		}
	}
	return out
}

func makeSeatMap(in []string) []seat {
	var out []seat
	totalSeats := 0
	for y, str := range in {
		for x, char := range str {
			t := seat{
				x: x,
				y: y,
			}
			switch string(char) {
			case ".":
				t.isFloor = true
				t.occupied = false
			case "L":
				totalSeats++
				t.isFloor = false
				t.occupied = false
			case "#":
				totalSeats++
				t.isFloor = false
				t.occupied = true
			}
			out = append(out, t)
		}
	}
	log.Warnf("There are %v seats available %v", totalSeats, len(out))
	return out
}

func checkRuleOne(one seat, two seat) (bool, string) {
	if one.x == two.x && one.y == two.y {
		return false, "sameSeat"
	}
	if two.occupied  || one.occupied{
		return false, "twoOccupied"
	}
	if one.isFloor || two.isFloor {
		return false, "isFloor"
	}
	if one.x == two.x && one.y == two.y-1 || one.x == two.x && one.y == two.y+1 || one.x == two.x-1 && one.y == two.y || one.x == two.x+1 && one.y == two.y{
		log.Infof("%v is next to %v", two, one)
		return true, "nextToOne"
	}
	return false, "anyOther"
}

func checkRuleTwo(one seat, two seat) (bool, string){
	if one.x == two.x && one.y == two.y {
		return false, "sameSeat"
	}
	if !two.occupied || !one.occupied{
		return false, "twoOccupied"
	}
	if one.isFloor || two.isFloor {
		return false, "isFloor"
	}
	if !one.occupied {
		return true, "nextToOne"
	}
	if one.x == two.x && one.y == two.y-1 || one.x == two.x && one.y == two.y+1 || one.x == two.x-1 && one.y == two.y || one.x == two.x+1 && one.y == two.y{
		//log.Infof("%v is next to %v", two, one)
		return true, "nextToOne"
	}
	return false, "anyOther"
}

func changeState(in []seat) []seat {
	drawMap(in)
	var out []seat
	oneCount := 0
	twoCount := 0
	for _, one := range in {

		if !one.isFloor {
			ruleOne := 0 // if one next to it is occupied
			ruleTwo := 0
			if !one.occupied {
				for _, two := range in {
					b, reason := checkRuleOne(one, two)
					if b{
						ruleOne++
					}
					log.Warn(reason)
				}

			}
			if one.occupied {
				for _, two := range in {
					b, reason := checkRuleTwo(one, two)
					if  b{
						ruleTwo++
					}
					log.Warn(reason)

				}
			}
			if ruleOne <= 1 && !one.occupied {
				t := seat{
					x:        one.x,
					y:        one.y,
					occupied: true,
					isFloor:  false,
				}
				oneCount++
				out = append(out, t)
				ruleOne = 0

			} else if ruleTwo >= 4 && one.occupied {
				if one.x == 0 && one.y == 0 {
					log.Info(ruleTwo)
				}
				t := seat{
					x:        one.x,
					y:        one.y,
					occupied: false,
					isFloor:  false,
				}
				twoCount++
				out = append(out, t)
				ruleTwo = 0
			} else {
				out = append(out, one)
				if one.isFloor {
					log.Panic()
				}
			}
		} else {
			out = append(out, one)
		}


	}
	drawMap(out)
	return out
}

func changeUntilConst(seatMap []seat) int {
	newMap := changeState(seatMap)

	if occupiedSeats(newMap) == occupiedSeats(seatMap){
		return occupiedSeats(newMap)
	}
	o := 0
	for i := range seatMap {
		if seatMap[i] == newMap[i] {
			o++
		}
	}
	if o >= len(newMap) {
		//log.Info(o, len(newMap))
		return occupiedSeats(newMap)
	}
	log.Infof("Occupied seats: %v",occupiedSeats(newMap))
	return changeUntilConst(newMap)
}

func drawMap(seatMap []seat) {
	fmt.Println("+++++++++++++ NEW MAP +++++++++++++")
	var t []string
	for i, seat := range seatMap {
		if i % 10 == 0 && i != 0{
			fmt.Println(t)
			t = []string{}
		}
		if seat.isFloor {
			t = append(t, ".")
		} else if seat.occupied {
			t = append(t, "#")
		} else if !seat.occupied {
			t = append(t, "L")
		}
	}
}

func main() {
	raw := utils.GetLines("011/input")
	seatMap := makeSeatMap(raw)
	//log.Info(occupiedSeats(seatMap))
	//newMap := changeState(seatMap)
	//log.Info(occupiedSeats(newMap))
	log.Info(changeUntilConst(seatMap))
}
