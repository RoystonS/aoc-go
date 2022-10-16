package main

import (
	"fmt"
	"strconv"

	"aoccommon"
)

// Finds the guard who was asleep the most
func findGuardTheMostAsleep(guardInfos map[int]map[int]int) (guardNumber int, totalMinutesAsleep int) {
	highestMinutes := -1
	highestMinutesGuard := -1

	for guardNumber, guardInfo := range guardInfos {
		differentMinutesAsleep := 0

		for _, count := range guardInfo {
			differentMinutesAsleep += count
		}

		if differentMinutesAsleep > highestMinutes {
			highestMinutes = differentMinutesAsleep
			highestMinutesGuard = guardNumber
		}
	}

	return highestMinutesGuard, highestMinutes
}

func computeStrategy1(lines []string) (guardNumber int, mostAsleepMinute int) {
	guardInfos := computeUsage(lines)
	mostAsleepGuardNumber, _ := findGuardTheMostAsleep(guardInfos)
	guardInfo := guardInfos[mostAsleepGuardNumber]

	highestMinute, _ := maxValue(guardInfo)

	return mostAsleepGuardNumber, highestMinute
}

func test1() {
	if guardNumber, highestMinute := computeStrategy1(testData); guardNumber*highestMinute != 240 {
		panic("unexpected " + strconv.Itoa(guardNumber*highestMinute))
	}
	fmt.Println("OK")
}

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	guardNumber, highestMinute := computeStrategy1(lines)
	fmt.Println(guardNumber * highestMinute)
}

func maxValue(values map[int]int) (key int, value int) {
	highestValue := -100000
	highestKey := -1

	for key, value := range values {
		if value > highestValue {
			highestKey = key
			highestValue = value
		}
	}

	return highestKey, highestValue
}
