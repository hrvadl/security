package ceasar

import "math"

const (
	latinUppercaseFirst = 65
	latinUppercaseLast  = 90
)

const (
	latinLowercaseFirst = 97
	latinLowercaseLast  = 122
)

const whitespaceCodepoint = 32

func newCyclicShifterForLowercase(shift int) *cyclicShifter {
	return newCyclicShifter(
		latinLowercaseFirst,
		latinLowercaseLast,
		shift,
	)
}

func newCyclicShifterForUppercase(shift int) *cyclicShifter {
	return newCyclicShifter(
		latinUppercaseFirst,
		latinUppercaseLast,
		shift,
	)
}

func newCyclicShifter(from, to rune, shift int) *cyclicShifter {
	return &cyclicShifter{
		from:  from,
		to:    to,
		shift: shift,
	}
}

type cyclicShifter struct {
	from  rune
	to    rune
	shift int
}

func (cs *cyclicShifter) Forward(target rune) rune {
	if cs.isWhitespace(target) {
		return target
	}

	if cs.biggerThanUpperLimit(target) {
		return cs.from + rune(math.Abs(float64(cs.to)-float64(target+rune(cs.shift)))) - 1
	}

	return target + rune(cs.shift)
}

func (cs *cyclicShifter) Backward(target rune) rune {
	if cs.isWhitespace(target) {
		return target
	}

	if cs.smallerThanLowerLimit(target) {
		return cs.to - (cs.from - (target - rune(cs.shift))) + 1
	}

	return target - rune(cs.shift)
}

func (cs *cyclicShifter) biggerThanUpperLimit(target rune) bool {
	return target+rune(cs.shift) > cs.to
}

func (cs *cyclicShifter) smallerThanLowerLimit(target rune) bool {
	return target-rune(cs.shift) < cs.from
}

func (cs *cyclicShifter) isWhitespace(r rune) bool {
	return r == whitespaceCodepoint
}
