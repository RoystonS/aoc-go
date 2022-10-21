package main

import (
	"fmt"
	"strconv"

	"aoccommon"
)

func calculatePart2(lines []string, maxDistance int) int {
	points := parsePoints(lines)
	bounds := computeBounds(points)

	pointsWithinDistance := 0

	// Start in the centre, and work round in a spiral.
	// Once we do a loop without interesting points, we're done,
	// as there's only one region.

	var currentX = (bounds.maxX + bounds.minX) >> 1
	var currentY = (bounds.maxY + bounds.minY) >> 1

	// Current direction
	var dx = 1
	var dy = 0

	// Bounds of the spiral
	var minX = currentX
	var maxX = currentX
	var minY = currentY
	var maxY = currentY

	foundItemWithinLoop := false

	for {
		current := Position{
			x: currentX,
			y: currentY,
		}

		totalDistance := 0
		for _, point := range points {
			totalDistance += manhattanDistance(current, point.Position)
		}

		// fmt.Printf("x: %d, y: %d; totDist: %d; dx: %d, dy: %d\n", currentX, currentY, totalDistance, dx, dy)

		if totalDistance <= maxDistance {
			foundItemWithinLoop = true
			pointsWithinDistance++
		}

		currentX = currentX + dx
		currentY = currentY + dy

		switch {
		case currentX > maxX:
			maxX++
			dx, dy = 0, -1
		case currentX < minX:
			minX--
			dx, dy = 0, 1
		case currentY > maxY:
			maxY++
			dx, dy = 1, 0
			if !foundItemWithinLoop {
				// We've completed a spiral loop without finding anything within distance
				// fmt.Printf("Bounds: %d,%d to %d, %d\n", minX, minY, maxX, maxY)
				return pointsWithinDistance
			}
			foundItemWithinLoop = false
		case currentY < minY:
			minY--
			dx, dy = -1, 0
		}
	}
}

func test2() {
	lines := []string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"}
	result := calculatePart2(lines, 30)
	if result != 16 {
		panic("test2 fail: " + strconv.Itoa(result))
	}
}

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := calculatePart2(lines, 10000)
	fmt.Println(result)
}
