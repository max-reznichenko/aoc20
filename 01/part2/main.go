package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	for _, x := range input {
		for _, y := range input {
			for _, z := range input {
				if x+y+z == 2020 {
					fmt.Println(x * y * z)
					os.Exit(1)
				}
			}
		}
	}
}

func readInput() []int {
	file, _ := os.Open("../input")
	defer file.Close()

	input := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	return input
}
