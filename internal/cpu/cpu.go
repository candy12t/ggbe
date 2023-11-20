package cpu

import (
	"errors"
	"fmt"

	"github.com/candy12t/ggbe/internal/peripherals"
)

var ErrNotImplementedOpcode = errors.New("Not Implemented opcode")

type ctx struct {
	opcode uint8
	cb     bool
}

type CPU struct {
	Regs *Registers
	Ctx  *ctx
}

func NewCPU() *CPU {
	return &CPU{
		Regs: newRegisters(),
		Ctx:  newCtx(),
	}
}

func newCtx() *ctx {
	return &ctx{}
}

func (c *CPU) Fetch(bus *peripherals.Peripherals) {
	c.Ctx.opcode = bus.Read(c.Regs.PC)
	c.Regs.PC++
	c.Ctx.cb = false
}

func (c *CPU) Decode(bus *peripherals.Peripherals) error {
	switch c.Ctx.opcode {
	case 0x00:
		c.nop(bus)
	default:
		return fmt.Errorf("%w: %v\n", ErrNotImplementedOpcode, c.Ctx.opcode)
	}
	return nil
}

func (c *CPU) EmulateCycle(bus *peripherals.Peripherals) {
	c.Decode(bus)
}

func (c *CPU) nop(bus *peripherals.Peripherals) {
	c.Fetch(bus)
}
