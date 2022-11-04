package main

import (
	"aoccommon"
	"fmt"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	g := parseLines(lines, 26)
	value := tsortToString(g)
	fmt.Println(value)
}
