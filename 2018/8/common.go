package main

import (
	"aoccommon"
	"strconv"
	"strings"
)

func parseNumbers(lines []string) *aoccommon.Queue[int] {
	strs := strings.Split(lines[0], " ")

	q := aoccommon.NewQueue[int]()

	for _, str := range strs {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		q.Queue(v)
	}

	return q
}

type Node struct {
	child_node_count int
	metadata_count   int
	child_nodes      []Node
	metadata         []int
}

func readNodeHeader(queue *aoccommon.Queue[int]) Node {
	child_node_count, hasValue := queue.Dequeue()
	if !hasValue {
		panic("Queue is empty")
	}
	var metadata_count int
	metadata_count, hasValue = queue.Dequeue()
	if !hasValue {
		panic("Queue is empty")
	}

	node := Node{
		child_node_count: child_node_count,
		metadata_count:   metadata_count,
	}
	return node
}

func readNode(queue *aoccommon.Queue[int]) Node {
	node := readNodeHeader(queue)

	children := make([]Node, node.child_node_count)
	node.child_nodes = children

	for count := 0; count < node.child_node_count; count++ {
		node.child_nodes[count] = readNode(queue)
	}

	metadata := make([]int, node.metadata_count)
	node.metadata = metadata
	for count := 0; count < node.metadata_count; count++ {
		metadataValue, hasValue := queue.Dequeue()
		if !hasValue {
			panic("end of queue")
		}
		node.metadata[count] = metadataValue
	}

	return node
}
