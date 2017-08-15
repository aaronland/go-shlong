package database

import (
	"database/sql"
	"errors"		
	_ "github.com/lib/pq"
	"github.com/thisisaaronland/go-shlong"
	"github.com/thisisaaronland/go-shlong/charset"	
)

type PostgresDB struct {
	shlong.Database
	db        *sql.DB
	charset   shlong.Charset
	max_tries int
}

func NewPostgresDB(dsn string) (*PostgresDB, error) {

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	charset, err := charset.NewDefaultCharset()

	e := PostgresDB{
		db:        db,
		charset:   charset,
		max_tries: 16,
	}

	return &e, nil
}

func (p *PostgresDB) Close() {
	p.db.Close()
}

func (p *PostgresDB) AddURL(long_url string) (string, error) {

     return "", errors.New("Please write me")
}

func (p *PostgresDB) GetShortURL(long_url string) (string, error) {

     return "", errors.New("Please write me")
}

func (p *PostgresDB) GetLongURL(short_url string) (string, error) {

     return "", errors.New("Please write me")
}
