package shlong

import (

)

type Engine interface {
	AddURL(long_url string) (string, error)
	GetLongURL(short_url string) (string, error)
	GetShortURL(long_url string) (string, error)
	Close()
}

type Charset interface {
     Characters() string
     GenerateId(length int) (string, error)
}
