package main

import (
	"fmt"

	"aoccommon"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)
	fmt.Printf("part2; lines: %d\n", len(lines))
}
