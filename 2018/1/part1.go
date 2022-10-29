package main

import (
	"aoccommon"
	"fmt"
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
