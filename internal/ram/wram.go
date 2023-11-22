package ram

// WRAM is Work RAM
// the size is 8 KiB
type WRAM struct {
	ram [0x2000]uint8
}

func NewWRAM() *WRAM {
	return &WRAM{}
}

func (w *WRAM) Read(addr uint16) uint8 {
	return w.ram[addr&0x1FFF]
}

func (w *WRAM) Write(addr uint16, val uint8) {
	w.ram[addr&0x1FFF] = val
}
