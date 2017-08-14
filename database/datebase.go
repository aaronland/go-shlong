package database

import (
	"errors"
	"github.com/thisisaaronland/go-shlong"
)

func NewDatabase(name string, dsn string) (shlong.Database, error) {

	if name == "buntdb" {
		return NewBuntDB(dsn)
	} else {
		return nil, errors.New("Invalid database engine")
	}
}
