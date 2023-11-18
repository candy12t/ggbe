package bootrom

type BootROM struct {
	rom    []uint8
	active bool
}

func New(rom []uint8) *BootROM {
	return &BootROM{
		rom:    rom,
		active: true,
	}
}

func (b *BootROM) Read(addr uint16) uint8 {
	return b.rom[addr]
}

func (b *BootROM) Write(_ uint16, val uint8) {
	b.active = b.active && (val == 0)
}

func (b *BootROM) IsActive() bool {
	return b.active
}
