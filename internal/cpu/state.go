package cpu

type State struct {
	step  int
	val8  uint8
	val16 uint16
}

func NewState() *State {
	return &State{}
}

func (s *State) Step() int {
	return s.step
}

func (s *State) SetStep(v int) {
	s.step = v
}

func (s *State) Val8() uint8 {
	return s.val8
}

func (s *State) SetVal8(v uint8) {
	s.val8 = v
}

func (s *State) Val16() uint16 {
	return s.val16
}

func (s *State) SetVal16(v uint16) {
	s.val16 = v
}
