package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type direct8 int

const (
	Direct8D direct8 = iota + 1
	Direct8DFF
)

type Direct8 struct {
	rs      *State
	ws      *State
	direct8 direct8
}

var _ IO8 = (*Direct8)(nil)

func (i *Direct8) Read(cpu *CPU, bus *peripherals.Peripherals) (uint8, bool) {
	switch i.rs.Step() {
	case 0:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.rs.SetVal8(v)
			i.rs.SetStep(1)

			if i.direct8 == Direct8DFF {
				i.rs.SetVal16(0xFF00 | uint16(v))
				i.rs.SetStep(2)
			}
		}
		return 0, false
	case 1:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.rs.SetVal16(uint16(v)<<8 | uint16(i.rs.Val8()))
			i.rs.SetStep(2)
		}
		return 0, false
	case 2:
		i.rs.SetVal8(bus.Read(i.rs.Val16()))
		i.rs.SetStep(3)
		return 0, false
	case 3:
		v := i.rs.Val8()
		i.rs.SetStep(0)
		return v, true
	default:
		panic("unreachable")
	}
}

func (i *Direct8) Write(cpu *CPU, bus *peripherals.Peripherals, val uint8) bool {
	switch i.ws.Step() {
	case 0:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.ws.SetVal8(v)
			i.ws.SetStep(1)

			if i.direct8 == Direct8DFF {
				i.ws.SetVal16(0xFF00 | uint16(v))
				i.ws.SetStep(2)
			}
		}
		return false
	case 1:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.ws.SetVal16(uint16(v)<<8 | uint16(i.ws.Val8()))
			i.ws.SetStep(2)
		}
		return false
	case 2:
		bus.Write(i.ws.Val16(), val)
		i.ws.SetStep(3)
		return false
	case 3:
		i.ws.SetStep(0)
		return true
	default:
		panic("unreachable")
	}
}
