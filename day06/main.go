package main

import (
	"aoc22/util"
	"fmt"
)

func indexOfFirstUniqueNChars(word string, n int) int {
	for i := 0; i < len(word)-n; i++ {
		charCounter := make(map[rune]bool)
		chars := []rune(word[i : i+n])
		for _, c := range chars {
			charCounter[c] = true
		}
		if len(charCounter) == n {
			return i + n
		}
	}
	return -1
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var indx int
	for scanner.Scan() {
		line := scanner.Text()
		indx = indexOfFirstUniqueNChars(line, 4)
	}
	return indx
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var indx int
	for scanner.Scan() {
		line := scanner.Text()
		indx = indexOfFirstUniqueNChars(line, 14)
	}

	return indx
}

func main() {
	fmt.Println("=== DAY 06a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 06b solution ===")
	fmt.Println(solveB())
}
