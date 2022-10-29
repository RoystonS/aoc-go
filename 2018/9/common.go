package main

import (
	"aoccommon"
	"container/list"
)

type Circle struct {
	items   *list.List
	current *list.Element
}

func (c *Circle) Insert(num int) {
	// fmt.Println("Insert", num)

	var el *list.Element
	if c.current == nil {
		c.items = list.New()
		el = c.items.PushFront(num)
	} else {
		oneClockwise := c.nextElement(c.current)
		el = c.items.InsertAfter(num, oneClockwise)
	}
	c.current = el
}

func (c *Circle) Remove7ToRight() int {
	// fmt.Printf("skipping back\n")
	sevenToRight := c.skipBack(c.current, 7)
	// fmt.Printf("7tor: %d\n", sevenToRight.Value)
	c.current = c.nextElement(sevenToRight)

	value := any(sevenToRight.Value).(int)
	c.items.Remove(sevenToRight)
	return value
}

func (c *Circle) skipBack(start *list.Element, count int) *list.Element {
	el := start

	for i := 0; i < count; i++ {
		// fmt.Printf("skipBack %d\n", i)
		el = c.prevElement(el)
	}

	return el
}

func (c *Circle) prevElement(e *list.Element) *list.Element {
	// fmt.Println("prevElement", e)
	prev := e.Prev()
	// fmt.Printf("prevElement %p\n", prev)
	if prev == nil {
		// We've run off the front. Go to the end
		return c.items.Back()
	}
	return prev
}

func (c *Circle) nextElement(e *list.Element) *list.Element {
	nextListElement := e.Next()
	if nextListElement == nil {
		// We've run off the end. Go to the front
		return c.items.Front()
	}
	return nextListElement
}

func runGame(players int, lastMarble int) int {
	scores := make([]int, players)

	circle := Circle{}
	for marble := 0; marble <= lastMarble; marble++ {
		player := marble % players

		if marble > 0 && marble%23 == 0 {
			removed := circle.Remove7ToRight()
			scores[player] += marble + removed
		} else {
			circle.Insert(marble)
		}
	}

	highestScore := 0
	for _, score := range scores {
		highestScore = aoccommon.Max(highestScore, score)
	}

	return highestScore
}
