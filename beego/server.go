package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test: Dynamic Route
func dynamicRoute(c *context.Context) {
	c.WriteString(fmt.Sprintf("team: %s, user: %s", c.Input.Param(":id"), c.Input.Param(":user")))
}

func main() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	app := beego.NewControllerRegister()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			app.Get(r.Path, dynamicRoute)
		case "POST":
			app.Post(r.Path, dynamicRoute)
		case "PUT":
			app.Put(r.Path, dynamicRoute)
		case "PATCH":
			app.Patch(r.Path, dynamicRoute)
		case "DELETE":
			app.Delete(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(http.ListenAndServe(":8080", app))
}
