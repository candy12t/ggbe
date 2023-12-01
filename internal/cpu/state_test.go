package cpu

import (
	"testing"
)

func TestState(t *testing.T) {
	state := NewState()

	if step := state.Step(); step != 0 {
		t.Errorf("state.Step() returned %d, want %d", step, 0)
	}

	state.SetStep(1)
	if step := state.Step(); step != 1 {
		t.Errorf("state.Step() returned %d, want %d", step, 1)
	}

	state.SetVal8(1)
	if v := state.Val8(); v != 1 {
		t.Errorf("state.Val8() returned %d, want %d", v, 1)
	}

	state.SetVal16(1)
	if v := state.Val16(); v != 1 {
		t.Errorf("state.Val16() returned %d, want %d", v, 1)
	}
}
