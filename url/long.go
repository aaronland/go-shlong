package url

type LongURL struct {
	URL
	u gourl.URL
}

func NewLongURL(raw string) (LongURL, error) {

	u, err := gourl.Parse(raw)

	if err != nil {
		return nil, err
	}

	lu := LongURL{
		u: u,
	}

	return lu, nil
}

func (lu LongURL) Hostname() string {
	return lu.u.Host
}

func (lu LongURL) Path() string {
	return lu.u.Path
}

func (lu LongURL) String() string {
	return lu.u.String()
}
