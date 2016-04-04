package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vishr/web-framework-benchmark/common"
)

// Test 1: Dynamic Route
func dynamicRoute(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}

func RegisterDynamicRoutes(e *echo.Echo) {
	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			e.Get(r.Path, dynamicRoute)
		case "POST":
			e.Post(r.Path, dynamicRoute)
		case "PUT":
			e.Put(r.Path, dynamicRoute)
		case "PATCH":
			e.Patch(r.Path, dynamicRoute)
		case "DELETE":
			e.Delete(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}
}
