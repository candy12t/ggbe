package cpu

type Cond int

const (
	CondNZ Cond = iota + 1
	CondZ
	CondNC
	CondC
)
