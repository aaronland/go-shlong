package shlong

type Charset interface {
	Characters() string
	GenerateId(length int) (string, error)
}

type Database interface {
	AddURL(long_url string) (string, error)
	GetLongURL(short_url string) (string, error)
	GetShortURL(long_url string) (string, error)
	Close()
}

type URL interface {
	Hostname() string
	Path() string
	String() string
}
