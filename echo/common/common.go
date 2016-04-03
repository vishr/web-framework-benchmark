package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Path param
func pathParam(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}

func RegisterDynamicRoutes(e *echo.Echo) {
	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			e.Get(r.Path, pathParam)
		case "POST":
			e.Post(r.Path, pathParam)
		case "PUT":
			e.Put(r.Path, pathParam)
		case "PATCH":
			e.Patch(r.Path, pathParam)
		case "DELETE":
			e.Delete(r.Path, pathParam)
		default:
			panic("method not supported")
		}
	}
}
