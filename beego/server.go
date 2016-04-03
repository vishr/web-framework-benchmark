package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test: Path param
func pathParam(c *context.Context) {
	c.WriteString(fmt.Sprintf("team: %s, user: %s", c.Input.Param(":id"), c.Input.Param(":user")))
}

func main() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	app := beego.NewControllerRegister()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			app.Get(r.Path, pathParam)
		case "POST":
			app.Post(r.Path, pathParam)
		case "PUT":
			app.Put(r.Path, pathParam)
		case "PATCH":
			app.Patch(r.Path, pathParam)
		case "DELETE":
			app.Delete(r.Path, pathParam)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(http.ListenAndServe(":8080", app))
}
