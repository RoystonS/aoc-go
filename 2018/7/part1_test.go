package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata")
	aoccommon.CheckError(err)

	g := parseLines(lines, 6)
	value := tsortToString(g)

	assert.Equal("CABDFE", value)
}
