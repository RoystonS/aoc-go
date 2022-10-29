package main

import "aoccommon"

// Counts the number of times each letter (rune) appears in a string
func count_letters(s string) map[rune]int {
	result := map[rune]int{}

	for _, r := range s {
		result[r] = result[r] + 1
	}

	return result
}

// Calculates the AOC 2018 day 2 checksum for a series of strings
func compute_checksum(words []string) (checksum int) {
	twos := 0
	threes := 0

	for _, word := range words {
		counts := count_letters(word)
		letters_by_count := aoccommon.PivotMap(counts)
		if _, ok := letters_by_count[2]; ok {
			twos++
		}
		if _, ok := letters_by_count[3]; ok {
			threes++
		}
	}

	return twos * threes
}
