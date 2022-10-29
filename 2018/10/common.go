package main

import (
	"aoccommon"
	"fmt"
	"regexp"
	"strconv"
)

type Point struct {
	x  int
	y  int
	vx int
	vy int
}

func parseLines(lines []string) []Point {
	result := make([]Point, len(lines))

	// position=< 9,  1> velocity=< 0,  2>
	re := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
	for i, line := range lines {
		match := re.FindAllStringSubmatch(line, -1)
		x, _ := strconv.Atoi(match[0][1])
		y, _ := strconv.Atoi(match[0][2])
		vx, _ := strconv.Atoi(match[0][3])
		vy, _ := strconv.Atoi(match[0][4])

		result[i].x = x
		result[i].y = y
		result[i].vx = vx
		result[i].vy = vy
	}

	return result
}

func runMovement(lines []string, output bool) {
	points := parseLines(lines)

	var oldWidth, oldHeight int

	for round := 0; ; round++ {

		minX := points[0].x
		minY := points[0].y
		maxX := points[0].x
		maxY := points[0].y

		for _, point := range points {
			minX = aoccommon.Min(minX, point.x)
			minY = aoccommon.Min(minY, point.y)
			maxX = aoccommon.Max(maxX, point.x)
			maxY = aoccommon.Max(maxY, point.y)
		}
		newWidth := maxX - minX
		newHeight := maxY - minY

		// fmt.Printf("Range: %d,%d -> %d,%d [%d,%d] from [%d,%d]\n", minX, minY, maxX, maxY, newWidth, newHeight, oldWidth, oldHeight)
		if round > 1 && (newWidth > oldWidth || newHeight > oldHeight) {
			applyDeltas(points, -1)
			filledPoints := computeFilledPoints(points)

			if output {
				// if round > 1 && newWidth < 100 {
				for row := minY; row <= maxY; row++ {
					for col := minX; col <= maxX; col++ {
						if (*filledPoints)[key(col, row)] {
							fmt.Printf("#")
						} else {
							fmt.Printf(".")
						}
					}
					fmt.Println()
				}
				fmt.Println()
			} else {
				fmt.Println(round - 1)
			}
			return
		}

		applyDeltas(points, 1)

		oldHeight = newHeight
		oldWidth = newWidth
	}
}

func applyDeltas(points []Point, multiplier int) {
	for i := range points {
		points[i].x += points[i].vx * multiplier
		points[i].y += points[i].vy * multiplier
	}
}

func computeFilledPoints(points []Point) *map[string]bool {
	filledPoints := map[string]bool{}
	for _, point := range points {
		filledPoints[key(point.x, point.y)] = true
	}
	return &filledPoints
}

func key(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
