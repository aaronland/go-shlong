package database

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/thisisaaronland/go-shlong"
	"github.com/thisisaaronland/go-shlong/charset"
	"github.com/thisisaaronland/go-shlong/url"	
	_ "log"
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

func (p *PostgresDB) AddURL(long_url shlong.URL) (shlong.URL, error) {

	short, err := p.GetShortURL(long_url)

	if err != nil {
		return nil, err
	}

	if short != nil {
		return short, nil
	}

	for i := 1; i < p.max_tries; i++ {

		id, err := p.charset.GenerateId(i)

		if err != nil {
			return nil, err
		}

		short_url, err := url.NewShortURLFromString(id)

		if err != nil {
			return nil, err
		}

		long, err := p.GetLongURL(short_url)

		if err != nil {
			return nil, err
		}

		if long != nil {
			continue
		}

		lu := long_url.String()
		su := short_url.String()
		
		sql := "INSERT INTO urls (long_url, short_url) VALUES ($1, $2) ON CONFLICT(short_url) DO UPDATE SET long_url=$3"
		_, err = p.db.Exec(sql, lu, su, lu)

		if err != nil {
			return nil, err
		}

		return short_url, nil
	}

	return nil, errors.New("Exceeded max tries")
}

func (p *PostgresDB) GetShortURL(long_url shlong.URL) (shlong.URL, error) {

     	lu := long_url.String()
	
	sql := "SELECT short_url FROM urls WHERE long_url = $1"
	row := p.db.QueryRow(sql, lu)

	var short_url string
	err := row.Scan(&short_url)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	return url.NewShortURLFromString(short_url)

}

func (p *PostgresDB) GetLongURL(short_url shlong.URL) (shlong.URL, error) {

	sql := "SELECT long_url FROM urls WHERE short_url = $1"
	row := p.db.QueryRow(sql, short_url)

	var long_url string
	err := row.Scan(&long_url)

	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		
		return nil, err
	}

	return url.NewLongURLFromString(long_url)
}
