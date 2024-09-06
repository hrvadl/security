package contentsign

const header = "\n---SIGNATURE---\n"

func NewAppender() *Appender {
	return &Appender{}
}

type Appender struct{}

func (a *Appender) AppendSign(data []byte, sign []byte) []byte {
	signed := make([]byte, 0, len(data)+len(header)+len(sign))
	copy(signed, data)
	withHeader := append(data, []byte(header)...)
	return append(withHeader, sign...)
}
