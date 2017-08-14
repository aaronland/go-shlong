package charset

import (
	"github.com/thisisaaronland/go-shlong"
)

func NewCharset(name string) (shlong.Charset, error) {

	return NewDefaultCharset()
}
