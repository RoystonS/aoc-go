package main

import (
	"fmt"

	"github.com/RoystonS/aoc-go/aoccommon"
)

func part1() {
	nums, err := aoccommon.ReadNumbers("input")
	aoccommon.CheckError(err)

	frequency := 0

	for _, num := range nums {
		frequency += num
	}
	fmt.Println(frequency)
}
