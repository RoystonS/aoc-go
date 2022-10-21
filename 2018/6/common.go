package main

import (
	"aoccommon"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Point struct {
	id int
	Position
}

type Bounds struct {
	minX int
	minY int
	maxX int
	maxY int
}

func parsePoints(lines []string) []Point {
	points := []Point{}

	nextId := 1
	for _, line := range lines {
		coords := strings.Split(line, ", ")
		x, err := strconv.Atoi(coords[0])
		aoccommon.CheckError(err)
		y, err := strconv.Atoi(coords[1])
		aoccommon.CheckError(err)

		points = append(points, Point{
			id: nextId,
			Position: Position{
				x: x,
				y: y,
			},
		})
		nextId++
	}

	return points
}

func computeBounds(points []Point) Bounds {
	minX := points[0].x
	maxX := points[0].x
	minY := points[0].y
	maxY := points[0].y

	for _, point := range points {
		minX = aoccommon.Min(minX, point.x)
		maxX = aoccommon.Max(maxX, point.x)
		minY = aoccommon.Min(minY, point.y)
		maxY = aoccommon.Max(maxY, point.y)
	}

	return Bounds{minX, minY, maxX, maxY}
}

func manhattanDistance(p1, p2 Position) (distance int) {
	return aoccommon.Abs(p1.x-p2.x) + aoccommon.Abs(p1.y-p2.y)
}
