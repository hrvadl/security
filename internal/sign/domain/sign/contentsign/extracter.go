package contentsign

import "bytes"

const signedDocParts = 2

func NewExtracter() *Extracter {
	return &Extracter{}
}

// Extracter is a small struct responsible
// for extracting signature from data in the
// signed file. It's meaning exactly opposite from
// Appender struct.
type Extracter struct{}

// ExtractSign sign takes a signed data with metadata
// as an argument. Then it tries to split data by signature
// header and returns data with signature and error (if any).
func (e *Extracter) ExtractSign(data []byte) ([]byte, []byte, error) {
	splits := bytes.Split(data, []byte(header))
	if len(splits) < signedDocParts {
		return nil, nil, ErrNotFound
	}

	return splits[0], bytes.Trim(splits[len(splits)-1], "\n"), nil
}
