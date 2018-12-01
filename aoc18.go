package main

import (
	"os"
	"strings"
)

var mapping = map[string]func(){}

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
