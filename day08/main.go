package main

import (
	"aoc22/util"
	"fmt"
)

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		matrix[i] = rows[startRow : startRow+m]
	}
	return matrix
}

func markVisibleTrees(forest [][]uint8) [][]bool {
	n_rows, n_cols := len(forest), len(forest[0])
	visible := Make2D[bool](n_rows, n_cols)

	var max_v uint8
	// left to right
	for i := 1; i < n_rows-1; i++ {
		max_v = forest[i][0]
		for j := 1; j < n_cols; j++ {
			if forest[i][j] > max_v {
				max_v = forest[i][j]
				visible[i][j] = true
			}
		}
	}
	// right to left
	for i := 1; i < n_rows-1; i++ {
		max_v = forest[i][n_cols-1]
		for j := n_cols - 2; j >= 0; j-- {
			if forest[i][j] > max_v {
				max_v = forest[i][j]
				visible[i][j] = true
			}
		}
	}

	// top to bottom
	for j := 1; j < n_cols-1; j++ {
		max_v = forest[0][j]
		for i := 1; i < n_rows; i++ {
			if forest[i][j] > max_v {
				max_v = forest[i][j]
				visible[i][j] = true
			}
		}
	}

	// bottom to top
	for j := 1; j < n_cols; j++ {
		max_v = forest[n_rows-1][j]
		for i := n_rows - 2; i >= 0; i-- {
			if forest[i][j] > max_v {
				max_v = forest[i][j]
				visible[i][j] = true
			}
		}
	}

	// border
	for i := 0; i < n_rows; i++ {
		visible[i][0] = true
		visible[i][n_cols-1] = true
	}

	for j := 0; j < n_cols; j++ {
		visible[0][j] = true
		visible[n_rows-1][j] = true
	}

	return visible
}

func inputToArray(scanner *util.FileScanner) [][]uint8 {
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	forest := Make2D[uint8](len(lines), len(lines[0]))

	for i, line := range lines {
		for j, c := range line {
			forest[i][j] = uint8(c - 48)
		}
	}

	return forest
}

func countVisibleTrees(visible [][]bool) int {
	n_rows, n_cols := len(visible), len(visible[0])

	var sum int
	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {
			if visible[i][j] == true {
				sum += 1
			}
		}
	}
	return sum
}

func calculateMaxScenicScore(forest [][]uint8) int {
	n_rows, n_cols := len(forest), len(forest[0])

	var maxScore int
	var currentScore int

	var currentHeight uint8

	var scoreLeft int
	var scoreRight int
	var scoreUp int
	var scoreDown int
	for i := 1; i < n_rows-1; i++ {
		for j := 1; j < n_cols-1; j++ {
			currentHeight = forest[i][j]
			scoreDown = 0
			scoreUp = 0
			scoreRight = 0
			scoreLeft = 0

			// to right
			for r := j + 1; r < n_cols; r++ {
				scoreRight += 1
				if forest[i][r] >= currentHeight {
					break
				}
			}
			// to left
			for l := j - 1; l >= 0; l-- {
				scoreLeft += 1
				if forest[i][l] >= currentHeight {
					break
				}
			}
			// to bottom
			for b := i + 1; b < n_rows; b++ {
				scoreDown += 1
				if forest[b][j] >= currentHeight {
					break
				}
			}
			// to top
			for t := i - 1; t >= 0; t-- {
				scoreUp += 1
				if forest[t][j] >= currentHeight {
					break
				}
			}

			currentScore = scoreUp * scoreDown * scoreLeft * scoreRight

			if currentScore > maxScore {
				maxScore = currentScore
			}

		}
	}

	return maxScore
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	forest := inputToArray(scanner)
	visible := markVisibleTrees(forest)
	return countVisibleTrees(visible)
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	forest := inputToArray(scanner)

	return calculateMaxScenicScore(forest)
}

func main() {
	fmt.Println("=== DAY 07a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 07b solution ===")
	fmt.Println(solveB())
}
