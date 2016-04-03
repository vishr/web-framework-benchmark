package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Path param
func pathParam(params martini.Params) string {
	return fmt.Sprintf("team: %s, user: %s", params["id"], params["user"])
}

func main() {
	m := martini.New()
	router := martini.NewRouter()
	m.Action(router.Handle)
	martini.Env = martini.Prod

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			router.Get(r.Path, pathParam)
		case "POST":
			router.Post(r.Path, pathParam)
		case "PUT":
			router.Put(r.Path, pathParam)
		case "PATCH":
			router.Patch(r.Path, pathParam)
		case "DELETE":
			router.Delete(r.Path, pathParam)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(http.ListenAndServe(":8080", m))
}
