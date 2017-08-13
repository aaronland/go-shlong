package engine

import (
	"errors"
	"fmt"
	"github.com/thisisaaronland/go-shlong"
	"github.com/tidwall/buntdb"
)

type BuntDBEngine struct {
	shlong.Engine
	db        *buntdb.DB
	max_tries int
}

func NewBuntDBEngine(dsn string) (*BuntDBEngine, error) {

	db, err := buntdb.Open(dsn)

	if err != nil {
		return nil, err
	}

	e := BuntDBEngine{
		db:        db,
		max_tries: 16,
	}

	return &e, nil
}

func (e *BuntDBEngine) Close() {
	e.db.Close()
}

func (e *BuntDBEngine) AddURL(long_url string) (string, error) {

	// log.Println("ADD", long_url)

	short, err := e.GetShortURL(long_url)

	if err != nil && err != buntdb.ErrNotFound {
		return "", err
	}

	// log.Printf("SHORT (for long) '%s'\n", short)

	if short != "" {
		return short, nil
	}

	for i := 1; i < e.max_tries; i++ {

		id, err := shlong.GenerateId(i)

		if err != nil {
			return "", err
		}

		// log.Println("SHORT", i, id)

		short_url := id

		long, err := e.GetLongURL(short_url)

		if err != nil && err != buntdb.ErrNotFound {
			return "", err
		}

		// log.Printf("LONG (for short) '%s'\n", long)

		if long != "" {
			continue
		}

		// log.Println("SET LONG", long_url)

		err = e.set(fmt.Sprintf("long#%s", long_url), short_url)

		if err != nil {
			return "", err
		}

		// log.Println("SET SHORT", short_url)

		err = e.set(fmt.Sprintf("short#%s", short_url), long_url)

		if err != nil {
			return "", err
		}

		return short_url, nil
	}

	return "", errors.New("Exceeded max tries")
}

func (e *BuntDBEngine) GetShortURL(long_url string) (string, error) {

	key := fmt.Sprintf("long#%s", long_url)
	return e.get(key)
}

func (e *BuntDBEngine) GetLongURL(short_url string) (string, error) {

	key := fmt.Sprintf("short#%s", short_url)
	return e.get(key)
}

func (e *BuntDBEngine) set(key string, value string) error {

	err := e.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})

	return err
}

func (e *BuntDBEngine) get(key string) (string, error) {

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
