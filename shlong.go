package shlong

import (
	"math/rand"
)

type Engine interface {
	AddURL(long_url string) (string, error)
	GetLongURL(short_url string) (string, error)
	GetShortURL(long_url string) (string, error)
	Close()
}

func GenerateId(length int) (string, error) {

	chars := "qwrtypsdfghjklzxcvbnm0123456789"
	len_chars := len(chars)

	id := ""

	for len(id) < length {

		i := rand.Intn(len_chars)
		id = id + string(chars[i])
	}

	return id, nil
}
