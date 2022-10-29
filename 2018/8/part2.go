package main

import (
	"aoccommon"
	"fmt"
)

func computeNodeValue(node *Node) int {
	result := 0

	for _, metadata := range node.metadata {
		if node.child_node_count == 0 {
			// Node has no children. The 'value' is the sum of the metadata
			result = result + metadata
		} else {
			// Node has children. The 'value' is the sum of the child values
			// as referenced by the metadata
			if metadata >= 1 && metadata <= node.child_node_count {
				index := metadata - 1
				result = result + computeNodeValue(&node.child_nodes[index])
			}
		}
	}

	return result
}

func computePart2(lines []string) int {
	q := parseNumbers(lines)

	rootNode := readNode(q)
	value := computeNodeValue(&rootNode)
	return value
}

func part2() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	result := computePart2(lines)
	fmt.Println(result)
}
