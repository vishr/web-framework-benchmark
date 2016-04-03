package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/vishr/web-framework-benchmark/echo/common"
)

func main() {
	e := echo.New()
	common.RegisterDynamicRoutes(e)
	e.Run(standard.New(":8080"))
}
