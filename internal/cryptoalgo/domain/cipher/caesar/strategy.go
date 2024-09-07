package caesar

func newShiftStrategy(shift int) *shiftStrategy {
	return &shiftStrategy{
		shift: shift,
	}
}

// shiftStrategy struct is responsible for determining
// whether given char is uppercase or lowercase.
// Then it will create cyclicShifter with the predefined limits
// and shift value.
type shiftStrategy struct {
	shifter *cyclicShifter
	shift   int
}

// Forward function shifts value in forward direction.
// Under the hood it sets active strategy for the
// given rune and delegates actual work for the strategy.
func (s *shiftStrategy) Forward(r rune) rune {
	s.setStrategyBasedOnRune(r)
	return s.shifter.Forward(r)
}

// Backward function shifts value in backward direction.
// Under the hood it sets active strategy for the
// given rune and delegates actual work for the strategy.
func (s *shiftStrategy) Backward(r rune) rune {
	s.setStrategyBasedOnRune(r)
	return s.shifter.Backward(r)
}

// TODO: optimize this. cuz it's allocating new obj on each
// function call.
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
