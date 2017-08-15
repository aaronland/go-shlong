prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/thisisaaronland/go-shlong; then rm -rf src/github.com/thisisaaronland/go-shlong; fi
	mkdir -p src/github.com/thisisaaronland/go-shlong
	mkdir -p src/github.com/thisisaaronland/go-shlong/charset
	mkdir -p src/github.com/thisisaaronland/go-shlong/database
	mkdir -p src/github.com/thisisaaronland/go-shlong/http
	mkdir -p src/github.com/thisisaaronland/go-shlong/utils
	cp *.go src/github.com/thisisaaronland/go-shlong/
	cp charset/*.go src/github.com/thisisaaronland/go-shlong/charset/
	cp database/*.go src/github.com/thisisaaronland/go-shlong/database/
	cp http/*.go src/github.com/thisisaaronland/go-shlong/http/
	cp utils/*.go src/github.com/thisisaaronland/go-shlong/utils/
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/tidwall/buntdb"
	@GOPATH=$(shell pwd) go get "github.com/lib/pq"
	# @GOPATH=$(shell pwd) go get "github.com/whosonfirst/go-ucd"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt charset/*.go
	go fmt database/*.go
	go fmt http/*.go
	go fmt utils/*.go

bin:    self
	if test ! -d bin; then mkdir bin; fi
	@GOPATH=$(shell pwd) go build -o bin/shlong cmd/shlong.go
	@GOPATH=$(shell pwd) go build -o bin/shlongd cmd/shlongd.go
