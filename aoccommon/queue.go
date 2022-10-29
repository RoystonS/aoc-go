package aoccommon

import (
	"container/list"
)

type Queue[T any] struct {
	items *list.List
}

func (l *Queue[T]) Init() *Queue[T] {
	l.items = list.New()
	return l
}

func (q *Queue[T]) IsEmpty() bool {
	return q.items.Len() == 0
}
func (q *Queue[T]) Queue(value T) {
	q.items.PushBack(value)
}
func (q *Queue[T]) Dequeue() (value T, hasValue bool) {
	if q.IsEmpty() {
		var noResult T
		return noResult, false
	} else {
		frontElem := q.items.Front()
		value := frontElem.Value
		q.items.Remove(frontElem)
		return any(value).(T), true
	}
}

func NewQueue[T any]() *Queue[T] { return new(Queue[T]).Init() }
