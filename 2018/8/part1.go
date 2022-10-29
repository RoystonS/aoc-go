package main

import (
	"aoccommon"
	"fmt"
)

func computePart1(lines []string) int {
	q := parseNumbers(lines)

	rootNode := readNode(q)
	metadataSum := 0

	var visitNode func(node Node)
	visitNode = func(node Node) {
		for _, childNode := range node.child_nodes {
			visitNode(childNode)
		}
		for _, metadata := range node.metadata {
			metadataSum += metadata
		}
	}
	visitNode(rootNode)

	return metadataSum
}

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)
	result := computePart1(lines)
	fmt.Println(result)
}
