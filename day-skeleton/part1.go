package main

import (
	"fmt"

	"aoccommon"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)
	fmt.Printf("part1; lines: %d\n", len(lines))
}
