package cpu

import (
	"github.com/candy12t/ggbe/internal/peripherals"
)

type IO8 interface {
	Read(cpu *CPU, bus *peripherals.Peripherals) (uint8, bool)
	Write(cpu *CPU, bus *peripherals.Peripherals, val uint8) bool
}

type IO16 interface {
	Read(cpu *CPU, bus *peripherals.Peripherals) (uint16, bool)
	Write(cpu *CPU, bus *peripherals.Peripherals, val uint16) bool
}

func (c *CPU) Read8(bus *peripherals.Peripherals, src IO8) (uint8, bool) {
	return src.Read(c, bus)
}

func (c *CPU) Write8(bus *peripherals.Peripherals, dst IO8, val uint8) bool {
	return dst.Write(c, bus, val)
}

func (c *CPU) Read16(bus *peripherals.Peripherals, src IO16) (uint16, bool) {
	return src.Read(c, bus)
}

func (c *CPU) Write16(bus *peripherals.Peripherals, dst IO16, val uint16) bool {
	return dst.Write(c, bus, val)
}
