package http

import (
	"fmt"
	"github.com/thisisaaronland/go-shlong"
	"github.com/thisisaaronland/go-shlong/url"
	_ "log"
	gohttp "net/http"
	gourl "net/url"
)

func LongToShortHandler(db shlong.Database, root string) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()
		raw := query.Get("url")

		long_url, err := url.NewLongURLFromString(raw)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		short_url, err := db.AddURL(long_url)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		root_url, err := gourl.Parse(root)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		root_url.Path = short_url.Path()

		rsp.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(rsp, root_url.String())
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}

func ShortToLongHandler(db shlong.Database, root string) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		path := req.URL.Path

		short_url, err := url.NewShortURLFromString(path)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		long_url, err := db.GetLongURL(short_url)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		if long_url == nil {
			gohttp.Error(rsp, "", gohttp.StatusNotFound)
			return
		}

		gohttp.Redirect(rsp, req, long_url.String(), gohttp.StatusFound)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
