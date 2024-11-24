package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	name  string
	value int
}

type OperationState struct {
	cycle           int
	value           int
	operationBuffer Operation
	valuesHistory   map[int]int
	bitmap          []bool
}

func NewOperationState() *OperationState {
	os := &OperationState{cycle: 1, value: 1}
	os.valuesHistory = make(map[int]int)
	return os
}

func loadOperation(line string) Operation {
	tokens := strings.Split(line, " ")
	var op Operation

	switch tokens[0] {
	case "noop":
		op = Operation{name: "noop"}
	case "addx":
		val, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		op = Operation{name: "addx", value: val}
	default:
		errorString := fmt.Sprintf("Unknown operation %s", tokens[0])
		panic(errorString)
	}

	return op
}

func (s *OperationState) incrementCycle() {
	s.cycle++
	if s.cycle%20 == 0 {
		s.valuesHistory[s.cycle] = s.value
	}
	var pixel bool
	paddle := (s.cycle - 1) % 40
	if paddle >= s.value-1 && paddle <= s.value+1 {
		pixel = true
	} else {
		pixel = false
	}
	s.bitmap = append(s.bitmap, pixel)
}

func (s *OperationState) GetSignalStrength(cycle int) int {
	return s.valuesHistory[cycle] * cycle
}

func (s *OperationState) Process(line string) {
	s.operationBuffer = loadOperation(line)
	s.tick()
}

func (s *OperationState) DrawSprite() {
	const width = 40
	chars := map[bool]rune{
		false: '.',
		true:  '#',
	}

	for i := 0; i < len(s.bitmap); i++ {
		if i%width == 0 {
			fmt.Println()
		}
		fmt.Printf("%c", chars[s.bitmap[i]])
	}
}

func (state *OperationState) tick() {
	state.incrementCycle()

	switch state.operationBuffer.name {
	case "noop":
	case "addx":
		state.value += state.operationBuffer.value
		state.incrementCycle()
	default:
		panic("unknown state" + state.operationBuffer.name)
	}
}
