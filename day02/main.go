package main

import (
	"aoc22/util"
	"fmt"
	"strings"
)

type Symbol int

const (
	ROCK     Symbol = 1
	PAPER           = 2
	SCISSORS        = 3
)

type MatchResult int

const (
	LOSS MatchResult = 0
	DRAW             = 3
	WIN              = 6
)

func readChoice(character string) Symbol {
	var s Symbol
	switch character {
	case "A", "X":
		s = ROCK
	case "B", "Y":
		s = PAPER
	case "C", "Z":
		s = SCISSORS
	default:
		panic("Unknonw character:" + character)
	}
	return s
}

func calculateMatchPoints(playerChoice Symbol, opponentChoice Symbol) int {
	var result MatchResult
	if playerChoice == opponentChoice {
		result = DRAW
	} else if playerChoice == ROCK {
		switch opponentChoice {
		case SCISSORS:
			result = WIN
		case PAPER:
			result = LOSS
		}
	} else if playerChoice == PAPER {
		switch opponentChoice {
		case ROCK:
			result = WIN
		case SCISSORS:
			result = LOSS
		}
	} else if playerChoice == SCISSORS {
		switch opponentChoice {
		case PAPER:
			result = WIN
		case ROCK:
			result = LOSS
		}
	} else {
		panic("Unhandled scenario")
	}
	fmt.Println(int(result) + int(playerChoice))
	return int(result) + int(playerChoice)
}

func main() {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var sumPoints int

	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Split(line, " ")

		opponentSymbol := readChoice(choices[0])
		mySymbol := readChoice(choices[1])

		sumPoints += calculateMatchPoints(mySymbol, opponentSymbol)
	}
	fmt.Println("=== DAY 02a solution ===")
	fmt.Println("sumPoints=", sumPoints)
}
