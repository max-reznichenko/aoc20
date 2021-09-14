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
	var invalidId int

	for i := 0; i < len(input)-preambleLen; i++ {
		window = input[i : i+preambleLen]
		checkVal = input[i+preambleLen]

		if !check(window, checkVal) {
			invalidId = i + preambleLen + 1
			break
		}
	}

	checkValInt, _ := strconv.Atoi(checkVal)

	var intInput []int
	for _, v := range input {
		iv, _ := strconv.Atoi(v)
		intInput = append(intInput, iv)
	}

	for i := 0; i < invalidId; i++ {
		for j := i; j < invalidId; j++ {
			seq := intInput[i:j]

			if len(seq) < 2 {
				continue
			}

			var sum int
			for _, v := range seq {
				sum += v
			}

			if sum == checkValInt {
				min, max := MinMax(seq)
				fmt.Println(min + max)
				os.Exit(0)
			}
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

// copied from https://stackoverflow.com/a/45976758
func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
