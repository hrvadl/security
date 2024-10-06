package des

import (
	"fmt"
	"strings"
)

func stringToBits(text string) []Bit {
	result := make([]Bit, 0, 64)
	for _, b := range text {
		bits := fmt.Sprintf("%.8b", b)
		result = append(result, strings.Split(bits, "")...)
	}
	return result
}
