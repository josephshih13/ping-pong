package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var cnt = 0

func getCnt(c echo.Context) error {
	
	ret := fmt.Sprintf("Pong %d\n", cnt)
	cnt = cnt + 1
	return c.String(http.StatusOK, ret)
}


func main() {
	// Echo instance

	// var cnt *int
	// cnt = new(int)
	// *cnt = 12345
	
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", getCnt)

	// Start server
	fmt.Println("Start Server from port 9936")
	e.Logger.Fatal(e.Start(":9936"))
}
