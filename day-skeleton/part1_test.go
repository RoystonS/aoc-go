package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func computePart1(lines []string) int {
	return 42
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata")
	aoccommon.CheckError(err)

	value := computePart1(lines)
	assert.Equal(42, value)
}
