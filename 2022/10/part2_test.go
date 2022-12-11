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

	value := computePart2(lines)
	assert.Equal(`Cycle   1 -> ##..##..##..##..##..##..##..##..##..##.. <- Cycle  41
Cycle  41 -> ###...###...###...###...###...###...###. <- Cycle  81
Cycle  81 -> ####....####....####....####....####.... <- Cycle 121
Cycle 121 -> #####.....#####.....#####.....#####..... <- Cycle 161
Cycle 161 -> ######......######......######......#### <- Cycle 201
Cycle 201 -> #######.......#######.......#######..... <- Cycle 241
`, value)
}
