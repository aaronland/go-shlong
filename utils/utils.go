package utils

import (
       "github.com/thisisaaronland/go-shlong"
       	"math/rand"       
)

func RandomStringFromCharset(cs shlong.Charset, length int) (string, error) {

	chars := cs.Characters()
	len_chars := len(chars)

	id := ""

	for len(id) < length {

		i := rand.Intn(len_chars)
		id = id + string(chars[i])
	}

	return id, nil
}

