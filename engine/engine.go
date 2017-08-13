package engine

import (
	"errors"
	"github.com/thisisaaronland/go-shlong"
)

func NewDBEngine(name string, dsn string) (shlong.Engine, error) {

	if name == "buntdb" {
		return NewBuntDBEngine(dsn)
	} else {
		return nil, errors.New("Invalid database engine")
	}
}
