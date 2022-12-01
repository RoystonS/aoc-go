package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata")
	aoccommon.CheckError(err)

	result := computePart2(lines)
	assert.Equal(45000, result)
}
