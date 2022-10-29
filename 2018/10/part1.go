package main

import (
	"aoccommon"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	runMovement(lines, true)
}
