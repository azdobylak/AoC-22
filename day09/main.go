package main

import (
	"aoc22/day09/rope"
	"aoc22/util"
	"fmt"
	"strconv"
	"strings"
)

func execute(head, tail *rope.Knot, cmd string, visited *map[rope.Knot]bool) {
	words := strings.Split(cmd, " ")
	n, err := strconv.Atoi(words[1])
	if err != nil {
		panic(err)
	}
	direction := rope.NewDirection(words[0])

	for i := 0; i < n; i++ {
		move(head, tail, direction)
		(*visited)[*tail] = true
	}
}

func move(head, tail *rope.Knot, d rope.Direction) {
	head.Move(d)
	tail.Follow(*head)
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	var H rope.Knot = rope.NewKnot("H")
	var T rope.Knot = rope.NewKnot("T")

	visited := make(map[rope.Knot]bool, 200)

	for scanner.Scan() {
		line := scanner.Text()
		execute(&H, &T, line, &visited)
	}

	return len(visited)
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	return -1
}

func main() {
	fmt.Println("=== DAY 07a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 07b solution ===")
	fmt.Println(solveB())
}
