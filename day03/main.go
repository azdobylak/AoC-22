package main

import (
	"aoc22/util"
	"fmt"
)

func itemsToMap(items string) map[rune]bool {
	inventory := make(map[rune]bool)
	for _, char := range items {
		inventory[char] = true
	}
	return inventory
}

func findCommonItem(firstHalf string, secondHalf string) rune {
	firstInventory := itemsToMap(firstHalf)
	secondInventory := itemsToMap(secondHalf)

	var commonItems []rune

	for key, _ := range firstInventory {
		if secondInventory[key] {
			commonItems = append(commonItems, key)
		}
	}

	if len(commonItems) > 1 {
		panic("Too many common characters: " + string(commonItems))
	}

	return commonItems[0]
}

func itemToPriority(item rune) rune {
	if item >= 'A' && item <= 'Z' {
		return item - 38
	} else if item >= 'a' && item <= 'z' {
		return item - 96
	} else {
		panic("Incorrect character: " + string(item))
	}
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var prioritySum int

	for scanner.Scan() {
		line := scanner.Text()
		halfLength := len(line) / 2

		firstHalf := line[:halfLength]
		secondHalf := line[halfLength:]

		commonItem := findCommonItem(firstHalf, secondHalf)
		prioritySum += int(itemToPriority(commonItem))
	}

	return prioritySum
}

func main() {
	fmt.Println("=== DAY 03a solution ===")
	fmt.Println("sumPoints=", solveA())
	//fmt.Println("=== DAY 03b solution ===")
	//fmt.Println("sumPoints=", solveB())
}
