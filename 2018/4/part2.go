package main

import (
	"fmt"
	"strconv"

	"aoccommon"
)

func computeStrategy2(lines []string) (guardNumber int, mostAsleepMinute int) {

	guardInfos := computeUsage(lines)

	maxTimes := -1
	maxTimesMinute := -1
	maxTimesGuardNumber := -1

	for guardNumber, guardInfo := range guardInfos {
		for minute, times := range guardInfo {
			if times > maxTimes {
				maxTimes = times
				maxTimesMinute = minute
				maxTimesGuardNumber = guardNumber
			}
		}
	}

	return maxTimesGuardNumber, maxTimesMinute
}

func test2() {
	if guardNumber, highestMinute := computeStrategy2(testData); guardNumber*highestMinute != 4455 {
		panic("unexpected " + strconv.Itoa(guardNumber*highestMinute))
	}
	fmt.Println("OK")
}

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	guardNumber, minute := computeStrategy2(lines)

	fmt.Println(guardNumber * minute)
}
