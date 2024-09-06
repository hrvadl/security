package contentsign

import "bytes"

func NewExtracter() *Extracter {
	return &Extracter{}
}

type Extracter struct{}

func (e *Extracter) ExtractSign(data []byte) ([]byte, []byte, error) {
	splits := bytes.Split(data, []byte(header))
	if len(splits) < 2 {
		return nil, nil, ErrNotFound
	}

	return splits[0], bytes.Trim(splits[1], "\n"), nil
}
