package analysis

import (
	"bytes"
)

func NewCeasarKeyDecryptor() *CeasarKeyDecryptor {
	return &CeasarKeyDecryptor{}
}

type CeasarKeyDecryptor struct{}

func (c *CeasarKeyDecryptor) GetKeyFor(source, encrypted []byte) (int, error) {
	sourceCharFreq := c.getRunesFrequency(bytes.Runes(source))
	encryptedCharFreq := c.getRunesFrequency(bytes.Runes(encrypted))
	return c.getDifferenceBetweenSingleMatchedPair(sourceCharFreq, encryptedCharFreq)
}

func (c *CeasarKeyDecryptor) getDifferenceBetweenSingleMatchedPair(
	sourceCharFreq map[rune]int,
	encryptedCharFreq map[rune]int,
) (int, error) {
	for sourceChar, sourceFreq := range sourceCharFreq {
		match := c.getCharsWithMatchedFrequency(sourceFreq, encryptedCharFreq)
		if isSinglePair := len(match) == 1; isSinglePair {
			return int(match[0] - sourceChar), nil
		}
	}

	return 0, ErrCannotDetectKey
}

func (c *CeasarKeyDecryptor) getCharsWithMatchedFrequency(
	target int,
	encryptedCharFreq map[rune]int,
) []rune {
	matched := make([]rune, 0, 1)
	for encryptedChar, encryptedFreq := range encryptedCharFreq {
		if encryptedFreq == target {
			matched = append(matched, encryptedChar)
		}
	}
	return matched
}

func (c *CeasarKeyDecryptor) getRunesFrequency(text []rune) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	return freq
}
