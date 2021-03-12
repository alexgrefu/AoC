package main

import (
	"bufio"
	"log"
	"os"
)

func parseAnswers(line string) map[rune]bool {
	group := map[rune]bool {}
	for _, c := range line {
		group[c] = true
	}
	return group
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allGroups := make([]map[rune]bool, 0)
	var line string

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() == "" {
			allGroups = append(allGroups, parseAnswers(line))
			line = ""
		} else {
			line += s.Text()
		}
	}

	if line != "" {
		//we have another group
		allGroups = append(allGroups, parseAnswers(line))
	}

	sum := 0
	for _, g := range allGroups {
		sum += len(g)
	}

	println(sum)
}
