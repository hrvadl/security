package cipher

import "bytes"

const keyCorrelation = 1

func NewWithRearranmegent(key []int) *WithRearrangement {
	return &WithRearrangement{
		key: key,
	}
}

type WithRearrangement struct {
	key []int
}

func (wr *WithRearrangement) Encrypt(msg []byte) ([]byte, error) {
	msgRunes := bytes.Runes(msg)
	if wr.keyLengthDoesNotMatchMessage(msgRunes) {
		return nil, ErrInvalidKeyLength
	}

	encrypted := wr.rearrange(msgRunes)
	return []byte(string(encrypted)), nil
}

func (wr *WithRearrangement) rearrange(msg []rune) []rune {
	encryptedMsg := make([]rune, len(msg))
	for i, r := range msg {
		rearrangedIndex := wr.key[i] - keyCorrelation
		encryptedMsg[rearrangedIndex] = r
	}
	return encryptedMsg
}

func (wr *WithRearrangement) keyLengthDoesNotMatchMessage(msg []rune) bool {
	return len(msg) != len(wr.key)
}
