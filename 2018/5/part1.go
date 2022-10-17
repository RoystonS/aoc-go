package main

import (
	"aoccommon"
	"fmt"
)

func test1() {
	polymerList, _ := fullyReact("dabAcCaCBAcCcaDA")
	value := runeListToString(polymerList)
	if value != "dabCBAcaDA" {
		panic("Test failed " + value)
	}

	polymerList, _ = fullyReact("aa")
	value = runeListToString(polymerList)
	if value != "aa" {
		panic("Test failed " + value)
	}

	fmt.Println("OK")
}

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	_, polymerLength := fullyReact(lines[0])
	fmt.Println(polymerLength)
}
