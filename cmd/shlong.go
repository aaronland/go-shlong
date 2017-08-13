package main

import (
       "flag"
       "github.com/thisisaaronland/go-shlong/engine"
       "log"
)

func main(){

     db_engine := flag.String("engine", "buntdb", "...")
     db_dsn := flag.String("dsn", ":memory:", "")
     
     flag.Parse()

     db, err := engine.NewDBEngine(*db_engine, *db_dsn)

     if err != nil {
     	log.Fatal(err)
     }

     defer db.Close()
     
     for _, long_url := range flag.Args() {

     	short_url, err := db.AddURL(long_url)

	if err != nil {
	   log.Fatal(err)
	}

	log.Println(long_url, short_url)
     }
}
