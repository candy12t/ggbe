package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type Imm8 struct {
	rs *State
}

var _ IO8 = (*Imm8)(nil)

var Imm8Instance = New()

func New() *Imm8 {
	return &Imm8{
		rs: NewState(),
	}
}

func (i *Imm8) Read(cpu *CPU, bus *peripherals.Peripherals) (uint8, bool) {
	switch i.rs.Step() {
	case 0:
		i.rs.SetVal8(bus.Read(cpu.Regs.PC))
		cpu.Regs.PC++
		i.rs.SetStep(1)
		return 0, false
	case 1:
		v := i.rs.Val8()
		i.rs.SetStep(0)
		return v, true
	default:
		panic("unreachable")
	}
}

func (i *Imm8) Write(cpu *CPU, bus *peripherals.Peripherals, val uint8) bool {
	panic("unreachable")
}
