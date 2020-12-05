package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type password struct {
	min    int
	max    int
	letter string
	value  string
}

func (p *password) isValid() bool {
	re := regexp.MustCompile(`` + p.letter + ``)
	matchCount := len(re.FindAllSubmatch([]byte(p.value), -1))

	if matchCount >= p.min && matchCount <= p.max {
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
		min, _ := strconv.Atoi(v[1])
		max, _ := strconv.Atoi(v[2])

		passwords = append(passwords, password{min: min, max: max, letter: v[3], value: v[4]})
	}

	return passwords
}
