package main

import (
	"aoccommon"
	"strconv"
	"strings"
)

type Assignment struct {
	from int
	to   int
}

type Pair struct {
	elf1 Assignment
	elf2 Assignment
}

func (a1 Assignment) covers(a2 Assignment) bool {
	return a1.from <= a2.from && a1.to >= a2.to
}

func (a1 Assignment) overlaps(a2 Assignment) bool {
	// The logic is easier to imagine in reverse
	doesNotOverlap := (a1.to < a2.from || a2.to < a1.from)
	return !doesNotOverlap
}

func (pair Pair) containsCover() bool {
	return pair.elf1.covers(pair.elf2) || pair.elf2.covers(pair.elf1)
}

func parseAssignment(text string) Assignment {
	bits := strings.Split(text, "-")

	from, err := strconv.Atoi(bits[0])
	aoccommon.CheckError(err)
	to, err := strconv.Atoi(bits[1])
	aoccommon.CheckError(err)

	return Assignment{from, to}
}

func parseLine(line string) Pair {
	elfStrings := strings.Split(line, ",")
	elf1 := parseAssignment(elfStrings[0])
	elf2 := parseAssignment(elfStrings[1])

	return Pair{elf1, elf2}
}

func computePart1(lines []string) int {
	total := 0

	for _, line := range lines {
		pair := parseLine(line)
		if pair.containsCover() {
			total++
		}
	}

	return total
}

func computePart2(lines []string) int {
	total := 0

	for _, line := range lines {
		pair := parseLine(line)
		if pair.elf1.overlaps(pair.elf2) {
			total++
		}
	}

	return total
}
