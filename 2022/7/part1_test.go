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

	value := computePart1(lines)
	assert.Equal(uint64(95437), value)
}
