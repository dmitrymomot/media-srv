package resizer

import (
	"bytes"
	"io"
	"io/ioutil"
)

// MockResize ...
func MockResize(err error) func(f io.ReadCloser, w, h int) (io.ReadSeeker, error) {
	return func(f io.ReadCloser, w, h int) (io.ReadSeeker, error) {
		if err != nil {
			return nil, err
		}
		b, _ := ioutil.ReadAll(f)
		return bytes.NewReader(b), nil
	}
}
