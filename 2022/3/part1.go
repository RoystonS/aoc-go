package main

import (
	"aoccommon"
	"fmt"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	fmt.Printf("%d\n", computePart1(lines))
}
