# go-shlong

This a Go port of Mike Migurski's `shlong` URL shortener. No, really.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Usage

### Simple

```
package main

import (
       "flag"
       "github.com/thisisaaronland/go-shlong/engine"
       "log"
)

func main(){

     flag.Parse()
     
     db, err := engine.NewDBEngine("buntdb", "shlong.db")

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
```

Note that `short_url` is just a short _code_ and not a fully qualified domain. Syntactic hoohah to prepend a domain and an optional path to a short code is in the works.

## Engines

`go-shlong` has support for a variety of database engines. Currently it only has support for one database engine.

### buntdb

_Please write me_

## Interfaces

### Engine

```
type Engine interface {
	AddURL(long_url string) (string, error)
	GetLongURL(short_url string) (string, error)
	GetShortURL(long_url string) (string, error)
	Close()
}
```

Note that it's entirely possible the `Engine` interface will be updated to expect (and return) `url.URL` thingies which means the engine itself will need to know about root domains.

## Tools

_It's possible that `shlong` and `shlongd` will be merged in to single tool. Dunno..._

### shlong

_Please write me_

### shlongd

_Please write me_

```
./bin/shlongd -root http://freshandnew.com

curl 'localhost:8888?url=http://www.freshandnew.org/2017/01/2017-recapping-2016/'
http://freshandnew.com/6
```

## See also

* https://github.com/migurski/shlong
* https://github.com/tidwall/buntdb
