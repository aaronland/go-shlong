package database

import (
	"errors"
	"fmt"
	"github.com/thisisaaronland/go-shlong"
	"github.com/thisisaaronland/go-shlong/charset"
	"github.com/thisisaaronland/go-shlong/url"
	"github.com/tidwall/buntdb"
)

type BuntDB struct {
	shlong.Database
	db        *buntdb.DB
	charset   shlong.Charset
	max_tries int
}

func NewBuntDB(dsn string) (*BuntDB, error) {

	db, err := buntdb.Open(dsn)

	if err != nil {
		return nil, err
	}

	charset, err := charset.NewDefaultCharset()

	e := BuntDB{
		db:        db,
		charset:   charset,
		max_tries: 16,
	}

	return &e, nil
}

func (e *BuntDB) Close() {
	e.db.Close()
}

func (e *BuntDB) AddURL(long_url shlong.URL) (shlong.URL, error) {

	// log.Println("ADD", long_url)

	short, err := e.GetShortURL(long_url)

	if err != nil && err != buntdb.ErrNotFound {
		return nil, err
	}

	// log.Printf("SHORT (for long) '%s'\n", short)

	if short != nil {
		return short, nil
	}

	for i := 1; i < e.max_tries; i++ {

		id, err := e.charset.GenerateId(i)

		if err != nil {
			return nil, err
		}

		// log.Println("SHORT", i, id)

		short_url, err := url.NewShortURLFromString(id)

		if err != nil {
			return nil, err
		}

		long, err := e.GetLongURL(short_url)

		if err != nil && err != buntdb.ErrNotFound {
			return nil, err
		}

		// log.Printf("LONG (for short) '%s'\n", long)

		if long != nil {
			continue
		}

		// log.Println("SET LONG", long_url)

		err = e.set(fmt.Sprintf("long#%s", long_url.String()), short_url.String())

		if err != nil {
			return nil, err
		}

		// log.Println("SET SHORT", short_url)

		err = e.set(fmt.Sprintf("short#%s", short_url.String()), long_url.String())

		if err != nil {
			return nil, err
		}

		return short_url, nil
	}

	return nil, errors.New("Exceeded max tries")
}

func (e *BuntDB) GetShortURL(lu shlong.URL) (shlong.URL, error) {

	key := fmt.Sprintf("long#%s", lu.String())
	su, err := e.get(key)

	if err != nil {
		return nil, err
	}

	return url.NewShortURLFromString(su)
}

func (e *BuntDB) GetLongURL(su shlong.URL) (shlong.URL, error) {

	key := fmt.Sprintf("short#%s", su.String())
	lu, err := e.get(key)

	if err != nil {
		return nil, err
	}

	return url.NewLongURLFromString(lu)
}

func (e *BuntDB) set(key string, value string) error {

	err := e.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})

	return err
}

func (e *BuntDB) get(key string) (string, error) {

	var value string

	err := e.db.View(func(tx *buntdb.Tx) error {

		val, err := tx.Get(key)

		if err != nil {
			return err
		}

		value = val
		return nil
	})

	if err != nil {
		return "", err
	}

	return value, nil
}
