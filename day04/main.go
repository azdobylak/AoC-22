package main

import (
	"aoc22/util"
	"fmt"
	"strconv"
	"strings"
)

type SectionsRange struct {
	min int
	max int
}

func newSectionsRange(input string) *SectionsRange {
	sections := strings.Split(input, "-")
	min, err1 := strconv.Atoi(sections[0])
	max, err2 := strconv.Atoi(sections[1])
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		panic("Unable to parse " + input)
	}
	return &SectionsRange{min: min, max: max}
}

func (r *SectionsRange) Contains(other *SectionsRange) bool {
	return r.min <= other.min && r.max >= other.max
}

func (r *SectionsRange) Overlaps(other *SectionsRange) bool {
	return !(r.min > other.max || r.max < other.min)
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var overlapingCount int
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")
		r1 := newSectionsRange(inputs[0])
		r2 := newSectionsRange(inputs[1])

		if r1.Contains(r2) || r2.Contains(r1) {
			overlapingCount += 1
		}
	}
	return overlapingCount
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var overlapingCount int
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")
		r1 := newSectionsRange(inputs[0])
		r2 := newSectionsRange(inputs[1])

		if r1.Overlaps(r2) {
			overlapingCount += 1
		}
	}

	return overlapingCount
}

func main() {
	fmt.Println("=== DAY 04a solution ===")
	fmt.Println("sumPoints=", solveA())
	fmt.Println("=== DAY 04b solution ===")
	fmt.Println("sumPoints=", solveB())
}
