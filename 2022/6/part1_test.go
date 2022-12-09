package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, computePart1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(6, computePart1("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(10, computePart1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(11, computePart1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
