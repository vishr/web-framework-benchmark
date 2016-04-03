package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/vishr/web-framework-benchmark/echo/common"
)

func main() {
	e := echo.New()
	common.RegisterDynamicRoutes(e)
	e.Run(fasthttp.New(":8080"))
}
