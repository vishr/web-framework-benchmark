package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/vishr/web-framework-benchmark/common"
	"github.com/zenazn/goji/web"
)

// Test 1: Path param
func pathParam(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(echo.ContentType, echo.TextPlainCharsetUTF8)
	io.WriteString(w, fmt.Sprintf("team: %s, user: %s", c.URLParams["id"], c.URLParams["user"]))
}

func main() {
	mux := web.New()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			mux.Get(r.Path, pathParam)
		case "POST":
			mux.Post(r.Path, pathParam)
		case "PUT":
			mux.Put(r.Path, pathParam)
		case "PATCH":
			mux.Patch(r.Path, pathParam)
		case "DELETE":
			mux.Delete(r.Path, pathParam)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
