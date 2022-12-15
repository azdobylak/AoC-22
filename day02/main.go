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

func pickWinningSymbol(choice Symbol) Symbol {
	var result Symbol
	switch choice {
	case SCISSORS:
		result = ROCK
	case ROCK:
		result = PAPER
	case PAPER:
		result = SCISSORS
	}
	return result
}

func pickLoosingSymbol(choice Symbol) Symbol {
	var result Symbol
	switch choice {
	case SCISSORS:
		result = PAPER
	case ROCK:
		result = SCISSORS
	case PAPER:
		result = ROCK
	}
	return result
}

func assignFixedChoice(playerInput string, opponent Symbol) Symbol {
	var playerResponse Symbol
	if playerInput == "X" {
		playerResponse = pickLoosingSymbol(opponent)
	} else if playerInput == "Y" {
		playerResponse = opponent
	} else if playerInput == "Z" {
		playerResponse = pickWinningSymbol(opponent)
	} else {
		panic("Wrong input: " + playerInput)
	}

	return playerResponse
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
		panic("Unhandled scenario: " + string(playerChoice) + " " + string(opponentChoice))
	}
	return int(result) + int(playerChoice)
}

func solveA() int {
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
	return sumPoints
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var sumPoints int
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Split(line, " ")

		opponentSymbol := readChoice(choices[0])
		mySymbol := assignFixedChoice(choices[1], opponentSymbol)

		sumPoints += calculateMatchPoints(mySymbol, opponentSymbol)
	}
	return sumPoints
}

func main() {
	fmt.Println("=== DAY 02a solution ===")
	fmt.Println("sumPoints=", solveA())
	fmt.Println("=== DAY 02b solution ===")
	fmt.Println("sumPoints=", solveB())
}
