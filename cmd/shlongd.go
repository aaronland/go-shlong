package main

import (
       "flag"
       "github.com/thisisaaronland/go-shlong/engine"
       "github.com/thisisaaronland/go-shlong/http"       
       "log"
       gohttp "net/http"
       "net/url"
)

func main(){

     db_engine := flag.String("engine", "buntdb", "...")
     db_dsn := flag.String("dsn", ":memory:", "")

     root := flag.String("root", "", "")
     
     flag.Parse()

     db, err := engine.NewDBEngine(*db_engine, *db_dsn)

     if err != nil {
     	log.Fatal(err)
     }

     defer db.Close()

     _, err = url.Parse(*root)

     if err != nil {
     	log.Fatal(err)
     }
     
     handler := http.ShlongHandler(db, *root)

     endpoint := "localhost:8888"
     
     err = gohttp.ListenAndServe(endpoint, handler)

     if err != nil {
     	log.Fatal(err)
     }
}
