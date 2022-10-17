package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"aoccommon"
)

func getUniqueRunes(s string) map[rune]bool {
	uniqueRunes := map[rune]bool{}
	for _, unit := range s {
		uniqueRunes[unicode.ToLower(unit)] = true
	}
	return uniqueRunes
}

func removeAllOccurrencesOfRuneFromStringCaseInsensitively(s string, delChar rune) string {
	var sb strings.Builder

	upper := unicode.ToUpper(delChar)
	lower := unicode.ToLower(delChar)

	for _, rune := range s {
		if rune != upper && rune != lower {
			sb.WriteRune(rune)
		}
	}
	return sb.String()
}

func minimizeReaction(polymer string) int {
	uniqueUnits := getUniqueRunes(polymer)

	var wg sync.WaitGroup
	var mut sync.Mutex

	// Work through each unit type, to see which one we need to
	// remove to obtain the shortest reacted result
	smallestReactedPolymerLength := len(polymer)
	for unit := range uniqueUnits {
		wg.Add(1)

		unit := unit

		go func() {
			defer wg.Done()
			filteredPolymer := removeAllOccurrencesOfRuneFromStringCaseInsensitively(polymer, unit)
			_, reactedPolymerLength := fullyReact(filteredPolymer)

			mut.Lock()
			defer mut.Unlock()
			if reactedPolymerLength < smallestReactedPolymerLength {
				smallestReactedPolymerLength = reactedPolymerLength
			}
		}()
	}

	wg.Wait()

	return smallestReactedPolymerLength
}

func test2() {
	value := minimizeReaction("dabAcCaCBAcCcaDA")
	if value != 4 {
		panic("test failed " + strconv.Itoa(value))
	}
	fmt.Println("OK")
}

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	fmt.Println(minimizeReaction(lines[0]))
}
