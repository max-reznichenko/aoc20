package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type password struct {
	pos1   int
	pos2   int
	letter string
	value  string
}

func (p *password) isValid() bool {
	var m int
	if string(p.value[p.pos1-1]) == p.letter {
		m++
	}
	if string(p.value[p.pos2-1]) == p.letter {
		m++
	}
	if m == 1 {
		return true
	}

	return false
}

func main() {
	passwords := readInput()

	var validCounter int
	for _, password := range passwords {
		if password.isValid() {
			validCounter++
		}
	}

	fmt.Println(validCounter)
}

func readInput() []password {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(\d+)-(\d+) ([[:alpha:]]): ([[:alpha:]]+)`)

	passwords := []password{}
	for scanner.Scan() {
		v := re.FindStringSubmatch(scanner.Text())
		pos1, _ := strconv.Atoi(v[1])
		pos2, _ := strconv.Atoi(v[2])

		passwords = append(passwords, password{pos1: pos1, pos2: pos2, letter: v[3], value: v[4]})
	}

	return passwords
}
