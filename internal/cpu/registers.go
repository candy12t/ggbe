package cpu

type Registers struct {
	A          uint8  // accumulator register
	B, C, D, E uint8  // general register
	F          uint8  // flag register. Lower 4bit always 0
	H, L       uint8  // general register
	SP         uint16 // stack pointer
	PC         uint16 // program counter
}

func newRegisters() *Registers {
	return &Registers{}
}

func (r *Registers) AF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F)
}

func (r *Registers) BC() uint16 {
	return uint16(r.B)<<8 | uint16(r.C)
}

func (r *Registers) DE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

func (r *Registers) HL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

func (r *Registers) WriteAF(val uint16) {
	r.A = uint8(val >> 8)
	r.F = uint8(val & 0xF0)
}

func (r *Registers) WriteBC(val uint16) {
	r.B = uint8(val >> 8)
	r.C = uint8(val)
}

func (r *Registers) WriteDE(val uint16) {
	r.D = uint8(val >> 8)
	r.E = uint8(val)
}

func (r *Registers) WriteHL(val uint16) {
	r.H = uint8(val >> 8)
	r.L = uint8(val)
}

func (r *Registers) ZF() bool {
	return r.F&0b_1000_0000 > 0
}

func (r *Registers) NF() bool {
	return r.F&0b_0100_0000 > 0
}

func (r *Registers) HF() bool {
	return r.F&0b_0010_0000 > 0
}

func (r *Registers) CF() bool {
	return r.F&0b_0001_0000 > 0
}

func (r *Registers) SetZF(zf bool) {
	if zf {
		r.F |= 0b_1000_0000
	} else {
		r.F &= 0b_0111_1111
	}
}

func (r *Registers) SetNF(nf bool) {
	if nf {
		r.F |= 0b_0100_0000
	} else {
		r.F &= 0b_1011_1111
	}
}

func (r *Registers) SetHF(hf bool) {
	if hf {
		r.F |= 0b_0010_0000
	} else {
		r.F &= 0b_1101_1111
	}
}

func (r *Registers) SetCF(cf bool) {
	if cf {
		r.F |= 0b_0001_0000
	} else {
		r.F &= 0b_1110_1111
	}
}
