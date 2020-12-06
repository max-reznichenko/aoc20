package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var maxSeatID int64
	for _, pass := range formatInput(readInput()) {
		sid := seatID(pass)
		if sid > maxSeatID {
			maxSeatID = sid
		}
	}

	fmt.Println(maxSeatID)
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
