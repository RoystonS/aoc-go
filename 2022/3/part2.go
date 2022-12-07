package main

import (
	"aoccommon"
	"fmt"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	fmt.Printf("%d\n", computePart2(lines))
}
