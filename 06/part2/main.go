package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	groups := readInput()

	var counter int
	for _, group := range groups {
		counter += count(group)
	}
	fmt.Println(counter)
}

func readInput() []string {
	data, _ := ioutil.ReadFile("../input")
	return strings.Split(string(data), "\n\n")
}

func count(group string) int {
	keys := make(map[byte]int)
	people := strings.Split(group, "\n")

	for _, person := range people {
		// byting a person to produce an answer #hm #not_bad #not_bad_at_all
		for _, answer := range []byte(person) {
			keys[answer]++
		}
	}

	var counter int
	for _, value := range keys {
		if value == len(people) {
			counter++
		}
	}
	return counter
}
