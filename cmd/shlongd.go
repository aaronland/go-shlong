package main

import (
       "flag"
       "fmt"	
       "github.com/thisisaaronland/go-shlong/engine"
       "github.com/thisisaaronland/go-shlong/http"       
       "log"
       gohttp "net/http"
       "net/url"
)

func main(){

     db_engine := flag.String("engine", "buntdb", "...")
     db_dsn := flag.String("dsn", ":memory:", "")

     http_host := flag.String("host", "localhost", "...")
     http_port := flag.Int("port", 8888, "...")
     
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
     
     handler, err := http.ShlongHandler(db, *root)

     if err != nil {
     	log.Fatal(err)
     }

     endpoint := fmt.Sprintf("%s:%d", *http_host, *http_port)

     err = gohttp.ListenAndServe(endpoint, handler)

     if err != nil {
     	log.Fatal(err)
     }
}
