package aoccommon

import "github.com/zyedidia/generic/list"

func IterateList[T any](list *list.List[T]) <-chan T {
	ch := make(chan T)
	go func() {
		node := list.Front
		for node != nil {
			ch <- node.Value
			node = node.Next
		}
		close(ch)
	}()
	return ch
}

func CountList[T any](list *list.List[T]) int {
	i := 0
	for range IterateList(list) {
		i++
	}
	return i
}

// Go doesn't have EITHER overloading or default parameter values!?
func ToArray[T any](l *list.List[T], count int) *[]T {
	if count == 0 {
		count = CountList(l)
	}

	array := make([]T, count)

	i := 0
	for item := range IterateList(l) {
		array[i] = item
		i++
	}

	return &array
}
