package main

import (
	"aoccommon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	lines, err := aoccommon.ReadLines("testdata2")
	aoccommon.CheckError(err)

	value := computePart2(lines)
	assert.Equal(36, value)
}
