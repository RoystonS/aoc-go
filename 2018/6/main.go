package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		part1()
	case "2":
		part2()
	case "test1":
		test1()
	case "test2":
		test2()
	}
}
