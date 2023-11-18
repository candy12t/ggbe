package peripherals

import (
	"testing"

	"github.com/candy12t/ggbe/internal/bootrom"
)

func TestRead(t *testing.T) {
	rom := make([]uint8, 256)
	b := bootrom.New(rom)
	p := New(b)

	tests := []struct {
		addr uint16
		want uint8
	}{
		{addr: 0x0000, want: 0x00},
		{addr: 0x00FF, want: 0x00},
		{addr: 0xC000, want: 0x00},
		{addr: 0xFDFF, want: 0x00},
		{addr: 0xFF80, want: 0x00},
		{addr: 0xFFFE, want: 0x00},
		{addr: 0xFFFF, want: 0xFF},
	}

	for _, tt := range tests {
		if got := p.Read(tt.addr); got != tt.want {
			t.Errorf("b.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestWrite(t *testing.T) {
	rom := make([]uint8, 256)
	b := bootrom.New(rom)
	p := New(b)

	tests := []struct {
		addr uint16
		val  uint8
		want uint8
	}{
		{addr: 0xC000, val: 0x01, want: 0x01},
		{addr: 0xFDFF, val: 0x02, want: 0x02},
		{addr: 0xFF80, val: 0x03, want: 0x03},
		{addr: 0xFFFE, val: 0x04, want: 0x04},
	}

	for _, tt := range tests {
		p.Write(tt.addr, tt.val)
		if got := p.Read(tt.addr); got != tt.want {
			t.Errorf("b.Read(%v) returned %v, want %v\n", tt.addr, got, tt.want)
		}
	}
}

func TestWrite_DisableBootROM(t *testing.T) {
	rom := make([]uint8, 256)
	b := bootrom.New(rom)
	p := New(b)

	if got := p.bootrom.IsActive(); got != true {
		t.Errorf("p.bootrom.IsActive() returned %v, want %v\n", got, true)
	}

	p.Write(0xFF50, 0x01)
	if got := p.bootrom.IsActive(); got != false {
		t.Errorf("p.bootrom.IsActive() returned %v, want %v\n", got, false)
	}

	if got := p.Read(0x0000); got != 0xFF {
		t.Errorf("b.Read(%v) returned %v, want %v\n", 0x0000, got, 0xFF)
	}
}
