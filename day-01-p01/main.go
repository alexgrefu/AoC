package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	input := readInput("input.txt")

	fmt.Printf("result: %d", findSumOfTwo(input, 2020))
}

func readInput(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var n int
	var numbers = make([]int, 0)
	for s.Scan() {
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}

	return numbers
}

func findSumOfTwo(input []int, sum int) int {
	sort.Ints(input)
	i, n := 0, len(input) - 1
	for {
		total := input[i] + input[n]
		if total == sum {
			return input[i] * input[n]
		}
		if total > sum { n-- }
		if total < sum { i++ }
	}
}