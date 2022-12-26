package main

import (
	"aoc22/util"
	"testing"
)

func TestTreeVisibility(t *testing.T) {
	var scanner *util.FileScanner = util.ReadFileScannerFromLocalPath("test_input.txt")
	forest := inputToArray(scanner)

	exptectedOutput := [][]bool{
		{true, true, true, true, true},
		{true, true, true, false, true},
		{true, true, false, true, true},
		{true, false, true, false, true},
		{true, true, true, true, true},
	}

	visible := markVisibleTrees(forest)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if visible[i][j] != exptectedOutput[i][j] {
				t.Errorf("Index [%d, %d] is marked as %t but we expected %t",
					i, j, visible[i][j], exptectedOutput[i][j])
			}
		}
	}

	trees_count := countVisibleTrees(visible)
	if trees_count != 21 {
		t.Errorf("Visible trees %d != 21", trees_count)
	}
}
