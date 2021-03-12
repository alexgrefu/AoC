package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxId, ids := getSeatIds(file)

	sort.Ints(ids)

	fmt.Printf("Part 1: Max seat ID %d\n", maxId)
	fmt.Printf("%v\n", ids)

	for i := 0; i < len(ids) - 2; i++ {
		if ids[i+1] - ids[i] > 1 {
			fmt.Printf("seat ID: %d", ids[i] + 1)
		}
	}
}

func getSeatIds(file *os.File) (int, []int) {
	maxId := 0
	ids := make([]int, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		input := strings.TrimSpace(s.Text())

		row := allocation(0, 127, input)
		column := allocation(0, 7, input[7:])
		seatId :=(row * 8) + column
		ids = append(ids, seatId)
		fmt.Printf("row %d, column: %d, seat ID: %d\n", row, column, seatId)

		if seatId > maxId {
			maxId = seatId
		}
	}
	
	return maxId, ids
}

func allocation(low, high int, letters string) int {
	r := rune(letters[0])
	//println("low:", low, "high:", high, "letters:", letters, "letter:", r)

	count := high - low
	diff := count / 2

	// invalid start number
	retVal := -1

	if count == 1 {
		if r == 'F' || r == 'L' {
			retVal = low
		} else if r == 'B' || r == 'R' {
			retVal = high
		}
	} else {
		if r == 'F' || r == 'L' {
			retVal = allocation(low, low+diff, letters[1:])
		} else if r == 'B' || r == 'R' {
			retVal = allocation(high-diff, high, letters[1:])
		}
	}
	return retVal
}
