package main

import (
       "flag"
       "fmt"	
       "github.com/thisisaaronland/go-shlong/database"
       "github.com/thisisaaronland/go-shlong/http"
       "log"
       gohttp "net/http"
       "net/url"
)

func main(){

     db_engine := flag.String("database", "buntdb", "...")
     db_dsn := flag.String("dsn", ":memory:", "")

     http_host := flag.String("host", "localhost", "...")
     http_port := flag.Int("port", 8888, "...")
     
     root := flag.String("root", "", "")
     
     flag.Parse()

     db, err := database.NewDatabase(*db_engine, *db_dsn)

     if err != nil {
     	log.Fatal(err)
     }

     defer db.Close()

     _, err = url.Parse(*root)

     if err != nil {
     	log.Fatal(err)
     }
     
     handler, err := http.ShortToLongHandler(db, *root)

     if err != nil {
     	log.Fatal(err)
     }

     endpoint := fmt.Sprintf("%s:%d", *http_host, *http_port)

     err = gohttp.ListenAndServe(endpoint, handler)

     if err != nil {
     	log.Fatal(err)
     }
}
