package aoccommon

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (value T, hasValue bool) {
	if s.IsEmpty() {
		var noResult T
		return noResult, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func NewStack[T any]() *Stack[T] { return new(Stack[T]) }
