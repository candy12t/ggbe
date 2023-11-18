package wram

import "testing"

func TestRead(t *testing.T) {
	w := New()

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x0000, want: 0},
		{addr: 0x2000, want: 0},
	}

	for _, tt := range tests {
		if got := w.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestWrite(t *testing.T) {
	w := New()

	tests := []struct {
		addr uint16
		val  uint8
		want uint8
	}{
		{addr: 0x0000, val: 1, want: 1},
		{addr: 0x2000, val: 2, want: 2},
	}

	for _, tt := range tests {
		w.Write(tt.addr, tt.val)
		if got := w.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestOverAddr(t *testing.T) {
	w := New()

	if got := w.Read(0x2001); got != 0 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x2001, got, 0)
	}

	w.Write(0x2001, 1)
	if got := w.Read(0x2001); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x2001, got, 1)
	}
	if got := w.Read(0x01); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x01, got, 1)
	}
}
