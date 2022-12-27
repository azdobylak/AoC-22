package main

import (
	"aoc22/util"
	"testing"
)

func TestScenicScore01(t *testing.T) {
	var scanner *util.FileScanner = util.ReadFileScannerFromLocalPath("test_input.txt")
	forest := inputToArray(scanner)

	score := calculateMaxScenicScore(forest)

	if score != 8 {
		t.Errorf("Incorrect score %d != %d", 8, score)
	}
}
