package main

import (
	"fmt"

	"aoccommon"
)

func test1() {
	words := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	checksum := compute_checksum(words)
	if checksum != 12 {
		panic(checksum)
	}
	fmt.Println("OK")
}

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	checksum := compute_checksum(lines)
	fmt.Println(checksum)
}
