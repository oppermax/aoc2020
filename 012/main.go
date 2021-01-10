package main

import (
	utils "github.com/oppermax/aoc2020"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type position struct {
	direction int
	east      int
	north     int
}

type adjustment struct {
	kind  string // L, R, N, W, E, S, F
	value int
}

func (a *adjustment) performAdjustment(in position) position {
	log.Infof("We are at %v", in)
	log.Infof("Adjusting position by %v %v", a.kind, a.value)
	switch a.kind {
	case "L":
		return position{
			direction: setDirection(a.value*-1, in),
			east:      in.east,
			north:     in.north,
		}
	case "R":
		return position{
			direction: setDirection(a.value, in),
			east:      in.east,
			north:     in.north,
		}
	case "N":
		t := in.north + a.value
		return position{
			direction: in.direction,
			east:      in.east,
			north:     t,
		}
	case "S":
		t := in.north - a.value
		return position{
			direction: in.direction,
			east:      in.east,
			north:     t,
		}
	case "E":
		t := in.east + a.value
		return position{
			direction: in.direction,
			east:      t,
			north:     in.north,
		}
	case "W":
		t := in.east - a.value
		return position{
			direction: in.direction,
			east:      t,
			north:     in.north,
		}
	case "F":
		switch in.direction {
		case 0:
			t := in.north + a.value
			return position{
				direction: in.direction,
				east:      in.east,
				north:     t,
			}
		case 90:
			t := in.east + a.value
			return position{
				direction: in.direction,
				east:      t,
				north:     in.north,
			}
		case 180:
			t := in.north - a.value
			return position{
				direction: in.direction,
				east:      in.east,
				north:     t,
			}
		case 270:
			t := in.east - a.value
			return position{
				direction: in.direction,
				east:      t,
				north:     in.north,
			}
		}
		log.Warnf("No case found %v %v", a.kind, in.direction)
	}
	return in
}

func setDirection(degrees int, pos position) int {
	x := (pos.direction + degrees) / 90
	switch x {
	case -1:
		x = 3
	case -2:
		x = 2
	case -3:
		x = 1
	case -4:
		x = 0
	}
	for x >= 4 {
		x -= 4
	}
	log.Infof("Position %v", x*90)
	return x * 90
}

func translate(lines []string) []adjustment {
	var out []adjustment
	for _, l := range lines {
		s, err := strconv.Atoi(l[1:])
		if err != nil {
			log.WithError(err)
		}
		t := adjustment{
			kind:  string(l[0]),
			value: s,
		}
		out = append(out, t)
	}
	return out
}

func main() {
	lines := utils.GetLines("012")
	pos := position{ // starting position, in the middle facing east
		direction: 90,
		east:      0,
		north:     0,
	}
	adjustments := translate(lines)
	for _, a := range adjustments {
		pos = a.performAdjustment(pos)
	}
	log.Infof("New position %v", pos)
	var md int
	if pos.north < 0 {
		md += pos.north * -1
	} else {
		md += pos.north
	}
	if pos.east < 0 {
		md += pos.east * -1
	} else {
		md += pos.east
	}
	log.Infof("Manhattan distance %v", md)

}
