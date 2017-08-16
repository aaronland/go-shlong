package url

import (
       "github.com/thisisaaronland/go-shlong"
       gourl "net/url"
)

type LongURL struct {
	shlong.URL
	u *gourl.URL
}

func NewLongURLFromString(raw string) (*LongURL, error) {

	u, err := gourl.Parse(raw)

	if err != nil {
		return nil, err
	}

	lu := LongURL{
		u: u,
	}

	return &lu, nil
}

func (lu *LongURL) Hostname() string {
	return lu.u.Host
}

func (lu *LongURL) Path() string {
	return lu.u.Path
}

func (lu *LongURL) String() string {
	return lu.u.String()
}
