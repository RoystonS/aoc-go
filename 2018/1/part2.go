package main

import (
	"fmt"

	"github.com/RoystonS/aoc-go/aoccommon"
)

func part2() {
	frequencies_seen := map[int]bool{}

	nums, err := aoccommon.ReadNumbers("input")
	aoccommon.CheckError(err)

	frequency := 0

	for {
		for _, num := range nums {
			frequency += num
			if frequencies_seen[frequency] {
				// First time we've seen this frequency twice. We're done.
				fmt.Println(frequency)
				return
			}
			frequencies_seen[frequency] = true
		}
	}
}
