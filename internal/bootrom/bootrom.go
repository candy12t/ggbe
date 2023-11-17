package bootrom

type BootROM struct {
	rom []uint8
}

func New(rom []uint8) *BootROM {
	return &BootROM{
		rom: rom,
	}
}

func (b *BootROM) Read(addr uint16) uint8 {
	return b.rom[addr]
}
