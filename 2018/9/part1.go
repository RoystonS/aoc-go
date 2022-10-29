package main

import (
	"aoccommon"
	"fmt"
	"regexp"
	"strconv"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	// line is of the form 441 players; last marble is worth 71032 points
	re := regexp.MustCompile(`(\d+) players.* worth (\d+) points`)
	matches := re.FindAllStringSubmatch(lines[0], -1)
	players, _ := strconv.Atoi(matches[0][1])
	lastMarble, _ := strconv.Atoi(matches[0][2])

	highestScore := runGame(players, lastMarble)
	fmt.Println(highestScore)
}
