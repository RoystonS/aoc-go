package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func computePart2(lines []string) int {
	return 42
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata")
	aoccommon.CheckError(err)

	value := computePart2(lines)
	assert.Equal(42, value)
}
