package main

import (
	"aoccommon"
)

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	runMovement(lines, false)
}
