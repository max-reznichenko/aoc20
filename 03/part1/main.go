package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()

	var idx, count int
	lenx := len(input[0])
	for _, row := range input {
		if string(row[idx%lenx]) == "#" {
			count++
		}
		idx += 3
	}

	fmt.Println(count)
}

func readInput() []string {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)

	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows
}
