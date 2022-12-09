package main

import (
	"aoccommon"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Computes a move from a knot position in some direction
func computeMove(x int, y int, direction string) (int, int) {
	switch direction {
	case "U":
		return x, y - 1
	case "D":
		return x, y + 1
	case "L":
		return x - 1, y
	}
	return x + 1, y
}

// Given the positions of a tail knot and its immediate head knot, computes
// the new position for the tail knot
func computeKnotMovement(headKnotX int, headKnotY int, tailKnotX int, tailKnotY int) (int, int) {
	xDelta := headKnotX - tailKnotX
	yDelta := headKnotY - tailKnotY

	roundedHalfXDelta := int(math.Round(float64(xDelta) / 2))
	roundedHalfYDelta := int(math.Round(float64(yDelta) / 2))

	if math.Abs(float64(xDelta)) > 1 || math.Abs(float64(yDelta)) > 1 {
		// We need to move because our position is at least 2 behind
		return tailKnotX + roundedHalfXDelta, tailKnotY + roundedHalfYDelta
	}

	return tailKnotX, tailKnotY
}

type Knot struct {
	x int
	y int
}

func computePart1(lines []string) int {
	// For part 1, we just have two knots
	return compute(lines, 2)
}

func computePart2(lines []string) int {
	// For part 2, we have ten knots
	return compute(lines, 10)
}

func compute(lines []string, knotCount int) int {
	knots := make([]Knot, knotCount)
	for i := 0; i < knotCount; i++ {
		knots[i] = Knot{x: 0, y: 0}
	}

	visitedTailPositions := map[string]bool{}

	visitedTailPositions[tailKey(knots[0].x, knots[0].y)] = true

	for _, line := range lines {
		bits := strings.Split(line, " ")
		direction := bits[0]
		distance, err := strconv.Atoi(bits[1])
		aoccommon.CheckError(err)

		for distance > 0 {
			newHeadX, newHeadY := computeMove(knots[0].x, knots[0].y, direction)
			// fmt.Printf("Head: %d, %d --%s %d-> %d, %d\n", knots[0].x, knots[0].y, direction, distance, newHeadX, newHeadY)
			knots[0].x = newHeadX
			knots[0].y = newHeadY

			for i := 1; i < knotCount; i++ {
				newX, newY := computeKnotMovement(knots[i-1].x, knots[i-1].y, knots[i].x, knots[i].y)

				knots[i].x = newX
				knots[i].y = newY
			}
			distance--
			visitedTailPositions[tailKey(knots[knotCount-1].x, knots[knotCount-1].y)] = true

			// dump(knots)
		}
	}

	return len(visitedTailPositions)
}

func tailKey(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

// Dumps out a visual display of the knots
func dump(knots []Knot) {
	for y := -10; y < 5; y++ {
		for x := -10; x < 10; x++ {
			ch := '.'

			if x == 0 && y == 0 {
				ch = 's'
			}

			for i, knot := range knots {
				if knot.x == x && knot.y == y {
					if i == 0 {
						ch = 'H'
						break
					} else {
						ch = rune('0' + i)
						break
					}
				}
			}
			fmt.Printf("%c", ch)
		}

		fmt.Printf(" %d\n", y)
	}
}
