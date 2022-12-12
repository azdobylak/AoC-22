package main

import (
	"fmt"
	"sort"
	"strconv"
    "aoc22/util"
)

func main() {
    var scanner *util.FileScanner = util.ReadFileScanner()
    defer scanner.Close()

	elfCalories := make(map[int]int)
	elfIndex := 1

	maxCalories := 0
	maxCaloriesIndex := elfIndex

	for scanner.Scan() {
		line := scanner.Text()
		if string(line) == "" {
			if elfCalories[elfIndex] > maxCalories {
				maxCalories = elfCalories[elfIndex]
				maxCaloriesIndex = elfIndex
			}

			elfIndex += 1
		} else {
			numCals, _ := strconv.Atoi(line)
			elfCalories[elfIndex] += numCals
		}
	}
	fmt.Println("=== DAY 01a solution ===")
	fmt.Println(maxCaloriesIndex, "index ->", maxCalories, " calories")

	vals := make([]int, 0, len(elfCalories))
	for _, v := range elfCalories {
		vals = append(vals, v)
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	var valsSum int
	fmt.Println("=== DAY 01b solution ===")
	fmt.Println("Top 3 values:", vals[0:3])
	for _, v := range vals[0:3] {
		valsSum += v
	}
	fmt.Println("Top 3 sum =", valsSum)
}
