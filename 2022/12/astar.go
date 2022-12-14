package main

import (
	"fmt"
	"sort"

	"github.com/zyedidia/generic/list"
)

type Position struct {
	row int
	col int
}

type Direction struct {
	rowDelta int
	colDelta int
}

var directions []Direction = []Direction{
	{-1, 0},
	{+1, 0},
	{0, -1},
	{0, +1},
}

func (pos *Position) id() string {
	return fmt.Sprintf("%d!%d", pos.row, pos.col)
}

func (pos *Position) Add(delta Direction) Position {
	return Position{
		row: pos.row + delta.rowDelta,
		col: pos.col + delta.colDelta,
	}
}

func (pos *Position) Equals(otherPosition *Position) bool {
	return pos.row == otherPosition.row && pos.col == otherPosition.col
}

type AStar struct {
	problem AStarProblem

	goalPos    *Position
	openSet    []*Position
	openSetIds map[string]bool
	closedSet  map[string]bool
	cameFrom   map[string]*Position
	fScore     map[string]int
	gScore     map[string]int
}

type AStarProblem interface {
	// Gets the positions that can be moved to from a specified position
	GetNeighbours(from *Position) *list.List[*Position]

	// Gets an estimate of the cost from one arbitrary position to another
	GetHeuristicCost(from *Position, to *Position) int

	// Gets the actual cost from one position to a neighbour
	GetActualCost(from *Position, to *Position) int
}

func NewAStar(problem AStarProblem, startPos *Position, goalPos *Position) *AStar {
	openSet := []*Position{startPos}
	openSetIds := map[string]bool{startPos.id(): true}

	return &AStar{
		problem:    problem,
		goalPos:    goalPos,
		openSet:    openSet,
		openSetIds: openSetIds,
		closedSet:  map[string]bool{},
		cameFrom:   map[string]*Position{},
		fScore:     map[string]int{},
		gScore:     map[string]int{},
	}
}

func (algo *AStar) Step() bool {
	if len(algo.openSet) == 0 {
		return false
	}

	// Pop the currentPos off the end of the slice
	end := len(algo.openSet) - 1
	currentPos := algo.openSet[end]
	algo.openSet = algo.openSet[0:end]
	delete(algo.openSetIds, currentPos.id())

	if currentPos.Equals(algo.goalPos) {
		return false
	}

	// algo.DumpPathTo(&currentPos)
	algo.closedSet[currentPos.id()] = true

	neighbours := algo.problem.GetNeighbours(currentPos).Front

	for neighbours != nil {
		neighbour := neighbours.Value
		neighbours = neighbours.Next

		nid := neighbour.id()
		if algo.closedSet[nid] {
			continue
		}

		gScoreForThisRoute := algo.gScore[currentPos.id()] + algo.problem.GetActualCost(currentPos, neighbour)

		if algo.openSetIds[nid] {
			if gScoreForThisRoute >= algo.gScore[nid] {
				// Worse path
				continue
			}
		} else {
			algo.openSet = append(algo.openSet, neighbour)
			algo.openSetIds[nid] = true
		}

		algo.cameFrom[nid] = currentPos
		algo.gScore[nid] = gScoreForThisRoute
		algo.fScore[nid] = gScoreForThisRoute + algo.problem.GetHeuristicCost(neighbour, algo.goalPos)
	}

	openSet := algo.openSet
	fScore := algo.fScore

	sort.Slice(openSet, func(i, j int) bool {
		pos1 := openSet[i]
		pos2 := openSet[j]
		return fScore[pos1.id()] > fScore[pos2.id()]
	})

	return true
}

func (algo *AStar) GetPreviousPositionInPath(pos *Position) *Position {
	return algo.cameFrom[pos.id()]
}

func (algo *AStar) PathLength() int {
	return algo.gScore[algo.goalPos.id()]
}
