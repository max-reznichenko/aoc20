package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bags := parseInput()
	fmt.Println(countBags("shiny gold", bags, 1) - 1)
}

func countBags(color string, bags map[string]map[string]int, parentCount int) int {
	currentCount := parentCount

	for cColor, count := range bags[color] {
		currentCount += parentCount * countBags(cColor, bags, count)
	}

	return currentCount
}

func parseInput() map[string]map[string]int {
	bags := make(map[string]map[string]int)

	zeroRe := regexp.MustCompile(`(.*) bags contain no other bags.`)
	manyRe := regexp.MustCompile(`(.*) bags contain (.*)`)
	subRe := regexp.MustCompile(`(\d+) (.*) bag`)

	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := []byte(scanner.Text())
		if zeroRe.Match(text) {
			subs := zeroRe.FindSubmatch(text)
			bags[string(subs[1])] = nil
		} else if manyRe.Match([]byte(scanner.Text())) {
			subs := manyRe.FindAllSubmatch(text, -1)
			bags[string(subs[0][1])] = make(map[string]int)

			for _, v := range strings.Split(string(subs[0][2]), ", ") {
				ssubs := subRe.FindAllSubmatch([]byte(v), -1)
				bcount, _ := strconv.Atoi(string(ssubs[0][1]))
				bags[string(subs[0][1])][string(ssubs[0][2])] = bcount
			}
		}
	}

	return bags
}
