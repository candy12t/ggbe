package cpu

import "testing"

func TestRegisters_Read16BitRegisters(t *testing.T) {
	r := &Registers{
		A:  0x12,
		B:  0x34,
		C:  0x56,
		D:  0x78,
		E:  0x9a,
		F:  0b_1010_0000,
		H:  0xbc,
		L:  0xde,
		SP: 0x1234,
		PC: 0x5678,
	}

	tests := []struct {
		name string
		fn   func() uint16
		want uint16
	}{
		{name: "AF", fn: r.AF, want: 0x12a0},
		{name: "BC", fn: r.BC, want: 0x3456},
		{name: "DE", fn: r.DE, want: 0x789a},
		{name: "HL", fn: r.HL, want: 0xbcde},
	}

	for _, tt := range tests {
		if got := tt.fn(); got != tt.want {
			t.Errorf("r.%v() returned %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

func TestRegisters_Write16BitRegisters(t *testing.T) {
	r := &Registers{
		A:  0x12,
		B:  0x34,
		C:  0x56,
		D:  0x78,
		E:  0x9a,
		F:  0b_1010_0000,
		H:  0xbc,
		L:  0xde,
		SP: 0x1234,
		PC: 0x5678,
	}

	tests := []struct {
		name    string
		writeFn func(val uint16)
		argVal  uint16
		readFn  func() uint16
		want    uint16
	}{
		{name: "WriteAF", writeFn: r.WriteAF, argVal: 0x1234, readFn: r.AF, want: 0x1230},
		{name: "WriteBC", writeFn: r.WriteBC, argVal: 0x5678, readFn: r.BC, want: 0x5678},
		{name: "WriteDE", writeFn: r.WriteDE, argVal: 0x9abc, readFn: r.DE, want: 0x9abc},
		{name: "WriteHL", writeFn: r.WriteHL, argVal: 0xdef0, readFn: r.HL, want: 0xdef0},
	}

	for _, tt := range tests {
		tt.writeFn(tt.argVal)
		if got := tt.readFn(); got != tt.want {
			t.Errorf("r.%v(%v) got %v, want %v\n", tt.name, tt.argVal, got, tt.want)
		}
	}
}

func TestRegisters_ReadFlagRegister(t *testing.T) {
	r := &Registers{
		A:  0x12,
		B:  0x34,
		C:  0x56,
		D:  0x78,
		E:  0x9a,
		F:  0b_1010_0000,
		H:  0xbc,
		L:  0xde,
		SP: 0x1234,
		PC: 0x5678,
	}

	tests := []struct {
		name string
		fn   func() bool
		want bool
	}{
		{name: "ZF", fn: r.ZF, want: true},
		{name: "NF", fn: r.NF, want: false},
		{name: "HF", fn: r.HF, want: true},
		{name: "CF", fn: r.CF, want: false},
	}

	for _, tt := range tests {
		if got := tt.fn(); got != tt.want {
			t.Errorf("r.%v() returned %v, want %v\n", tt.name, got, tt.want)
		}
	}
}

func TestRegisters_WriteFlagRegister(t *testing.T) {
	r := &Registers{
		A:  0x12,
		B:  0x34,
		C:  0x56,
		D:  0x78,
		E:  0x9a,
		F:  0b_1010_0000,
		H:  0xbc,
		L:  0xde,
		SP: 0x1234,
		PC: 0x5678,
	}

	tests := []struct {
		name    string
		writeFn func(bool)
		argVal  bool
		readFn  func() bool
		want    bool
	}{
		{name: "SetZF", writeFn: r.SetZF, argVal: true, readFn: r.ZF, want: true},
		{name: "SetNF", writeFn: r.SetNF, argVal: true, readFn: r.NF, want: true},
		{name: "SetHF", writeFn: r.SetHF, argVal: true, readFn: r.HF, want: true},
		{name: "SetCF", writeFn: r.SetCF, argVal: true, readFn: r.CF, want: true},
		{name: "SetZF", writeFn: r.SetZF, argVal: false, readFn: r.ZF, want: false},
		{name: "SetNF", writeFn: r.SetNF, argVal: false, readFn: r.NF, want: false},
		{name: "SetHF", writeFn: r.SetHF, argVal: false, readFn: r.HF, want: false},
		{name: "SetCF", writeFn: r.SetCF, argVal: false, readFn: r.CF, want: false},
	}

	for _, tt := range tests {
		tt.writeFn(tt.argVal)
		if got := tt.readFn(); got != tt.want {
			t.Errorf("r.%v(%v) got %v, want %v\n", tt.name, tt.argVal, got, tt.want)
		}
	}
}
