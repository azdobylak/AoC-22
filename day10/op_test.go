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
