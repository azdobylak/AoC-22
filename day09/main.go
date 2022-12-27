package main

import (
	"aoc22/day09/rope"
	"aoc22/util"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func execute(rp *rope.Rope, cmd string, visited *map[Position]bool) {
	words := strings.Split(cmd, " ")
	n, err := strconv.Atoi(words[1])
	if err != nil {
		panic(err)
	}
	direction := rope.NewDirection(words[0])

	for i := 0; i < n; i++ {
		rp.Move(direction)
		tail := rp.Tail()
		pos := Position{x: tail.X, y: tail.Y}
		(*visited)[pos] = true
	}
}

func move(head, tail *rope.Knot, d rope.Direction) {
	head.Move(d)
	tail.Follow(*head)
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	rp := rope.NewRope(2)
	visited := make(map[Position]bool, 200)

	for scanner.Scan() {
		line := scanner.Text()
		execute(&rp, line, &visited)
	}

	return len(visited)
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	rp := rope.NewRope(10)
	visited := make(map[Position]bool, 400)

	for scanner.Scan() {
		line := scanner.Text()
		execute(&rp, line, &visited)
	}

	return len(visited)
}

func main() {
	fmt.Println("=== DAY 07a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 07b solution ===")
	fmt.Println(solveB())
}
