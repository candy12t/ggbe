package cpu

import (
	"testing"

	"github.com/candy12t/ggbe/internal/bootrom"
	"github.com/candy12t/ggbe/internal/peripherals"
)

func TestRead(t *testing.T) {
	bus := peripherals.New(bootrom.New([]uint8{0x01, 0x02, 0x03}))
	cpu := NewCPU()

	tests := []struct {
		wantDone1  bool
		wantDone2  bool
		wantValue2 uint8
	}{
		{
			wantDone1:  false,
			wantDone2:  true,
			wantValue2: 1,
		},
		{
			wantDone1:  false,
			wantDone2:  true,
			wantValue2: 2,
		},
		{
			wantDone1:  false,
			wantDone2:  true,
			wantValue2: 3,
		},
	}

	for _, tt := range tests {
		_, done := Imm8Instance.Read(cpu, bus)
		if done != tt.wantDone1 {
			t.Errorf("got %v, want %v\n", done, tt.wantDone1)
		}

		got, done := Imm8Instance.Read(cpu, bus)
		if done != tt.wantDone2 {
			t.Errorf("got %v, want %v\n", done, tt.wantDone2)
		}
		if got != tt.wantValue2 {
			t.Errorf("got %v, want %v\n", got, tt.wantValue2)
		}
	}
}
