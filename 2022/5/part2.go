package main

import (
	"aoccommon"
	"fmt"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := computePart2(lines)
	fmt.Println(result)
}
