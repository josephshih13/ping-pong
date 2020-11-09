package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cnt = 0

func getCnt(c echo.Context) error {

	ret := fmt.Sprintf("Pong %d\n", cnt)
	// d1 := []byte(ret)
	// err := ioutil.WriteFile("/home/ec2-user/environment/pong.txt", d1, 0644)
	// check(err)
	cnt = cnt + 1
	return c.String(http.StatusOK, ret)
}

func getCnt2(c echo.Context) error {

	ret := fmt.Sprintf("Pong %d\n", cnt)
	// d1 := []byte(ret)
	// err := ioutil.WriteFile("/home/ec2-user/environment/pong.txt", d1, 0644)
	// check(err)
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
	e.GET("/pong", getCnt)
	e.GET("/internal", getCnt2)

	// Start server
	fmt.Println("Start Server from port 9936")
	e.Logger.Fatal(e.Start(":9936"))
}
