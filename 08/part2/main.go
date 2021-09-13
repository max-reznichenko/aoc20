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
	for i, cmd := range readInput() {
		cmds := readInput() // reread to avoid conflicts

		switch cmd.op {
		case "acc":
			continue
		case "jmp":
			cmds[i].op = "nop"
		case "nop":
			cmds[i].op = "jmp"
		default:
			fmt.Println("should not happen")
			os.Exit(1)
		}

		if acc, ok := calcAcc(cmds); ok {
			fmt.Println(acc)
			os.Exit(0)
		}
	}
}

func calcAcc(cmds []*Cmd) (int, bool) {
	var acc, si int
	var cmd *Cmd

	for {
		cmd = cmds[si]

		if cmd.vis > 0 {
			return acc, false
		}

		cmd.vis++
		switch cmd.op {
		case "nop":
			if si+1 == len(cmds) {
				return acc, true
			}
			si++
		case "acc":
			acc += cmd.arg
			if si+1 == len(cmds) {
				return acc, true
			}
			si++
		case "jmp":
			if si+1 == len(cmds) {
				return acc, true
			}
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
