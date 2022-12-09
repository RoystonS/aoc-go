package main

import (
	"github.com/gammazero/deque"
)

func computePart1(line string) int {
	return compute(line, 4)
}

func computePart2(line string) int {
	return compute(line, 14)
}

func compute(line string, distinctChars int) int {
	// We keep the last 'n' characters in a deque, and check to
	// see if they're different
	d := deque.New[rune]()

	for i, ch := range line {
		if d.Len() == distinctChars {
			// Our deque already contains the correct number of characters, so pop off
			// the first one so it's ready to receive the next
			d.PopFront()
		}
		d.PushBack(ch)

		// If the deque is populated, check to see if its characters are all different
		if d.Len() == distinctChars {
			q := map[rune]bool{}

			for j := 0; j < distinctChars; j++ {
				q[d.At(j)] = true
			}

			if len(q) == distinctChars {
				// Yes. All the characters are different.
				return i + 1
			}
		}
	}

	panic("Not found")
}
