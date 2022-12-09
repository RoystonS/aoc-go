package main

import (
	"aoccommon"
	"fmt"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := computePart1(lines[0])
	fmt.Printf("%d\n", result)
}
