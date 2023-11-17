package cpu

import (
	"fmt"

	"github.com/candy12t/ggbe/internal/peripherals"
)

type ctx struct {
	opcode uint8
	cb     bool
}

type CPU struct {
	Regs Registers
	Ctx  ctx
}

func (c *CPU) Fetch(bus *peripherals.Peripherals) {
	c.Ctx.opcode = bus.Read(c.Regs.PC)
	c.Regs.PC++
	c.Ctx.cb = false
}

func (c *CPU) Decode(bus *peripherals.Peripherals) {
	switch c.Ctx.opcode {
	case 0x00:
		c.nop(bus)
	default:
		panic(fmt.Sprintf("Not Implemented: %v\n", c.Ctx.opcode))
	}
}

func (c *CPU) EmulateCycle(bus *peripherals.Peripherals) {
	c.Decode(bus)
}

func (c *CPU) nop(bus *peripherals.Peripherals) {
	c.Fetch(bus)
}
