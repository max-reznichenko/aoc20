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
		group = strings.ReplaceAll(group, "\n", "")
		counter += len(unique([]byte(group)))
	}
	fmt.Println(counter)
}

func readInput() []string {
	data, _ := ioutil.ReadFile("../input")
	return strings.Split(string(data), "\n\n")
}

// https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(byteSlice []byte) []byte {
	keys := make(map[byte]bool)
	list := []byte{}
	for _, entry := range byteSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
