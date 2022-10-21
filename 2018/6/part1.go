package main

import (
	"fmt"
	"strconv"

	"aoccommon"
)

func findClosest(fromPoint Position, points []Point) *Point {

	minDistance := manhattanDistance(fromPoint, points[0].Position)
	var minDistancePoint Point
	minDistancePoint = points[0]
	countAtDistance := 1

	for _, point := range points {
		distance := manhattanDistance(fromPoint, point.Position)
		// fmt.Printf("Distance from %d to %d is %d\n", fromPoint, point, distance)
		switch {
		case distance == minDistance:
			if minDistancePoint != point {
				countAtDistance++
			}
		case distance < minDistance:
			minDistance = distance
			minDistancePoint = point
			countAtDistance = 1
		}
	}

	if countAtDistance == 1 {
		return &minDistancePoint
	} else {
		return nil
	}
}

func computeDistances(points []Point) (countOfClosestTo map[int]int, isInfinite map[int]bool) {
	bounds := computeBounds(points)

	countOfClosestTo = map[int]int{}
	isInfinite = map[int]bool{}

	// Compute the closest input point for every point on the surface one
	// unit around all the input points.  Any point on the edge with a single
	// closest input point is an infinite area.
	for x := bounds.minX - 1; x <= bounds.maxX+1; x++ {
		for y := bounds.minY - 1; y <= bounds.maxY+1; y++ {
			isOnEdge := x == bounds.minX-1 || x == bounds.maxX+1 || y == bounds.minY-1 || y == bounds.maxY+1

			pos := Position{x: x, y: y}
			closest := findClosest(pos, points)
			if closest != nil {
				id := (*closest).id
				countOfClosestTo[id]++

				if isOnEdge {
					isInfinite[id] = true
				}
			}
		}
	}

	return countOfClosestTo, isInfinite
}

func calculatePart1(points []Point) int {

	countOfClosestTo, isInfinite := computeDistances(points)

	largestFinite := 0
	for id, count := range countOfClosestTo {
		if !isInfinite[id] && count > largestFinite {
			largestFinite = count
		}
	}

	return largestFinite
}

func test1() {
	lines := []string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"}
	points := parsePoints(lines)
	largestFinite := calculatePart1(points)
	if largestFinite != 17 {
		panic("test fail: " + strconv.Itoa(largestFinite))
	}
}

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	points := parsePoints(lines)
	largestFinite := calculatePart1(points)
	fmt.Println(largestFinite)
}
