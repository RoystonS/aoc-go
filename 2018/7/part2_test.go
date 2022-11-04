package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourbasic/graph"
)

func computePart2(g *graph.Mutable, workers int) int {
	timedTsort(g, workers)

	return 42
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata")
	aoccommon.CheckError(err)

	g := parseLines(lines, 6)
	value := computePart2(g, 2)
	assert.Equal(value, 42)
}
