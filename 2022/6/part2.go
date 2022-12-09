package main

import (
	"aoccommon"
	"fmt"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := computePart2(lines[0])
	fmt.Printf("%d\n", result)
}
