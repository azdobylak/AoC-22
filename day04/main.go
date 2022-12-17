package main

import (
	"aoc22/util"
	"fmt"
	"strconv"
	"strings"
)

type SectionRange struct {
	min int
	max int
}

func newSectionRange(input string) *SectionRange {
	sections := strings.Split(input, "-")
	min, err1 := strconv.Atoi(sections[0])
	max, err2 := strconv.Atoi(sections[1])
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		panic("Unable to parse " + input)
	}
	return &SectionRange{min: min, max: max}
}

func RangeContainsOther(R *SectionRange, other *SectionRange) bool {
	return R.min <= other.min && R.max >= other.max
}

func doesRangesOverlap(r1 *SectionRange, r2 *SectionRange) bool {
	return RangeContainsOther(r1, r2) || RangeContainsOther(r2, r1)
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var overlapingCount int
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")
		range1 := newSectionRange(inputs[0])
		range2 := newSectionRange(inputs[1])

		if doesRangesOverlap(range1, range2) {
			overlapingCount += 1
		}
	}
	return overlapingCount
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	for scanner.Scan() {
	}

	return 0
}

func main() {
	fmt.Println("=== DAY 04a solution ===")
	fmt.Println("sumPoints=", solveA())
	fmt.Println("=== DAY 04b solution ===")
	fmt.Println("sumPoints=", solveB())
}
