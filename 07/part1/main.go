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
	counter := 0
	for color := range bags {
		if containsShinyBag(bags, color) {
			counter++
		}
	}
	fmt.Println(counter)
}

func containsShinyBag(bags map[string]map[string]int, color string) bool {
	// 1. find color definition
	// 2. get color children
	// 3. if no children => return false
	// 4. if any of children = shiny => return true
	// 5. ELSE run self for each of children
	for cColor := range bags[color] {
		if cColor == "shiny gold" || containsShinyBag(bags, cColor) {
			return true
		}
	}
	return false
}

func parseInput() map[string]map[string]int {
	bags := make(map[string]map[string]int)

	zeroRe := regexp.MustCompile(`(.*) bags contain no other bags.`)
	manyRe := regexp.MustCompile(`(.*) bags contain (.*)`)

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
			re := regexp.MustCompile(`(\d+) (.*) bag`)
			for _, v := range strings.Split(string(subs[0][2]), ", ") {
				ssubs := re.FindAllSubmatch([]byte(v), -1)
				bcount, _ := strconv.Atoi(string(ssubs[0][1]))
				bags[string(subs[0][1])][string(ssubs[0][2])] = bcount
			}
		}
	}

	return bags
}
