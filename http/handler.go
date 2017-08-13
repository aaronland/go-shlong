package http

import (
	"fmt"
	"github.com/thisisaaronland/go-shlong"
	_ "log"
	gohttp "net/http"
	"net/url"
)

func ShlongHandler(db shlong.Engine, root string) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()
		raw := query.Get("url")

		u, err := url.ParseRequestURI(raw)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		short, err := db.AddURL(u.String())

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		root_url, err := url.Parse(root)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		root_url.Path = short

		rsp.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(rsp, root_url.String())
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
