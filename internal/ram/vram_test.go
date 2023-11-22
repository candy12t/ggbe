package ram

import "testing"

func TestVRAM_Read(t *testing.T) {
	r := NewVRAM()

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x0000, want: 0},
		{addr: 0x1FFF, want: 0},
	}

	for _, tt := range tests {
		if got := r.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestVRAM_Write(t *testing.T) {
	r := NewVRAM()

	tests := []struct {
		addr uint16
		val  uint8
		want uint8
	}{
		{addr: 0x0000, val: 1, want: 1},
		{addr: 0x1FFF, val: 2, want: 2},
	}

	for _, tt := range tests {
		r.Write(tt.addr, tt.val)
		if got := r.Read(tt.addr); got != tt.want {
			t.Errorf("h.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestVRAM_OverAddr(t *testing.T) {
	r := NewVRAM()

	if got := r.Read(0x2000); got != 0 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x2000, got, 0)
	}

	r.Write(0x2000, 1)
	if got := r.Read(0x2000); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x2000, got, 1)
	}
	if got := r.Read(0x00); got != 1 {
		t.Errorf("h.Read(%v) returned %v, want %v\n", 0x00, got, 1)
	}
}
