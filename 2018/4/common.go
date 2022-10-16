package main

import (
	"aoccommon"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// File format:
// [1518-11-12 00:00] Guard #3011 begins shift
// [1518-05-21 00:32] wakes up
// [1518-11-12 00:30] falls asleep

var (
	testData = []string{"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep", "[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep", "[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep", "[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep", "[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep", "[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep", "[1518-11-05 00:55] wakes up"}
)

// Converts a list of guard statements into a map from guard number
// to a map from minute number to the number of times the guard is asleep
// for that minute
func computeUsage(lines []string) map[int]map[int]int {
	sort.Strings(lines)

	dateRegexp := regexp.MustCompile(`^\[(\d+)-(\d+)-(\d+) (\d+):(\d+)\] `)
	guardRegexp := regexp.MustCompile(`Guard #(\d+) begins shift`)

	currentGuard := -1
	sleepStartedAt := -1

	// A map from guard number to a mapping from minute number to usage of it
	guardInfos := map[int]map[int]int{}

	var guardInfo map[int]int

	for _, line := range lines {
		dateTimeMatches := dateRegexp.FindAllStringSubmatch(line, -1)

		minutes, err := strconv.Atoi(dateTimeMatches[0][5])
		aoccommon.CheckError(err)

		guardMatches := guardRegexp.FindAllStringSubmatch(line, -1)
		if len(guardMatches) > 0 {
			currentGuard, err = strconv.Atoi(guardMatches[0][1])
			aoccommon.CheckError(err)

			sleepStartedAt = -1
			guardInfo = guardInfos[currentGuard]
			if guardInfo == nil {
				guardInfo = map[int]int{}
				guardInfos[currentGuard] = guardInfo
			}
		}

		if strings.HasSuffix(line, "falls asleep") {
			sleepStartedAt = minutes
		}

		if strings.HasSuffix(line, "wakes up") {
			for min := sleepStartedAt; min < minutes; min++ {
				guardInfo[min]++
			}
		}
	}

	return guardInfos
}
