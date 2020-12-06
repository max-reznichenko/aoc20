package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()
	fmt.Println(treeCount(input, 1, 1) * treeCount(input, 3, 1) * treeCount(input, 5, 1) * treeCount(input, 7, 1) * treeCount(input, 1, 2))
}

func treeCount(input []string, xSlope, ySlope int) int {
	var idx, count int
	lenx := len(input[0])
	for idy := 0; idy < len(input); idy += ySlope {
		row := input[idy]
		if string(row[idx%lenx]) == "#" {
			count++
		}
		idx += xSlope
	}

	return count
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
