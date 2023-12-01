package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type indirect int

const (
	IndirectBC indirect = iota + 1
	IndirectDE
	IndirectHL
	IndirectCFF
	IndirectHLD
	IndirectHLI
)

type Indirect struct {
	rs       *State
	ws       *State
	indirect indirect
}

var _ IO8 = (*Indirect)(nil)

func (i *Indirect) Read(cpu *CPU, bus *peripherals.Peripherals) (uint8, bool) {
	switch i.rs.Step() {
	case 0:
		v := i.read(cpu, bus)
		i.rs.SetVal8(v)
		i.rs.SetStep(1)
		return 0, false
	case 1:
		v := i.rs.Val8()
		i.rs.SetStep(1)
		return v, true
	default:
		panic("unreachable")
	}
}

func (i *Indirect) Write(cpu *CPU, bus *peripherals.Peripherals, val uint8) bool {
	switch i.ws.Step() {
	case 0:
		i.write(cpu, bus, val)
		i.ws.SetStep(1)
		return false
	case 1:
		i.ws.SetStep(0)
		return true
	default:
		panic("unreachable")
	}
}

func (i *Indirect) write(cpu *CPU, bus *peripherals.Peripherals, val uint8) {
	switch i.indirect {
	case IndirectBC:
		bus.Write(cpu.Regs.BC(), val)
	case IndirectDE:
		bus.Write(cpu.Regs.DE(), val)
	case IndirectHL:
		bus.Write(cpu.Regs.HL(), val)
	case IndirectCFF:
		bus.Write(0xFF00|uint16(cpu.Regs.C), val)
	case IndirectHLD:
		addr := cpu.Regs.HL()
		cpu.Regs.WriteHL(addr - 1)
		bus.Write(addr, val)
	case IndirectHLI:
		addr := cpu.Regs.HL()
		cpu.Regs.WriteHL(addr + 1)
		bus.Write(addr, val)
	default:
		panic("unreachable")
	}
}

func (i *Indirect) read(cpu *CPU, bus *peripherals.Peripherals) uint8 {
	switch i.indirect {
	case IndirectBC:
		return bus.Read(cpu.Regs.BC())
	case IndirectDE:
		return bus.Read(cpu.Regs.DE())
	case IndirectHL:
		return bus.Read(cpu.Regs.HL())
	case IndirectCFF:
		return bus.Read(0xFF00 | uint16(cpu.Regs.C))
	case IndirectHLD:
		addr := cpu.Regs.HL()
		cpu.Regs.WriteHL(addr - 1)
		return bus.Read(addr)
	case IndirectHLI:
		addr := cpu.Regs.HL()
		cpu.Regs.WriteHL(addr + 1)
		return bus.Read(addr)
	default:
		panic("unreachable")
	}
}
