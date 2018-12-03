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
