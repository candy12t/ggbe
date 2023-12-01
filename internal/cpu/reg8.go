package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type Reg8 int

var _ IO8 = (*Reg8)(nil)

const (
	Reg8A Reg8 = iota + 1
	Reg8B
	Reg8C
	Reg8D
	Reg8E
	Reg8H
	Reg8L
)

func (r Reg8) Read(cpu *CPU, _ *peripherals.Peripherals) (uint8, bool) {
	switch r {
	case Reg8A:
		return cpu.Regs.A, true
	case Reg8B:
		return cpu.Regs.B, true
	case Reg8C:
		return cpu.Regs.C, true
	case Reg8D:
		return cpu.Regs.D, true
	case Reg8E:
		return cpu.Regs.E, true
	case Reg8H:
		return cpu.Regs.H, true
	case Reg8L:
		return cpu.Regs.L, true
	default:
		panic("unreachable")
	}
}

func (r Reg8) Write(cpu *CPU, _ *peripherals.Peripherals, val uint8) bool {
	switch r {
	case Reg8A:
		cpu.Regs.A = val
	case Reg8B:
		cpu.Regs.B = val
	case Reg8C:
		cpu.Regs.C = val
	case Reg8D:
		cpu.Regs.D = val
	case Reg8E:
		cpu.Regs.E = val
	case Reg8H:
		cpu.Regs.H = val
	case Reg8L:
		cpu.Regs.L = val
	default:
		panic("unreachable")
	}
	return true
}
