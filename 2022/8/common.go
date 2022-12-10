package main

import "fmt"

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

type Forest struct {
	// The height of each tree
	heights [][]int8

	// The maximum height seen FROM each position in each direction
	maxHeight map[string]int8
}

func NewForest() *Forest {
	return &Forest{heights: nil, maxHeight: nil}
}

func (forest *Forest) rows() int {
	return len(forest.heights)
}

func (forest *Forest) cols() int {
	return len(forest.heights[0])
}

func (forest *Forest) ReadFrom(lines []string) {
	cols := len(lines[0])
	rows := len(lines)

	heights := make([][]int8, rows)
	inner := make([]int8, cols*rows)
	for i := range heights {
		heights[i] = inner[i*cols : (i+1)*cols]
	}
	forest.heights = heights

	forest.maxHeight = map[string]int8{}

	for y, line := range lines {
		for x, ch := range line {
			heights[y][x] = int8(ch - '0')
		}
	}
}

func (forest *Forest) isVisible(x int, y int) bool {
	height := forest.heights[y][x]
	return height > forest.maxHeightFrom(x, y, North) ||
		height > forest.maxHeightFrom(x, y, East) ||
		height > forest.maxHeightFrom(x, y, South) ||
		height > forest.maxHeightFrom(x, y, West)
}

func (forest *Forest) ScenicScore(x int, y int) int {
	return forest.viewingDistanceFrom(x, y, North) * forest.viewingDistanceFrom(x, y, South) * forest.viewingDistanceFrom(x, y, East) * forest.viewingDistanceFrom(x, y, West)
}

func (forest *Forest) isOnEdge(x int, y int, d Direction) bool {
	return (x == 0 && d == West) || (y == 0 && d == North) ||
		(x == forest.cols()-1 && d == East) || (y == forest.rows()-1 && d == South)
}

func getNeighbourPosition(x int, y int, d Direction) (int, int) {
	neighbourX := x
	neighbourY := y
	switch d {
	case North:
		neighbourY--
	case South:
		neighbourY++
	case East:
		neighbourX++
	case West:
		neighbourX--
	}

	return neighbourX, neighbourY
}

func (forest *Forest) maxHeightFrom(x int, y int, d Direction) int8 {
	key := visibilityKey(x, y, d)

	if forest.maxHeight[key] == 0 {
		// We've not already been calculated
		if forest.isOnEdge(x, y, d) {
			// We're on the appropriate edge, so we're visible by definition
			forest.maxHeight[key] = -1
		}
	}

	if forest.maxHeight[key] == 0 {
		// We're not on an edge
		neighbourX, neighbourY := getNeighbourPosition(x, y, d)

		maxHeightInDirection := forest.maxHeightFrom(neighbourX, neighbourY, d)
		// fmt.Printf("maxHeightInDirection from %d, %d in %d is %d\n", neighbourX, neighbourY, d, maxHeightInDirection)
		if forest.heights[neighbourY][neighbourX] > maxHeightInDirection {
			// fmt.Printf("Neighbour is highest, at %d\n", forest.heights[neighbourY][neighbourX])
			forest.maxHeight[key] = forest.heights[neighbourY][neighbourX]
		} else {
			// fmt.Printf("Neighbour is not highest. Keeping max of %d\n", maxHeightInDirection)
			forest.maxHeight[key] = maxHeightInDirection
		}
	}

	return forest.maxHeight[key]
}

func (forest *Forest) viewingDistanceFrom(x int, y int, d Direction) int {
	height := forest.heights[y][x]
	distance := 0

	// Run in the specified direction until we hit something our height or bigger, or the edge
	for !forest.isOnEdge(x, y, d) {
		distance++
		x, y = getNeighbourPosition(x, y, d)
		h := forest.heights[y][x]
		if h >= height {
			return distance
		}
	}
	return distance
}

func (forest *Forest) VisibilityCount() int {
	count := 0
	for y, arr := range forest.heights {
		for x := range arr {
			if forest.isVisible(x, y) {
				count++
			}
		}
	}

	return count
}

func (forest *Forest) MaximumScenicScore() int {
	max := 0
	for y, arr := range forest.heights {
		for x := range arr {
			dist := forest.ScenicScore(x, y)
			if dist > max {
				max = dist
			}
		}
	}

	return max

}
func (forest *Forest) dumpVisibilities() {
	for y, arr := range forest.heights {
		for x, height := range arr {
			if forest.isVisible(x, y) {
				fmt.Printf("(%d)", height)
			} else {
				fmt.Printf(" %d ", height)
			}
		}
		fmt.Println()
	}
}

func visibilityKey(x int, y int, d Direction) string {
	return fmt.Sprintf("%d!%d!%d", x, y, d)
}

// Forest is represented as a 2D array of ints
// Need to track visibility of each tree,
// whether it's visible from N, S, E, W.
// Could mess about with spiral patterns, filling in,
// or just memoize. e.g. for any tree T, it's visible in
// direction X if it's on the X-most border OR its
// next X-most neighbour K is visible in direction X and
// T is bigger than K

func computePart1(lines []string) int {
	forest := NewForest()
	forest.ReadFrom(lines)

	// forest.dumpVisibilities()
	return forest.VisibilityCount()
}

func computePart2(lines []string) int {
	forest := NewForest()
	forest.ReadFrom(lines)

	return forest.MaximumScenicScore()
}
