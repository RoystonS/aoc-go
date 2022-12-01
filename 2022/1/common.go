package main

import (
	"aoccommon"
	"sort"
	"strconv"

	"github.com/zyedidia/generic/list"
)

type Elf struct {
	lines *list.List[string]
}

func NewElf() *Elf {
	return &Elf{
		lines: list.New[string](),
	}
}

func (elf Elf) total() int {
	tot := 0

	node := elf.lines.Front
	for node != nil {
		cals, err := strconv.Atoi(node.Value)
		aoccommon.CheckError(err)
		tot += cals
		node = node.Next
	}

	return tot
}

func listLen[T any](list *list.List[T]) int {
	count := 0

	node := list.Front
	for node != nil {
		count++
		node = node.Next
	}

	return count
}

func toArray[V any](list *list.List[V]) []V {
	length := listLen(list)

	arr := make([]V, length)

	firstNode := list.Front
	i := 0

	firstNode.Each(func(val V) {
		arr[i] = val
		i++
	})

	return arr
}

func parse(lines []string) *list.List[*Elf] {
	elves := list.New[*Elf]()

	elf := NewElf()

	for _, line := range lines {
		if len(line) == 0 {
			elves.PushBack(elf)
			elf = NewElf()
		} else {
			elf.lines.PushBack(line)
		}
	}

	if elf.lines.Front != nil {
		elves.PushBack(elf)
	}

	return elves
}

func computePart1(lines []string) int {
	elves := sortedElves(lines)

	return elves[0].total()
}

func sortedElves(lines []string) []*Elf {
	elves := parse(lines)
	elvesArray := toArray(elves)

	sort.SliceStable(elvesArray, func(i int, j int) bool {
		e1 := (elvesArray)[i]
		e2 := (elvesArray)[j]
		return e1.total() > e2.total()
	})

	return elvesArray
}

func computePart2(lines []string) int {
	elves := sortedElves(lines)

	return elves[0].total() + elves[1].total() + elves[2].total()
}
