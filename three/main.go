package main

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func findTree(layers []string) int {
	var treeCount int
	steps := 0

	sLayers := []string{}
	for i, layer := range layers {
		if i%2 == 0 {
			sLayers = append(sLayers, layer)
		}
	}
	layers = sLayers

	for i, layer := range layers {

		layer = strings.Repeat(layer, i+10)
		if string(layer[steps]) == "#" {
			treeCount++
			steps += 1 // three steps to the right
		} else if string(layer[steps]) == "." {
			steps += 1 // three steps to the right

		} else {
			log.Panicf("Found a")
		}

	}
	return treeCount
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

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	defer input.Close()
	layers := readFile(input)
	out := findTree(layers)
	log.Infof("Found %v trees", out)
	log.Infof("Trees for two: %v", 94*99*214*91*46)
}
