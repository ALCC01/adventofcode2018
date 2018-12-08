package main

import (
	"aoc18/challenges"
	"os"
	"strings"
)

var mapping = map[string]func(){
	"d1c1": challenges.D1C1,
	"d1c2": challenges.D1C2,
	"d2c1": challenges.D2C1,
	"d2c2": challenges.D2C2,
	"d3c1": challenges.D3C1,
	"d3c2": challenges.D3C2,
	"d4c1": challenges.D4C1,
	"d4c2": challenges.D4C2,
	"d5c1": challenges.D5C1,
	"d5c2": challenges.D5C2,
	"d6c1": challenges.D6C1,
	"d6c2": challenges.D6C2,
	"d7c1": challenges.D7C1,
	"d8c1": challenges.D8C1,
	"d8c2": challenges.D8C2,
}

const helpMessage = `Advent of Code 2018

Usage: aoc18 [challenge]
Example: aoc18 d1c2 < ./input.txt

Copyright (c) 2018 Alberto Coscia - Released under the Unlicense <https://unlicense.org/UNLICENSE>`

func main() {
	if len(os.Args) == 1 {
		help()
	} else {
		challenge := strings.ToLower(os.Args[1])
		if mapping[challenge] != nil {
			mapping[os.Args[1]]()
		} else {
			help()
		}
	}
}

func help() {
	println(helpMessage)
}
