package main

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func findTree(layers []string) int{
	var treeCount int
	steps := 0
	for i, layer := range layers{

		log.Infof("%v: %v", i, layer)
		layer = strings.Repeat(layer, i+10)
			if string(layer[steps]) == "#" {
				treeCount++
				log.Infof("Found a tree %v", string(layer[:steps+1]))
				steps = i * 3 // three steps to the right
			} else if string(layer[steps]) == "." {
				steps = i * 3 // three steps to the right

			} else {
				log.Panicf("Found a %s", layer[steps])
			}
		log.Infof("Steps: %v", steps)


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

func main(){
	input, err := os.Open("input")
	if err != nil {
		log.Infof("Could not read file: %v", err)
	}
	defer input.Close()
	layers := readFile(input)
	out := findTree(layers)
	log.Infof("Found %v trees", out)
}
