package hram

import (
	"testing"
)

func TestRead(t *testing.T) {
	h := New()

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x00, want: 0},
		{addr: 0x80, want: 0},
	}

	for _, tt := range tests {
		if got := h.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestWrite(t *testing.T) {
	h := New()

	tests := []struct {
		addr uint16
		val  uint8
		want uint8
	}{
		{addr: 0x00, val: 1, want: 1},
		{addr: 0x80, val: 2, want: 2},
	}

	for _, tt := range tests {
		h.Write(tt.addr, tt.val)
		if got := h.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestOverAddr(t *testing.T) {
	h := New()

	if got := h.Read(0x81); got != 0 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x81, got, 0)
	}

	h.Write(0x81, 1)
	if got := h.Read(0x81); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x81, got, 1)
	}
	if got := h.Read(0x01); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x01, got, 1)
	}
}
