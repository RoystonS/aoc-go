package main

import (
	"aoccommon"
	"regexp"
	"strconv"
)

func getStackCount(lines []string) (int, int) {
	stackCountRegexp := regexp.MustCompile(`(\d+)\s*$`)

	for lineIndex, line := range lines {
		matches := stackCountRegexp.FindAllStringSubmatch(line, -1)

		if matches != nil {
			match := matches[0]
			count, err := strconv.Atoi(match[len(match)-1])
			aoccommon.CheckError(err)
			return count, lineIndex
		}
	}

	panic("Could not find count of stacks")
}

type Movement struct {
	count     int
	fromStack int
	toStack   int
}

func parseMovement(line string) Movement {
	movementRegex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := movementRegex.FindAllStringSubmatch(line, -1)

	count, _ := strconv.Atoi(matches[0][1])
	fromStackNumber, _ := strconv.Atoi(matches[0][2])
	toStackNumber, _ := strconv.Atoi(matches[0][3])

	return Movement{
		count:     count,
		fromStack: fromStackNumber - 1,
		toStack:   toStackNumber - 1,
	}
}

type CrateMover func(fromStack *aoccommon.Stack[rune], toStack *aoccommon.Stack[rune], count int)

func crateMover9000(fromStack *aoccommon.Stack[rune], toStack *aoccommon.Stack[rune], count int) {
	for i := 0; i < count; i++ {
		item, hasValue := fromStack.Pop()
		if !hasValue {
			panic("No item to move")
		}
		toStack.Push(item)
	}
}

func crateMover9001(fromStack *aoccommon.Stack[rune], toStack *aoccommon.Stack[rune], count int) {
	tempStack := aoccommon.NewStack[rune]()
	crateMover9000(fromStack, tempStack, count)
	crateMover9000(tempStack, toStack, count)
}

func computePart1(lines []string) string {
	return compute(lines, crateMover9000)
}

func computePart2(lines []string) string {
	return compute(lines, crateMover9001)
}

func compute(lines []string, crateMover CrateMover) string {
	// How many stacks are there?
	stackCount, stacksLineIndex := getStackCount(lines)

	// Make empty stacks
	stacks := make([]*aoccommon.Stack[rune], stackCount)
	for i, _ := range stacks {
		stacks[i] = aoccommon.NewStack[rune]()
	}

	// Populate the initial stacks
	for i := stacksLineIndex - 1; i >= 0; i-- {
		line := lines[i]

		runes := []rune(line)

		// Pick out the chars
		// [N] [C]
		// They're 4 apart, starting at index 1
		for stackIndex := 0; stackIndex < stackCount; stackIndex++ {
			r := runes[1+stackIndex*4]
			if r != ' ' {
				stacks[stackIndex].Push(r)
			}
		}
	}

	// Run the movements
	for _, line := range lines[(stacksLineIndex + 2):] {
		movement := parseMovement(line)

		crateMover(stacks[movement.fromStack], stacks[movement.toStack], movement.count)
	}

	// Collect the stack tips
	finalStackHeads := make([]rune, stackCount)
	for stackIndex, stack := range stacks {
		item, hasValue := stack.Pop()
		if hasValue {
			finalStackHeads[stackIndex] = item
		} else {
			finalStackHeads[stackIndex] = ' '
		}
	}

	return string(finalStackHeads)
}
