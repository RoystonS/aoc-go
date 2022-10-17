package main

import (
	"container/list"
	"strings"
	"unicode"
)

// Fully React an entire polymer, returning a list of the remaining units and the length
func fullyReact(line string) (*list.List, int) {
	list := list.New()

	length := 0

	// Add each input unit one at a time, checking if it
	// can be reacted with the previous unit
	for _, input := range line {
		list.PushBack(input)
		length++

		if length >= 2 {
			lastUnit := list.Back().Value.(rune)
			lastUnitButOne := list.Back().Prev().Value.(rune)
			// fmt.Println(lastUnitButOne, lastUnit)

			lastUnitPolarity := unicode.IsUpper(lastUnit)
			lastUnitButOnePolarity := unicode.IsUpper(lastUnitButOne)

			if lastUnitPolarity != lastUnitButOnePolarity && unicode.ToLower(lastUnit) == unicode.ToLower(lastUnitButOne) {
				list.Remove(list.Back())
				list.Remove(list.Back())
				length -= 2
			}
		}
	}

	return list, length
}

// Converts a list of runes into a string
func runeListToString(list *list.List) string {
	var sb strings.Builder

	elem := list.Front()
	for elem != nil {
		sb.WriteRune(elem.Value.(rune))
		elem = elem.Next()
	}
	return sb.String()
}
