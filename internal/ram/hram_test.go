package ram

import (
	"testing"
)

func TestHRAM_Read(t *testing.T) {
	r := NewHRAM()

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x00, want: 0},
		{addr: 0x7F, want: 0},
	}

	for _, tt := range tests {
		if got := r.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestHRAM_Write(t *testing.T) {
	r := NewHRAM()

	tests := []struct {
		addr uint16
		val  uint8
		want uint8
	}{
		{addr: 0x00, val: 1, want: 1},
		{addr: 0x7F, val: 2, want: 2},
	}

	for _, tt := range tests {
		r.Write(tt.addr, tt.val)
		if got := r.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestHRAM_OverAddr(t *testing.T) {
	r := NewHRAM()

	if got := r.Read(0x80); got != 0 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x80, got, 0)
	}

	r.Write(0x80, 1)
	if got := r.Read(0x80); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x80, got, 1)
	}
	if got := r.Read(0x00); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x00, got, 1)
	}
}
