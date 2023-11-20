package cpu

import (
	"testing"

	"github.com/candy12t/ggbe/internal/bootrom"
	"github.com/candy12t/ggbe/internal/peripherals"
)

func TestCPU_Featch(t *testing.T) {
	b := bootrom.New([]uint8{0x12, 0x34})
	bus := peripherals.New(b)
	c := NewCPU()

	c.Fetch(bus)
	if got := c.Ctx.opcode; got != 0x12 {
		t.Errorf("c.Ctx.opcode returned %v, want %v\n", got, 0x12)
	}
	if got := c.Regs.PC; got != 0x01 {
		t.Errorf("c.Regs.PC returned %v, want %v\n", got, 0x01)
	}
}

func TestCPU_Decode_nop(t *testing.T) {
	b := bootrom.New([]uint8{0x00, 0x12, 0x34})
	bus := peripherals.New(b)
	c := NewCPU()

	c.Fetch(bus)
	c.Decode(bus)
	if got := c.Ctx.opcode; got != 0x12 {
		t.Errorf("c.Ctx.opcode returned %v, want %v\n", got, 0x12)
	}
	if got := c.Regs.PC; got != 0x02 {
		t.Errorf("c.Regs.PC returned %v, want %v\n", got, 0x02)
	}
}
