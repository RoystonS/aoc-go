package main

import (
	"aoccommon"
	"fmt"
	"regexp"
	"strconv"

	"github.com/zyedidia/generic/list"
)

const PCRegister = "pc"

type Registers *map[string]int

type Instruction interface {
	execute(registers Registers)
}

type Noop struct {
}

func (op Noop) execute(registers Registers) {
	(*registers)[PCRegister] += 1
}

type Add struct {
	register string
	operand  int
}

func (op Add) execute(registers Registers) {
	value := (*registers)[op.register]
	value += op.operand
	(*registers)[op.register] = value
	(*registers)[PCRegister] += 2
}

func parse(lines []string) *list.List[Instruction] {
	noopRegex := regexp.MustCompile(`noop`)
	addRegex := regexp.MustCompile(`add(.)\s+(-?\d+)`)

	instructions := list.New[Instruction]()
	for _, line := range lines {
		if noopRegex.MatchString(line) {
			instructions.PushBack(Noop{})
		} else {
			addMatch := addRegex.FindAllStringSubmatch(line, -1)
			if addMatch != nil {
				register := addMatch[0][1]
				operand, err := strconv.Atoi(addMatch[0][2])
				aoccommon.CheckError(err)
				instructions.PushBack(Add{register: register, operand: operand})
			} else {
				panic("Unexpected line: " + line)
			}
		}
	}

	return instructions
}

func runPart1(instructions *list.List[Instruction]) int {
	registers := map[string]int{}
	registers["x"] = 1
	registers["pc"] = 1

	sumOfStrengths := 0
	nextCycleToGrab := 20

	for i := range iterateList(instructions) {
		xBefore := registers["x"]
		i.execute(&registers)
		pc := registers[PCRegister]

		if pc > nextCycleToGrab {
			signalStrength := xBefore * nextCycleToGrab
			sumOfStrengths += signalStrength
			nextCycleToGrab += 40
		}
	}

	return sumOfStrengths
}

const screenWidth = 40
const screenHeight = 6

func runPart2(instructions *list.List[Instruction]) []rune {
	registers := map[string]int{}
	registers["x"] = 1
	registers["pc"] = 1

	screenMemory := make([]rune, screenHeight*screenWidth)
	for i := 0; i < screenHeight*screenWidth; i++ {
		screenMemory[i] = '.'
	}
	drawnPc := 1

	for i := range iterateList(instructions) {
		xBefore := registers["x"]

		i.execute(&registers)
		pc := registers[PCRegister]

		for crtPosition := drawnPc; crtPosition < pc; crtPosition++ {
			// We're 1-based, so we need some shenanigans for the % to work as desired
			crtHorizPosition := 1 + (crtPosition-1)%screenWidth
			if crtHorizPosition >= xBefore && crtHorizPosition <= xBefore+2 {
				screenMemory[crtPosition-1] = '#'
			}
		}
		drawnPc = pc

	}

	return screenMemory
}

func toScreen(screenMemory []rune) string {
	result := ""

	for row := 0; row < screenHeight; row++ {
		rowStart := row * screenWidth
		rowEnd := rowStart + screenWidth
		result += fmt.Sprintf("Cycle %3d -> %s <- Cycle %3d\n", 1+rowStart, string(screenMemory[rowStart:rowEnd]), 1+rowEnd)
	}

	return result
}

func iterateList[T any](list *list.List[T]) <-chan T {
	ch := make(chan T)
	go func() {
		node := list.Front
		for node != nil {
			ch <- node.Value
			node = node.Next
		}
		close(ch)
	}()
	return ch
}

func computePart1(lines []string) int {
	instructions := parse(lines)
	return runPart1(instructions)
}

func computePart2(lines []string) string {
	instructions := parse(lines)
	screenMemory := runPart2(instructions)

	return toScreen(screenMemory)
}
