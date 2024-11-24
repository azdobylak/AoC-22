package main

import (
	"aoc22/util"
	"fmt"
)

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	state := NewOperationState()
	for scanner.Scan() {
		line := scanner.Text()
		state.Process(line)
	}
	signalStrengths := [6]int{
		state.GetSignalStrength(20),
		state.GetSignalStrength(60),
		state.GetSignalStrength(100),
		state.GetSignalStrength(140),
		state.GetSignalStrength(180),
		state.GetSignalStrength(220),
	}
	var strengthsSum int
	for i := 0; i < len(signalStrengths); i++ {
		strengthsSum += signalStrengths[i]
	}

	return strengthsSum
}

func main() {
	fmt.Println("=== DAY 10a solution ===")
	fmt.Println(solveA())
	//fmt.Println("=== DAY 10b solution ===")
	//fmt.Println(solveB())
}
