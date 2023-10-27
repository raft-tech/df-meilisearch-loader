package xml

import (
	"bytes"
	xj "github.com/basgys/goxml2json"
	"io"
)

func ToJSON(in io.Reader) (*bytes.Buffer, error) {
	return xj.Convert(in)
}
