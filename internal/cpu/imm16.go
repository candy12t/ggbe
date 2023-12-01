package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type Imm16 struct {
	rs *State
}

var _ IO16 = (*Imm16)(nil)

func (i *Imm16) Read(cpu *CPU, bus *peripherals.Peripherals) (uint16, bool) {
	switch i.rs.Step() {
	case 0:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.rs.SetVal8(v)
			i.rs.SetStep(1)
		}
		return 0, false
	case 1:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.rs.SetVal16(uint16(v)<<8 | uint16(i.rs.Val8()))
			i.rs.SetStep(2)
		}
		return 0, false
	case 2:
		v := i.rs.Val16()
		i.rs.SetStep(0)
		return v, true
	default:
		panic("unreachable")
	}
}

func (i *Imm16) Write(cpu *CPU, bus *peripherals.Peripherals, val uint16) bool {
	panic("unreachable")
}
