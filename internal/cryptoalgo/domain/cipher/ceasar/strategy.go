package ceasar

func NewShiftStrategy(shift int) *shiftStrategy {
	return &shiftStrategy{
		shift: shift,
	}
}

type Shifter interface {
	Forward(target rune) rune
	Backward(target rune) rune
}

type shiftStrategy struct {
	shifter Shifter
	shift   int
}

func (s *shiftStrategy) Forward(r rune) rune {
	s.setStrategyBasedOnRune(r)
	return s.shifter.Forward(r)
}

func (s *shiftStrategy) Backward(r rune) rune {
	s.setStrategyBasedOnRune(r)
	return s.shifter.Backward(r)
}

func (s *shiftStrategy) setStrategyBasedOnRune(r rune) {
	switch {
	case s.isUpperCaseLatin(r):
		s.shifter = newCyclicShifterForUppercase(s.shift)
	case s.isLowerCaseLatin(r):
		s.shifter = newCyclicShifterForLowercase(s.shift)
	}
}

func (s *shiftStrategy) isUpperCaseLatin(r rune) bool {
	return r >= latinUppercaseFirst && r <= latinUppercaseLast
}

func (s *shiftStrategy) isLowerCaseLatin(r rune) bool {
	return r >= latinLowercaseFirst && r <= latinLowercaseLast
}
