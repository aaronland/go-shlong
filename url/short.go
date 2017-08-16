package url

import (
       "fmt"
       "github.com/thisisaaronland/go-shlong"
)

type ShortURL struct {
	shlong.URL
	host string
	path string
}

func NewShortURLFromString(path string) (*ShortURL, error) {

	su := ShortURL{
		host: "shlong",
		path: path,
	}

	return &su, nil
}

func (su *ShortURL) Hostname() string {
	return su.host
}

func (su *ShortURL) Path() string {
	return su.path
}

func (su *ShortURL) String() string {
	return fmt.Sprintf("urn:%s:%s", su.Hostname(), su.Path())
}
