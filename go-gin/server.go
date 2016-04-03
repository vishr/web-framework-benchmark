package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Path param
func pathParam(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			g.GET(r.Path, pathParam)
		case "POST":
			g.POST(r.Path, pathParam)
		case "PUT":
			g.PUT(r.Path, pathParam)
		case "PATCH":
			g.PATCH(r.Path, pathParam)
		case "DELETE":
			g.DELETE(r.Path, pathParam)
		default:
			panic("method not supported")
		}
	}

	log.Fatal(g.Run(":8080"))
}
