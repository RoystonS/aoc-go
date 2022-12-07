package main

func getPriority(value rune) int {
	if value >= 'a' && value <= 'z' {
		return 1 + int(value-'a')
	}
	if value >= 'A' && value <= 'Z' {
		return 27 + int(value-'A')
	}
	panic("Unexpected char")
}

func toSet(value string) map[rune]bool {
	m := map[rune]bool{}

	for _, shChar := range value {
		m[shChar] = true
	}

	return m
}

func computePart1(lines []string) int {
	total := 0

	for _, line := range lines {
		// fmt.Printf("Line: %s\n", line)

		lineLength := len(line)
		halfWay := lineLength / 2
		firstHalf := line[0:halfWay]
		secondHalf := line[halfWay:]

		secondHalfMap := toSet(secondHalf)

		for _, fhChar := range firstHalf {
			if secondHalfMap[fhChar] {
				priority := getPriority(fhChar)
				// fmt.Printf("priority for %s %c is %d\n", line, fhChar, priority)
				total += priority
				break
			}
		}
	}
	return total
}

func computePart2(lines []string) int {
	total := 0

	for i := 0; i < len(lines); i += 3 {
		groupLines := lines[i : i+3]

		secondLineSet := toSet(groupLines[1])
		thirdLineSet := toSet(groupLines[2])

		for _, ch := range groupLines[0] {
			if secondLineSet[ch] && thirdLineSet[ch] {
				priority := getPriority(ch)
				total += priority
				break
			}
		}
	}

	return total
}
