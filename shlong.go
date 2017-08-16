package shlong

type Charset interface {
	Characters() string
	GenerateId(length int) (string, error)
}

type Database interface {
	AddURL(long_url URL) (URL, error)
	GetLongURL(short_url URL) (URL, error)
	GetShortURL(long_url URL) (URL, error)
	Close()
}

type URL interface {
	Hostname() string
	Path() string
	String() string
}
