package main

import (
	"strings"
)

type Play int8

const (
	Rock Play = iota
	Paper
	Scissors
)

type Result int8

const (
	Win Result = iota
	Lose
	Draw
)

func (play Play) play(otherPlay Play) Result {
	if play == otherPlay {
		return Draw
	}

	switch play {
	case Rock:
		if otherPlay == Scissors {
			return Win
		}
	case Paper:
		if otherPlay == Rock {
			return Win
		}
	case Scissors:
		if otherPlay == Paper {
			return Win
		}
	}
	return Lose
}

var plays = []Play{Rock, Paper, Scissors}

func parseStrategy1(value string) Play {
	basicIndex := strings.Index("ABC", value)
	if basicIndex >= 0 {
		return plays[basicIndex]
	}

	ourStrategy := strings.Index("XYZ", value)
	return plays[ourStrategy]
}

func parseStrategy2(value string, opponentPlay Play) Play {
	var winningPlay Play
	var losingPlay Play

	switch opponentPlay {
	case Rock:
		winningPlay = Paper
		losingPlay = Scissors
	case Paper:
		winningPlay = Scissors
		losingPlay = Rock
	case Scissors:
		winningPlay = Rock
		losingPlay = Paper
	}

	switch value {
	case "X":
		// need to lose
		return losingPlay
	case "Z":
		// need to win
		return winningPlay
	}

	// need to draw
	return opponentPlay
}

func getPlayResultScore(ourPlay Play, opponentPlay Play) int {
	result := ourPlay.play(opponentPlay)
	switch result {
	case Win:
		return 6
	case Draw:
		return 3
	}

	return 0
}

func getPlayScore(ourPlay Play) int {
	switch ourPlay {
	case Rock:
		return 1
	case Paper:
		return 2
	}
	return 3
}

func computePart1(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		bits := strings.Split(line, " ")
		opponentPlay := parseStrategy1(bits[0])
		ourPlay := parseStrategy1(bits[1])

		resultScore := getPlayResultScore(ourPlay, opponentPlay)
		playScore := getPlayScore(ourPlay)

		totalScore += resultScore + playScore
	}
	return totalScore
}

func computePart2(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		bits := strings.Split(line, " ")
		opponentPlay := parseStrategy1(bits[0])
		ourPlay := parseStrategy2(bits[1], opponentPlay)

		resultScore := getPlayResultScore(ourPlay, opponentPlay)
		playScore := getPlayScore(ourPlay)

		totalScore += resultScore + playScore
	}
	return totalScore
}
