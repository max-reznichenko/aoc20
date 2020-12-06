package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	seats := []int{}
	for _, pass := range formatInput(readInput()) {
		seats = append(seats, int(seatID(pass)))
	}
	sort.Ints(seats)

	for idx, seat := range seats {
		if idx != 0 && seat-seats[idx-1] > 1 {
			fmt.Println(seat - 1)
		}
	}
}

func seatID(pass string) int64 {
	row, _ := strconv.ParseInt(pass[0:7], 2, 64)
	column, _ := strconv.ParseInt(pass[7:], 2, 64)
	return row*8 + column
}

func formatInput(input []string) []string {
	passes := []string{}
	for _, pass := range input {
		pass = strings.ReplaceAll(pass, "B", "1")
		pass = strings.ReplaceAll(pass, "F", "0")
		pass = strings.ReplaceAll(pass, "R", "1")
		pass = strings.ReplaceAll(pass, "L", "0")

		passes = append(passes, pass)
	}
	return passes
}

func readInput() []string {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)

	passes := []string{}
	for scanner.Scan() {
		passes = append(passes, scanner.Text())
	}

	return passes
}
