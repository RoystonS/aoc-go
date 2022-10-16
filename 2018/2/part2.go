package main

import (
	"fmt"
	"strings"

	"aoccommon"
)

func are_one_character_different(s1 string, s2 string) bool {
	found_difference := false

	s2runes := []rune(s2)
	for index, ch1 := range s1 {
		if ch1 != s2runes[index] {
			if found_difference {
				return false
			} else {
				found_difference = true
			}
		}
	}

	return found_difference
}

func get_common_characters(s1 string, s2 string) string {
	var sb strings.Builder

	s2runes := []rune(s2)
	for index, ch1 := range s1 {
		if ch1 == s2runes[index] {
			sb.WriteRune(ch1)
		}
	}

	return sb.String()
}

func test2() {
	if !are_one_character_different("axc", "abc") {
		panic("axc")
	}
	if are_one_character_different("abcd", "abdc") {
		panic("abcd")
	}
	if get_common_characters("abfde", "abjde") != "abde" {
		panic("get-common")
	}

	fmt.Println("OK")
}

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	for index, line1 := range lines {
		for _, line2 := range lines[index+1:] {
			if are_one_character_different(line1, line2) {
				fmt.Println(get_common_characters(line1, line2))
			}
		}
	}
}
