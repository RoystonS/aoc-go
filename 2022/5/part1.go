package main

import (
	"aoccommon"
	"fmt"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := computePart1(lines)
	fmt.Println(result)
}
