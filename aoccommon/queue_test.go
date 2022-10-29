package aoccommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEmpty(t *testing.T) {
	assert := assert.New(t)

	q := NewQueue[int]()
	assert.Equal(true, q.IsEmpty())
	q.Queue(42)
	assert.Equal(false, q.IsEmpty())
}

func TestQueue(t *testing.T) {
	assert := assert.New(t)

	q := NewQueue[int]()
	q.Queue(1)
	q.Queue(2)
	q.Queue(3)

	value, hasValue := q.Dequeue()
	assert.Equal(hasValue, true)
	assert.Equal(1, value)

	value, hasValue = q.Dequeue()
	assert.Equal(hasValue, true)
	assert.Equal(2, value)

	value, hasValue = q.Dequeue()
	assert.Equal(hasValue, true)
	assert.Equal(3, value)

	_, hasValue = q.Dequeue()
	assert.Equal(hasValue, false)
}
