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

func findCommonItem(items string, otherItems []string) rune {
	firstInventory := itemsToMap(items)
	var otherInventories []map[rune]bool
	for _, _items := range otherItems {
		otherInventories = append(otherInventories, itemsToMap(_items))
	}
	var commonItems []rune

	for key, _ := range firstInventory {
		var includesAll = true
		for _, invent := range otherInventories {
			if !invent[key] {
				includesAll = false
				break
			}
		}
		if includesAll {
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

		commonItem := findCommonItem(firstHalf, []string{secondHalf})
		prioritySum += int(itemToPriority(commonItem))
	}

	return prioritySum
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var prioritySum int
	n := 0
	items := [3]string{}

	for scanner.Scan() {
		items[n%3] = scanner.Text()

		if (n+1)%3 == 0 {
			commonItem := findCommonItem(items[0], items[1:3])
			prioritySum += int(itemToPriority(commonItem))
		}
		n++
	}

	return prioritySum
}

func main() {
	fmt.Println("=== DAY 03a solution ===")
	fmt.Println("sumPoints=", solveA())
	fmt.Println("=== DAY 03b solution ===")
	fmt.Println("sumPoints=", solveB())
}
