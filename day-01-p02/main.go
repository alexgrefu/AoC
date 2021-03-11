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

	fmt.Printf("result: %d", findSumOfThree(input, 2020))
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

func findSumOfThree(input []int, sum int) int {
	sort.Ints(input)
	for a := 0; a < len(input)-2; a++ {
		i, n := a+1, len(input)-1
		for {
			total := input[a] + input[i] + input[n]
			if total == sum {
				return input[a] * input[i] * input[n]
			}
			if total > sum {
				if n == 0 { break }
				n--
			}
			if total < sum {
				if i == len(input)-1 { break }
				i++
			}
		}
	}
	panic("corrupted input file")
}