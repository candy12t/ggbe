package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type Reg16 int

var _ IO16 = (*Reg16)(nil)

const (
	Reg16AF Reg16 = iota + 1
	Reg16BC
	Reg16DE
	Reg16HL
	Reg16SP
)

func (r Reg16) Read(cpu *CPU, _ *peripherals.Peripherals) (uint16, bool) {
	switch r {
	case Reg16AF:
		return cpu.Regs.AF(), true
	case Reg16BC:
		return cpu.Regs.BC(), true
	case Reg16DE:
		return cpu.Regs.DE(), true
	case Reg16HL:
		return cpu.Regs.HL(), true
	case Reg16SP:
		return cpu.Regs.SP, true
	default:
		panic("unreachable")
	}
}

func (r Reg16) Write(cpu *CPU, _ *peripherals.Peripherals, val uint16) bool {
	switch r {
	case Reg16AF:
		cpu.Regs.WriteAF(val)
	case Reg16BC:
		cpu.Regs.WriteBC(val)
	case Reg16DE:
		cpu.Regs.WriteDE(val)
	case Reg16HL:
		cpu.Regs.WriteHL(val)
	case Reg16SP:
		cpu.Regs.SP = val
	default:
		panic("unreachable")
	}
	return true
}
