package analysis

import (
	"bytes"
)

func NewCaesarKeyDecryptor() *CaesarKeyDecryptor {
	return &CaesarKeyDecryptor{}
}

// CaesarKeyDecryptor tries to guess the key
// for a given input and encrypted with Caesar
// method text.
type CaesarKeyDecryptor struct{}

// GetKeyFor method tries to guess key for the decrypted/encrypted text pair.
// Under the hood it will count chars for decrypted/encrypted text and then
// it will try to calculate the caesar shift if possible.
func (c *CaesarKeyDecryptor) GetKeyFor(source, encrypted []byte) (int, error) {
	sourceCharFreq := c.getRunesFrequency(bytes.Runes(source))
	encryptedCharFreq := c.getRunesFrequency(bytes.Runes(encrypted))
	return c.getDifferenceBetweenSingleMatchedPair(sourceCharFreq, encryptedCharFreq)
}

func (c *CaesarKeyDecryptor) getDifferenceBetweenSingleMatchedPair(
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

func (c *CaesarKeyDecryptor) getCharsWithMatchedFrequency(
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

func (c *CaesarKeyDecryptor) getRunesFrequency(text []rune) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	return freq
}
