package main

import (
	"restaurant-management/pkg/containers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Serve(e)
}

