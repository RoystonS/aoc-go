package aoccommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackEmpty(t *testing.T) {
	assert := assert.New(t)

	q := NewStack[int]()
	assert.Equal(true, q.IsEmpty())
	q.Push(42)
	assert.Equal(false, q.IsEmpty())
}

func TestStack(t *testing.T) {
	assert := assert.New(t)

	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	value, hasValue := s.Pop()
	assert.Equal(hasValue, true)
	assert.Equal(3, value)

	value, hasValue = s.Pop()
	assert.Equal(hasValue, true)
	assert.Equal(2, value)

	value, hasValue = s.Pop()
	assert.Equal(hasValue, true)
	assert.Equal(1, value)

	_, hasValue = s.Pop()
	assert.Equal(hasValue, false)
}
