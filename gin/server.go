package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Dynamic Route
func dynamicRoute(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			g.GET(r.Path, dynamicRoute)
		case "POST":
			g.POST(r.Path, dynamicRoute)
		case "PUT":
			g.PUT(r.Path, dynamicRoute)
		case "PATCH":
			g.PATCH(r.Path, dynamicRoute)
		case "DELETE":
			g.DELETE(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(g.Run(":8080"))
}
