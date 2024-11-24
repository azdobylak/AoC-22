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
}

func NewOperationState() *OperationState {
	return &OperationState{cycle: 1, value: 1}
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

func (state *OperationState) Process(line string) {
	state.operationBuffer = loadOperation(line)
	state.tick()
}

func (state *OperationState) tick() {
	state.cycle += 1

	switch state.operationBuffer.name {
	case "noop":
	case "addx":
		state.value += state.operationBuffer.value
		state.cycle += 1
	default:
		panic("unknown state" + state.operationBuffer.name)
	}
}
