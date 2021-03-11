package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)



func main() {
	fmt.Println("valid passwords: ", countValidPasswords("input.txt"))
}

func countValidPasswords(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := bufio.NewScanner(file)
	var valid int
	for s.Scan() {
		var min, max int
		var letter rune
		var passwd string

		if _, err := fmt.Sscanf(s.Text(), "%d-%d %c: %s", &min, &max, &letter, &passwd); err != nil {
			log.Fatal(err)
		}
		r := record{min, max, letter, make(map[rune]int)}
		for _, l := range passwd {
			r.Password[l]++
		}

		if r.isValid() { valid++ }
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return valid
}


