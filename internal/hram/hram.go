package hram

// HRAM is High RAM
// HRAM is 128 B
type HRAM struct {
	ram [0x80]uint8
}

func New() *HRAM {
	return &HRAM{}
}

func (h *HRAM) Read(addr uint16) uint8 {
	return h.ram[addr&0x7f]
}

func (h *HRAM) Write(addr uint16, val uint8) {
	h.ram[addr&0x7f] = val
}
