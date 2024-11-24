package main

import (
	"fmt"
	"testing"
)

func assertStateValue(t *testing.T, state *OperationState, expCycle int, expValue int) {
	if state.value != expValue {
		t.Errorf("Operation state value is incorrect: (%d != %d) expected", state.value, expValue)
	}
	if state.cycle != expCycle {
		t.Errorf("Operation state cycle is incorrect: (%d != %d) expected", state.cycle, expCycle)
	}
}

func assertSignalStrength(t *testing.T, state *OperationState, cycle int, expSignalStrength int) {
	signalStrength := state.GetSignalStrength(cycle)
	if signalStrength != expSignalStrength {
		t.Errorf("Signal strength is incorrect: (%d != %d) expected", signalStrength, expSignalStrength)
	}
}

func TestSmallInputDay10(t *testing.T) {
	os := NewOperationState()
	fmt.Println(os.value)

	assertStateValue(t, os, 1, 1)
	os.Process("noop")
	assertStateValue(t, os, 2, 1)
	os.Process("addx 3")
	assertStateValue(t, os, 4, 4)
	os.Process("addx -5")
	assertStateValue(t, os, 6, -1)
}

func TestDay10A(t *testing.T) {
	instructions := [...]string{
		"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4", "noop",
		"addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx -35",
		"addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15",
		"noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5", "noop",
		"noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7", "noop", "noop",
		"noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop", "addx 1", "noop",
		"noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7", "noop", "addx 1", "addx -33",
		"noop", "noop", "noop", "addx 2", "noop", "noop", "noop", "addx 8", "noop", "addx -1",
		"addx 2", "addx 1", "noop", "addx 17", "addx -9", "addx 1", "addx 1", "addx -3", "addx 11", "noop",
		"noop", "addx 1", "noop", "addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3",
		"addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop", "addx -9",
		"addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop", "noop", "addx -1",
		"addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15", "addx -21", "addx 22", "addx -6", "addx 1",
		"noop", "addx 2", "addx 1", "noop", "addx -10", "noop", "noop", "addx 20", "addx 1", "addx 2",
		"addx 2", "addx -6", "addx -11", "noop", "noop", "noop",
	}
	os := NewOperationState()
	for _, instr := range instructions {
		os.Process(instr)
	}
	assertSignalStrength(t, os, 20, 420)
	assertSignalStrength(t, os, 60, 1140)
	assertSignalStrength(t, os, 100, 1800)
	assertSignalStrength(t, os, 140, 2940)
	assertSignalStrength(t, os, 180, 2880)
	assertSignalStrength(t, os, 220, 3960)
}
