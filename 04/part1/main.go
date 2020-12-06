package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var count int
	for _, passport := range readInput() {
		if isValid(passport) {
			count++
		}
	}

	fmt.Println(count)
}

func isValid(passport string) bool {
	checks := []string{
		"byr:",
		"iyr:",
		"eyr:",
		"hgt:",
		"hcl:",
		"ecl:",
		"pid:",
		// "cid",
	}

	for _, check := range checks {
		if !strings.Contains(passport, check) {
			return false
		}
	}
	return true
}

func readInput() []string {
	input, _ := ioutil.ReadFile("../input")
	return strings.Split(string(input), "\n\n")
}
