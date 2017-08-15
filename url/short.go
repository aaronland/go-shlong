package url

type ShortURL struct {
	URL
	host string
	path string
}

func NewShortURL(path string) (ShortURL, error) {

	su := ShortURL{
		host: "shlong",
		path: path,
	}

	return su
}

func (su ShortURL) Hostname() string {
	return su.host
}

func (su ShortURL) Path() string {
	return su.path
}

func (su ShortURL) String() string {
	return fmt.Sprintf("urn:%s:%s", su.Host(), su.Path())
}
