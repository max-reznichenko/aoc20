package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	const preambleLen = 25
	var window []string
	var checkVal string

	for i := 0; i < len(input)-preambleLen; i++ {
		window = input[i : i+preambleLen]
		checkVal = input[i+preambleLen]

		if !check(window, checkVal) {
			fmt.Println(checkVal)
			os.Exit(0)
		}
	}
}

func check(window []string, value string) bool {
	combinations := Combinations(window, 2)
	checkVal, _ := strconv.Atoi(value)

	for _, cmb := range combinations {
		val1, _ := strconv.Atoi(cmb[0])
		val2, _ := strconv.Atoi(cmb[1])

		if (val1 + val2) == checkVal {
			return true
		}
	}

	return false
}

func readInput() []string {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

// Copied from https://github.com/mxschmitt/golang-combinations/blob/master/combinations.go
//
// Combinations returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func Combinations(set []string, n int) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
