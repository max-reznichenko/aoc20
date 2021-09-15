package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seatMap := readInput()
	flipMap(seatMap, "#")

	var occupied int
	for _, row := range seatMap {
		for _, seat := range row {
			if seat == "#" {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
}

func flipMap(seatMap [][]string, flipTo string) {
	var flipCoordinates [][]int

	for y, row := range seatMap {
		for x, seat := range row {
			var adjSeats []string

			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					if i == 0 && j == 0 {
						continue
					}
					adjSeats = append(adjSeats, getSeat(seatMap, x+i, y+j))
				}
			}

			var occupied int
			for _, v := range adjSeats {
				if v == "#" {
					occupied++
				}
			}

			if (seat == "L" && occupied == 0) || (seat == "#" && occupied >= 4) {
				flipCoordinates = append(flipCoordinates, []int{x, y})
			}
		}
	}

	for _, xy := range flipCoordinates {
		seatMap[xy[1]][xy[0]] = flipTo
	}

	if len(flipCoordinates) > 0 {
		if flipTo == "#" {
			flipMap(seatMap, "L")
		} else {
			flipMap(seatMap, "#")
		}
	}
}

// return a seat state [L, #, .]
// if seat is out of bonds, returns an empty string
func getSeat(seatMap [][]string, x, y int) string {
	defer func() {
		recover()
	}()

	return seatMap[y][x]
}

func readInput() [][]string {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)

	var seatMap [][]string

	for scanner.Scan() {
		inputRow := scanner.Text()
		var row []string
		for _, s := range strings.Split(inputRow, "") {
			row = append(row, s)
		}

		seatMap = append(seatMap, row)
	}

	return seatMap
}
