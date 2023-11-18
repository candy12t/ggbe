package bootrom

import (
	"testing"
)

func TestRead(t *testing.T) {
	rom := []uint8{0, 1, 2, 3, 4, 5}
	b := New(rom)

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x00, want: 0},
		{addr: 0x01, want: 1},
		{addr: 0x02, want: 2},
		{addr: 0x03, want: 3},
		{addr: 0x04, want: 4},
		{addr: 0x05, want: 5},
	}

	for _, tt := range tests {
		if got := b.Read(tt.addr); got != tt.want {
			t.Errorf("b.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestWrite(t *testing.T) {
	rom := []uint8{0, 1, 2, 3, 4, 5}
	b := New(rom)

	if got := b.IsActive(); got != true {
		t.Errorf("b.IsActive() returned %v, want %v\n", got, true)
	}

	b.Write(0x00, 0)
	if got := b.IsActive(); got != true {
		t.Errorf("b.IsActive() returned %v, want %v\n", got, true)
	}

	b.Write(0x00, 1)
	if got := b.IsActive(); got != false {
		t.Errorf("b.IsActive() returned %v, want %v\n", got, false)
	}
}
