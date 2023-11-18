package peripherals

import (
	"github.com/candy12t/ggbe/internal/bootrom"
	"github.com/candy12t/ggbe/internal/hram"
	"github.com/candy12t/ggbe/internal/wram"
)

// Peripherals is MMIO (memory mapped I/O)
type Peripherals struct {
	bootrom *bootrom.BootROM
	wram    *wram.WRAM
	hram    *hram.HRAM
}

func New(bootrom *bootrom.BootROM) *Peripherals {
	return &Peripherals{
		bootrom: bootrom,
		wram:    wram.New(),
		hram:    hram.New(),
	}
}

func (p *Peripherals) Read(addr uint16) uint8 {
	switch {
	case 0x0000 <= addr && addr <= 0x00FF:
		if p.bootrom.IsActive() {
			return p.bootrom.Read(addr)
		}
		return 0xFF
	case 0xC000 <= addr && addr <= 0xFDFF:
		return p.wram.Read(addr)
	case 0xFF80 <= addr && addr <= 0xFFFE:
		return p.hram.Read(addr)
	default:
		return 0xFF
	}
}

func (p *Peripherals) Write(addr uint16, val uint8) {
	switch {
	case 0xC000 <= addr && addr <= 0xFDFF:
		p.wram.Write(addr, val)
	case addr == 0xFF50:
		p.bootrom.Write(addr, val)
	case 0xFF80 <= addr && addr <= 0xFFFE:
		p.hram.Write(addr, val)
	}
}
