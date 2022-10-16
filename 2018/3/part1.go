package main

import (
	"aoccommon"
	"fmt"
)

func part1() {
	claims, err := read_claims("input")
	aoccommon.CheckError(err)

	tiles := apply_claims(claims)

	// Count the number of tiles with >= 2 claims
	count := 0
	for y := range tiles {
		for x := range tiles[y] {
			if tiles[x][y] >= 2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
