package aoccommon

import (
	"bufio"
	"os"
	"strconv"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadNumbers(path string) ([]int, error) {
	lines, err := ReadLines(path)
	CheckError(err)

	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		CheckError(err)
		nums = append(nums, num)
	}
	return nums, err
}
