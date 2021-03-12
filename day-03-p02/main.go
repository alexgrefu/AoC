package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type slope struct {
	right int
	down int
}

func main() {
	slopes := [] slope {
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	result := 1
	treeMap := generateTreemap(slopes, "input.txt")
	for _, s := range slopes {
		n := countTrees(treeMap, s)
		result *= n
		fmt.Printf("slope: %#v, trees encountered: %d\n", s, n)
	}

	fmt.Printf("result: %d", result)
}

func generateTreemap(slopes []slope, fileName string) [][]bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var lines = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// find the max right from the given slopes in order to generate a map
	// large enough to fit all slopes
	maxRight := 0
	for _, s := range slopes {
		if s.right > maxRight { maxRight = s.right }
	}

	// the vertical slots are the number of lines from the input file
	verticalSlots := len(lines)
	// the initial non-repeated horizontal slots is the length of any line from the input file
	// we take the first one
	horizontalSlots := len(lines[0])

	// the multiplication factor: in order to get a solution lines have to be
	// repeated horizontally until we have a window large enough to calculate a solution
	// because the slope is: down x, right y, the horizontal slots count has to be y times larger
	// then the vertical count
	multiplicationFactor := ((verticalSlots * maxRight) / horizontalSlots) + 1 // we add one to account for division rounding

	treeMap := make ([][]bool, verticalSlots)
	horizontalCap := len(strings.Repeat(lines[0], multiplicationFactor))
	for i := range treeMap {
		treeMap[i] = make([]bool, horizontalCap)
	}

	// populate the 2D map with true value for all coords that have a tree and false if they don't have a tree
	for i, str := range lines {
		lines[i] = strings.Repeat(str, multiplicationFactor)
		signs := []rune(lines[i])
		for k, val := range signs {
			if val == '.' {
				treeMap[i][k] = false
			} else if val == '#' {
				treeMap[i][k] = true
			} else {
				panic("unreachable")
			}
		}
	}

	return treeMap
}

func countTrees(treeMap [][]bool, s slope) int {

	hIndex := 0
	trees := 0
	for vIndex := s.down; vIndex < len(treeMap); vIndex += s.down {
		hIndex += s.right
		if treeMap[vIndex][hIndex] { trees++ }
	}

	return trees
}
