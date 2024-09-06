package contentsign

const header = "\n---SIGNATURE---\n"

func NewAppender() *Appender {
	return &Appender{}
}

// Appender is a struct responsible for appending
// signature as a metadata header to the end
// of the file. It's meaning exactly opposite from the
// Extracter struct.
type Appender struct{}

// AppendSign sign takes a data as an argument.
// Then it appends given signature with the header
// to the end of the file.
func (a *Appender) AppendSign(data []byte, sign []byte) []byte {
	signed := make([]byte, 0, len(data)+len(header)+len(sign))
	copy(signed, data)
	withHeader := append(data, []byte(header)...)
	return append(withHeader, sign...)
}
