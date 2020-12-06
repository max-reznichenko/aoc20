package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func newPassport(p string) *passport {
	p = strings.ReplaceAll(p, "\n", " ")
	pFields := strings.Split(p, " ")
	pass := passport{}

	for _, field := range pFields {
		kv := strings.Split(field, ":")

		if kv[0] == "byr" {
			pass.byr = kv[1]
		}
		if kv[0] == "iyr" {
			pass.iyr = kv[1]
		}
		if kv[0] == "eyr" {
			pass.eyr = kv[1]
		}
		if kv[0] == "hgt" {
			pass.hgt = kv[1]
		}
		if kv[0] == "hcl" {
			pass.hcl = kv[1]
		}
		if kv[0] == "ecl" {
			pass.ecl = kv[1]
		}
		if kv[0] == "pid" {
			pass.pid = kv[1]
		}
		if kv[0] == "cid" {
			pass.cid = kv[1]
		}
	}
	return &pass
}

func (p *passport) isValid() bool {
	if p.isValidByr() && p.isValidIyr() && p.isValidEyr() &&
		p.isValidHgt() &&
		p.isValidHcl() && p.isValidEcl() && p.isValidPid() {
		return true
	}
	return false
}

func (p *passport) isValidByr() bool {
	byr, err := strconv.Atoi(p.byr)
	if err == nil && byr >= 1920 && byr <= 2002 {
		return true
	}
	return false
}

func (p *passport) isValidIyr() bool {
	iyr, err := strconv.Atoi(p.iyr)
	if err == nil && iyr >= 2010 && iyr <= 2020 {
		return true
	}
	return false
}

func (p *passport) isValidEyr() bool {
	eyr, err := strconv.Atoi(p.eyr)
	if err == nil && eyr >= 2020 && eyr <= 2030 {
		return true
	}
	return false
}

func (p *passport) isValidHgt() bool {
	if strings.HasSuffix(p.hgt, "cm") {
		v := strings.ReplaceAll(p.hgt, "cm", "")
		hv, err := strconv.Atoi(v)

		if err == nil && hv >= 150 && hv <= 193 {
			return true
		}
	}
	if strings.HasSuffix(p.hgt, "in") {
		v := strings.ReplaceAll(p.hgt, "in", "")
		hv, err := strconv.Atoi(v)

		if err == nil && hv >= 59 && hv <= 76 {
			return true
		}
	}

	return false
}

func (p *passport) isValidHcl() bool {
	if strings.HasPrefix(p.hcl, "#") {
		hv := strings.ReplaceAll(p.hcl, "#", "")
		re := regexp.MustCompile(`^[0123456789abcdef]{6}$`)
		if re.Match([]byte(hv)) {
			return true
		}
	}
	return false
}

func (p *passport) isValidEcl() bool {
	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}

}

func (p *passport) isValidPid() bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	if re.Match([]byte(p.pid)) {
		return true
	}
	return false
}

func main() {
	var count int
	for _, passInput := range readInput() {
		p := newPassport(passInput)
		if p.isValid() {
			count++
		}
	}

	fmt.Println(count)
}

func readInput() []string {
	input, _ := ioutil.ReadFile("../input")
	return strings.Split(string(input), "\n\n")
}
