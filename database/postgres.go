package database

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/thisisaaronland/go-shlong"
	"github.com/thisisaaronland/go-shlong/charset"
	"log"
)

/*

CREATE TABLE urls (short_url VARCHAR(255) PRIMARY KEY, long_url TEXT);
CREATE INDEX by_long_url ON urls (long_url)

./bin/shlong -engine postgres -dsn "dbname=urls" http://www.freshandnew.org/2017/01/2017-recapping-2016/

*/

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

	short, err := p.GetShortURL(long_url)

	if err != nil {
		return "", err
	}

	if short != "" {
		return short, nil
	}

	for i := 1; i < p.max_tries; i++ {

		id, err := p.charset.GenerateId(i)

		if err != nil {
			return "", err
		}

		short_url := id

		long, err := p.GetLongURL(short_url)

		if err != nil {
			return "", err
		}

		if long != "" {
			continue
		}

		sql := "INSERT INTO urls (long_url, short_url) VALUES ($1, $2) ON CONFLICT(short_url) DO UPDATE SET long_url=$3"
		_, err = p.db.Exec(sql, long_url, short_url, long_url)

		if err != nil {
			return "", err
		}

		return short_url, nil
	}

	return "", errors.New("Exceeded max tries")
}

func (p *PostgresDB) GetShortURL(long_url string) (string, error) {

	sql := "SELECT short_url FROM urls WHERE long_url = $1"
	row := p.db.QueryRow(sql, long_url)

	var short_url string
	err := row.Scan(&short_url)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil
		}

		return "", err
	}

	return short_url, nil

}

func (p *PostgresDB) GetLongURL(short_url string) (string, error) {

	sql := "SELECT long_url FROM urls WHERE short_url = $1"
	row := p.db.QueryRow(sql, short_url)

	var long_url string
	err := row.Scan(&long_url)

	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return "", nil
		}
		return "", err
	}

	return long_url, nil
}
