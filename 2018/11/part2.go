package main

import (
	"aoccommon"
	"fmt"
	"strconv"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	serial, _ := strconv.Atoi(lines[0])
	x, y, _, size := computePart2(serial)
	fmt.Printf("%d,%d,%d\n", x, y, size)
}
