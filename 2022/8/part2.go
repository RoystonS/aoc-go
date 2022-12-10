package main

import (
	"aoccommon"
	"fmt"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	// other()
	result := computePart2(lines)
	fmt.Printf("%d\n", result)
}
