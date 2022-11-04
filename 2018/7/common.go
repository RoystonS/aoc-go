package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/yourbasic/graph"
)

func parseLines(lines []string, vertexCount int) *graph.Mutable {
	g := graph.New(vertexCount)
	re := regexp.MustCompile(`Step (.) must be finished before step (.) can begin`)

	for _, line := range lines {
		res := re.FindAllStringSubmatch(line, -1)
		step1, _ := utf8.DecodeRuneInString(res[0][1])
		step2, _ := utf8.DecodeRuneInString(res[0][2])

		g.Add(int(step1-65), int(step2-65))
	}

	return g
}

func tsort(g *graph.Mutable) []int {
	result := make([]int, g.Order())
	indegree := make([]int, g.Order())

	for v := range indegree {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			indegree[w]++
			return
		})
	}

	for i := 0; i < g.Order(); i++ {
		for v, degree := range indegree {
			if degree == 0 {
				result[i] = v
				indegree[v] = -1

				g.Visit(v, func(w int, _ int64) (skip bool) {
					indegree[w]--
					return
				})
				break
			}
		}
	}

	return result
}

type Worker struct {
	id                int
	workingOnTask     int
	availableAtSecond int
}

func timedTsort(g *graph.Mutable, workerCount int) []int {
	result := make([]int, g.Order())
	indegree := make([]int, g.Order())
	workers := make([]Worker, workerCount)

	// Initialize workers
	for index := range workers {
		workers[index].id = index + 1
		workers[index].workingOnTask = -1
	}

	// Calculate indegrees of nodes in the graph
	for v := range indegree {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			indegree[w]++
			return
		})
	}

	// currentTime := 0

	for i := 0; i < 5; i++ {
		// Find next unblocked task, in alphabetical order
		nextTask := -1
		for v, degree := range indegree {
			if degree == 0 {
				nextTask = v
				indegree[v] = -1
				break
			}
		}

		fmt.Println()
		fmt.Printf("Next unblocked task: %d/%c\n", nextTask, nextTask+65)

		// When is the next worker available?
		var nextAvailableWorker = &Worker{
			id:                -1,
			workingOnTask:     -1,
			availableAtSecond: math.MaxInt,
		}

		// Find the next available worker
		for index, worker := range workers {
			if worker.availableAtSecond < nextAvailableWorker.availableAtSecond {
				fmt.Printf(" worker interesting: %d %d\n", index, worker.availableAtSecond)
				if nextTask >= 0 {
					// Any available worker is good
					fmt.Println("  good1!")
					nextAvailableWorker = &workers[index]
				} else {
					// There are no tasks available so we're only interested in working workers
					if worker.workingOnTask >= 0 {
						fmt.Println("  good2!")
						nextAvailableWorker = &workers[index]
					}
				}
			}
		}

		fmt.Printf("Next available worker %d (was on %c) at time %d\n", nextAvailableWorker.id, 65+nextAvailableWorker.workingOnTask, nextAvailableWorker.availableAtSecond)

		if nextAvailableWorker.workingOnTask >= 0 {
			// Their task is complete
			fmt.Printf(" This means that task %c is now complete\n", nextAvailableWorker.workingOnTask+65)
			g.Visit(nextAvailableWorker.workingOnTask, func(w int, _ int64) (skip bool) {
				fmt.Printf("  Reducing indegree of task %c from %d to %d\n", w+65, indegree[w], indegree[w]-1)
				indegree[w]--
				return
			})
		}

		if nextTask > 0 {
			durationOfTask := 60 + nextTask + 1
			fmt.Printf(" Worker %d being assigned to task %c with duration %d (end: %d)\n", nextAvailableWorker.id, nextTask+65, durationOfTask, nextAvailableWorker.availableAtSecond+durationOfTask)

			nextAvailableWorker.availableAtSecond = nextAvailableWorker.availableAtSecond + durationOfTask
			nextAvailableWorker.workingOnTask = nextTask

			// currentTime = nextAvailableWorker.availableAtSecond
		} else {
			// We have spare workers but nothing for them to do!
			fmt.Println("idle workers!")
		}

		fmt.Println(workers)
	}
	return result
}

func tsortToString(g *graph.Mutable) string {
	var sb strings.Builder
	vertices := tsort(g)
	for _, vertex := range vertices {
		sb.WriteRune(rune(vertex + 65))
	}

	return sb.String()
}
