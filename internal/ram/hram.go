package ram

// HRAM is High RAM
// the size is 128 B
type HRAM struct {
	ram [0x80]uint8
}

func NewHRAM() *HRAM {
	return &HRAM{}
}

func (h *HRAM) Read(addr uint16) uint8 {
	return h.ram[addr&0x7F]
}

func (h *HRAM) Write(addr uint16, val uint8) {
	h.ram[addr&0x7F] = val
}
