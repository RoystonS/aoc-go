package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(19, computePart2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(23, computePart2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(23, computePart2("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(29, computePart2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(26, computePart2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))

}
