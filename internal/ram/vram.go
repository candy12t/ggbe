package ram

// VRAM is Video RAM
// the size is 8 KiB
type VRAM struct {
	ram [0x2000]uint8
}

func NewVRAM() *VRAM {
	return &VRAM{}
}

func (v *VRAM) Read(addr uint16) uint8 {
	return v.ram[addr&0x1FFF]
}

func (v *VRAM) Write(addr uint16, val uint8) {
	v.ram[addr&0x1FFF] = val
}
