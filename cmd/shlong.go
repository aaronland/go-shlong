package main

import (
       "flag"
       "github.com/thisisaaronland/go-shlong/database"
       "github.com/thisisaaronland/go-shlong/url"       
       "log"
)

func main(){

     db_engine := flag.String("database", "buntdb", "...")
     db_dsn := flag.String("dsn", ":memory:", "")
     
     flag.Parse()
     
     db, err := database.NewDatabase(*db_engine, *db_dsn)

     if err != nil {
     	log.Fatal(err)
     }

     defer db.Close()
     
     for _, str_url := range flag.Args() {

     	long_url, err := url.NewLongURLFromString(str_url)

	if err != nil {
	   log.Fatal(err)
	}
	
     	short_url, err := db.AddURL(long_url)

	if err != nil {
	   log.Fatal(err)
	}

	log.Println(long_url, short_url)
     }
}
