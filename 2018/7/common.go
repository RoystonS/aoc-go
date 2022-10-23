package main

import (
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

func tsortToString(g *graph.Mutable) string {
	var sb strings.Builder
	vertices := tsort(g)
	for _, vertex := range vertices {
		sb.WriteRune(rune(vertex + 65))
	}

	return sb.String()
}
