package main

import (
	"aoccommon"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/zyedidia/generic/list"
)

type Item struct {
	worryLevel int
}

type Monkey struct {
	items            *list.List[Item]
	operation        func(initialWorryLevel int) int
	testDivisibleBy  int
	trueDestination  int
	falseDestination int
}

func parseMonkeys(lines []string) []*Monkey {
	startingItemsRe := regexp.MustCompile(`Starting items: (.*)`)
	operationRe := regexp.MustCompile(`Operation: new = old (.) (.*)`)
	testRe := regexp.MustCompile(`Test: divisible by (\d+)`)
	throwRe := regexp.MustCompile(`throw to monkey (\d+)`)

	lineIndex := 0

	monkeyCount := 0
	monkeys := list.New[*Monkey]()

	for lineIndex < len(lines) {
		monkeyCount++
		monkey := &Monkey{items: list.New[Item]()}
		monkeys.PushBack(monkey)

		// Monkey 0:
		lineIndex++

		// Starting items: 3, 5, 2, 1
		startingItemsMatch := startingItemsRe.FindAllStringSubmatch(lines[lineIndex], -1)
		for _, itemWorryLevelString := range strings.Split(startingItemsMatch[0][1], ", ") {
			initialWorryLevel, _ := strconv.Atoi(itemWorryLevelString)
			monkey.items.PushBack(Item{worryLevel: initialWorryLevel})
		}
		lineIndex++

		// operation: new = old +|* x
		operationMatch := operationRe.FindAllStringSubmatch(lines[lineIndex], -1)
		operator := operationMatch[0][1]
		operand := operationMatch[0][2]

		operation := func(old int) int {
			otherValue, err := strconv.Atoi(operand)
			if err != nil {
				// Failed to parse number => 'old'
				otherValue = old
			}
			switch operator {
			case "+":
				return old + otherValue
			case "*":
				return old * otherValue
			}
			panic("Unexpected operator: " + operator)
		}
		monkey.operation = operation
		lineIndex++

		// Test: divisible by 19
		testMatch := testRe.FindAllStringSubmatch(lines[lineIndex], -1)
		divisor, err := strconv.Atoi(testMatch[0][1])
		aoccommon.CheckError(err)
		monkey.testDivisibleBy = divisor
		lineIndex++

		throwMatch := throwRe.FindAllStringSubmatch(lines[lineIndex], -1)
		target, err := strconv.Atoi(throwMatch[0][1])
		aoccommon.CheckError(err)
		monkey.trueDestination = target
		lineIndex++

		throwMatch = throwRe.FindAllStringSubmatch(lines[lineIndex], -1)
		target, err = strconv.Atoi(throwMatch[0][1])
		aoccommon.CheckError(err)
		monkey.falseDestination = target
		lineIndex++

		lineIndex++
	}

	monkeyArray := aoccommon.ToArray(monkeys, monkeyCount)
	return monkeyArray
}

func run(lines []string, roundCount int, doReduceWorryLevel bool) int {
	monkeys := parseMonkeys(lines)
	inspectionCounts := make([]int, len(monkeys))

	runMonkeyRounds(monkeys, inspectionCounts, roundCount, doReduceWorryLevel)

	// for monkeyIndex, count := range inspectionCounts {
	// 	fmt.Printf("Monkey %d inspected items %d times.\n", monkeyIndex, count)
	// }

	sort.Ints(inspectionCounts)
	return inspectionCounts[len(inspectionCounts)-2] * inspectionCounts[len(inspectionCounts)-1]
}

func runMonkeyRounds(monkeys []*Monkey, inspectionCounts []int, roundCount int, doReduceWorryLevel bool) {

	// For part 2, the numbers we'll be dealing with are _HUGE_, so instead of
	// keeping their full numbers, we'll just keep them modulo some value.
	// Because the tests are whether we're divisible by some number, we can keep
	// everything modulo that number. BUT each monkey has a different value,
	// so we need to use the product of all the monkeys' divisors.
	// For example, with a full value of 1501, 1501 % 17 == 5, 1501 % 19 == 0;
	// 17 * 19 == 323;  1501 % 323 == 209; 209 % 17 == 5, 209 % 19 == 0
	// i.e. we can do all of the arithmetic (for part 2) modulo the product of all the test divisors
	// (This wouldn't work for part 1 as that's doing rounded division by 3)
	part2CommonModulus := 1
	for _, monkey := range monkeys {
		part2CommonModulus *= monkey.testDivisibleBy
	}

	for round := 0; round < roundCount; round++ {
		for monkeyId, monkey := range monkeys {
			itemNode := monkey.items.Front
			for itemNode != nil {
				nextNode := itemNode.Next

				oldWorryLevel := itemNode.Value.worryLevel
				newWorryLevel := monkey.operation(oldWorryLevel)

				inspectionCounts[monkeyId] += 1
				// fmt.Printf("Monkey %d, item worry level %d -> %d\n", monkeyId, oldWorryLevel, newWorryLevel)

				if doReduceWorryLevel {
					newWorryLevel = newWorryLevel / 3
				} else {
					newWorryLevel = newWorryLevel % part2CommonModulus
				}
				itemNode.Value.worryLevel = newWorryLevel

				targetMonkey := monkey.falseDestination
				if (newWorryLevel % monkey.testDivisibleBy) == 0 {
					// It is divisible
					targetMonkey = monkey.trueDestination
				}

				// This monkey is throwing the item
				monkey.items.Remove(itemNode)

				// To the target
				monkeys[targetMonkey].items.PushBack(itemNode.Value)
				// fmt.Printf("Monkey %d throwing item with worry level %d to monkey %d\n", monkeyId, newWorryLevel, targetMonkey)
				itemNode = nextNode
			}
		}

		// dumpMonkeys(monkeys, round+1)
	}
}

func dumpMonkeys(monkeys *[]*Monkey, round int) {
	fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", round)
	for monkeyIndex, monkey := range *monkeys {
		fmt.Printf("Monkey %d: ", monkeyIndex)
		for item := range aoccommon.IterateList(monkey.items) {
			fmt.Printf("%d ", item.worryLevel)
		}
		fmt.Println()
	}

}
func computePart1(lines []string) int {
	return run(lines, 20, true)
}

func computePart2(lines []string) int {
	return run(lines, 10000, false)
}
