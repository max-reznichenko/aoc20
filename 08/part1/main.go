package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	op  string
	arg int
	vis int
}

func main() {
	var acc, si int
	var cmd *Cmd
	cmds := readInput()

	for {
		cmd = cmds[si]

		if cmd.vis > 0 {
			fmt.Println(acc)
			os.Exit(0)
		}

		cmd.vis++
		switch cmd.op {
		case "nop":
			si++
		case "acc":

			acc += cmd.arg
			si++
		case "jmp":
			si += cmd.arg
		default:
			fmt.Println("should not happen")
		}

	}
}

func readInput() []*Cmd {
	file, _ := os.Open("../input")
	scanner := bufio.NewScanner(file)

	var cmds []*Cmd

	for scanner.Scan() {
		ins := strings.Split(scanner.Text(), " ")
		arg, _ := strconv.Atoi(ins[1])
		cmds = append(cmds, &Cmd{
			op:  ins[0],
			arg: arg,
		})
	}

	return cmds
}
