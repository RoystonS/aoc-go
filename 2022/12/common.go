package main

import (
	"aoccommon"
	"strings"
	"sync"

	"github.com/zyedidia/generic/list"
)

func find(lines []string, ch rune) Position {
	for row, line := range lines {
		col := strings.IndexRune(line, ch)
		if col >= 0 {
			return Position{row: row, col: col}
		}
	}
	panic("Position not found")
}

func canTraverseFromHeightToHeight(from byte, to byte) bool {
	if to == 'E' {
		to = 'z'
	}
	if from == 'S' {
		from = 'a'
	}

	if to == from+1 {
		return true
	}

	return to <= from
}

type Day13Problem struct {
	width  int
	height int

	lines []string
}

func (problem Day13Problem) GetHeuristicCost(from *Position, to *Position) int {
	// Manhattan distance is the best possible cost from one position to another
	return aoccommon.Abs(to.row-from.row) + aoccommon.Abs(to.col-from.col)
}
func (problem Day13Problem) GetActualCost(from *Position, to *Position) int {
	return 1
}

func (problem Day13Problem) GetNeighbours(pos *Position) *list.List[*Position] {
	result := list.New[*Position]()

	currentHeight := problem.lines[pos.row][pos.col]

	for _, dir := range directions {
		neighbour := pos.Add(dir)
		if neighbour.row < 0 || neighbour.row >= problem.height || neighbour.col < 0 || neighbour.col >= problem.width {
			continue
		}
		neighbourHeight := problem.lines[neighbour.row][neighbour.col]
		if !canTraverseFromHeightToHeight(currentHeight, neighbourHeight) {
			continue
		}
		result.PushBack(&neighbour)
	}
	return result
}

func computePart1(lines []string) int {
	startPos := find(lines, 'S')
	endPos := find(lines, 'E')

	problem := Day13Problem{
		lines:  lines,
		height: len(lines),
		width:  len(lines[0]),
	}

	algo := NewAStar(problem, &startPos, &endPos)

	for algo.Step() {
	}

	return algo.PathLength()
}

func computePart2(lines []string) int {
	endPos := find(lines, 'E')

	problem := Day13Problem{
		lines:  lines,
		height: len(lines),
		width:  len(lines[0]),
	}

	algos := make([]*AStar, 0, 100)

	var wg sync.WaitGroup

	for row, line := range lines {
		for col, ch := range line {
			if ch == 'a' {
				algo := NewAStar(problem, &Position{row: row, col: col}, &endPos)
				algos = append(algos, algo)

				wg.Add(1)

				go func() {
					defer wg.Done()
					for algo.Step() {
					}
				}()
			}
		}
	}

	wg.Wait()

	shortestLength := algos[0].PathLength()
	for _, algo := range algos {
		algoLength := algo.PathLength()
		if algoLength > 0 {
			shortestLength = aoccommon.Min(shortestLength, algoLength)
		}
	}
	return shortestLength
}
