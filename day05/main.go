package main

import (
	"aoc22/util"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func toStackReversed(elements []string) *Stack[string] {
	var items []string
	for i := len(elements) - 1; i >= 0; i-- {
		items = append(items, elements[i])
	}
	return &Stack[string]{items: items}
}

type Stack[T interface{}] struct {
	items []T
}

func createInitialStacks() map[int]*Stack[string] {
	var stacks map[int]*Stack[string]
	stacks = map[int]*Stack[string]{
		1: toStackReversed([]string{"D", "Z", "T", "H"}),
		2: toStackReversed([]string{"S", "C", "G", "T", "W", "R", "Q"}),
		3: toStackReversed([]string{"H", "C", "R", "N", "Q", "F", "B", "P"}),
		4: toStackReversed([]string{"Z", "H", "F", "N", "C", "L"}),
		5: toStackReversed([]string{"S", "Q", "F", "L", "G"}),
		6: toStackReversed([]string{"S", "C", "R", "B", "Z", "W", "P", "V"}),
		7: toStackReversed([]string{"J", "F", "Z"}),
		8: toStackReversed([]string{"Q", "H", "R", "Z", "V", "L", "D"}),
		9: toStackReversed([]string{"D", "L", "Z", "F", "N", "G", "H", "B"}),
	}

	return stacks
}

func (this *Stack[T]) Push(item T) {
	this.items = append(this.items, item)
}

func (this *Stack[T]) Peek() T {
	return this.items[len(this.items)-1]
}

func (this *Stack[T]) Pop() T {
	i := len(this.items)

	el := this.items[i-1]
	this.items = this.items[:i-1]

	return el
}

func move(stacks map[int]*Stack[string], order MoveOrder) {
	for i := 1; i <= order.number; i++ {
		el := stacks[order.from].Pop()
		stacks[order.to].Push(el)
	}
}

func moveGroup(stacks map[int]*Stack[string], order MoveOrder) {
	items := make([]string, order.number)
	for i := order.number - 1; i >= 0; i-- {
		items[i] = stacks[order.from].Pop()
	}

	for i := 0; i < len(items); i++ {
		stacks[order.to].Push(items[i])
	}
}

type MoveOrder struct {
	number int
	from   int
	to     int
}

func parseInputOrder(line string) MoveOrder {
	re := regexp.MustCompile("[0-9]+")
	inp := re.FindAllString(line, -1)

	n, _ := strconv.Atoi(inp[0])
	from, _ := strconv.Atoi(inp[1])
	to, _ := strconv.Atoi(inp[2])

	return MoveOrder{number: n, from: from, to: to}
}

func solveA() string {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	stacks := createInitialStacks()
	for scanner.Scan() {
		line := scanner.Text()
		order := parseInputOrder(line)
		move(stacks, order)
	}
	var out bytes.Buffer
	for i := 1; i <= len(stacks); i++ {
		out.WriteString(stacks[i].Peek())
	}
	return out.String()
}

func solveB() string {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	stacks := createInitialStacks()
	for scanner.Scan() {
		line := scanner.Text()
		order := parseInputOrder(line)
		moveGroup(stacks, order)
	}

	var out bytes.Buffer
	for i := 1; i <= len(stacks); i++ {
		out.WriteString(stacks[i].Peek())
	}
	return out.String()
}

func main() {

	fmt.Println("=== DAY 05a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 05b solution ===")
	fmt.Println(solveB())
}
