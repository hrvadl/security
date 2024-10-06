package des

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func fromBits(b []Bit) (string, error) {
	var buider strings.Builder

	for symbol := range slices.Chunk(b, byteSize) {
		byteStr, err := bitSliceToStr(symbol)
		if err != nil {
			return "", errors.Join(ErrFailedToConvertFromBits, err)
		}

		bb, err := strconv.ParseUint(byteStr, 2, byteSize)
		if err != nil {
			return "", errors.Join(ErrFailedToConvertFromBits, err)
		}

		if err := buider.WriteByte(byte(bb)); err != nil {
			return "", errors.Join(ErrFailedToConvertFromBits, err)
		}
	}

	return buider.String(), nil
}

func toBits(text string) ([]Bit, error) {
	result := make([]Bit, 0, allowedTextLength)

	for _, b := range text {
		bitsStr := strings.Split(fmt.Sprintf("%.8b", b), "")
		bits, err := strSliceToBits(bitsStr)
		if err != nil {
			return nil, errors.Join(ErrFailedToConvertToBits, err)
		}
		result = append(result, bits...)
	}

	return result, nil
}

func bitSliceToStr(ss []Bit) (string, error) {
	var b strings.Builder
	for _, s := range ss {
		if _, err := b.WriteString(strconv.FormatUint(uint64(s), 10)); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func strSliceToBits(str []string) ([]Bit, error) {
	res := make([]Bit, 0, len(str))

	for _, b := range str {
		bb, err := strconv.ParseUint(b, 10, 8)
		if err != nil {
			return nil, err
		}
		res = append(res, Bit(bb))
	}

	return res, nil
}
