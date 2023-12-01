package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type Direct16 struct {
	ws *State
}

var _ IO16 = (*Direct16)(nil)

func (i *Direct16) Read(cpu *CPU, bus *peripherals.Peripherals) (uint16, bool) {
	panic("unreachable")
}

func (i *Direct16) Write(cpu *CPU, bus *peripherals.Peripherals, val uint16) bool {
	switch i.ws.Step() {
	case 0:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.ws.SetVal8(v)
			i.ws.SetStep(1)
		}
		return false
	case 1:
		if v, ok := cpu.Read8(bus, Imm8Instance); ok {
			i.ws.SetVal16(uint16(v)<<8 | uint16(i.ws.Val8()))
			i.ws.SetStep(2)
		}
		return false
	case 2:
		bus.Write(i.ws.Val16(), uint8(val))
		i.ws.SetStep(3)
		return false
	case 3:
		bus.Write(i.ws.Val16()+1, uint8(val>>8))
		i.ws.SetStep(4)
		return false
	case 4:
		i.ws.SetStep(0)
		return true
	default:
		panic("unreachable")
	}
}
