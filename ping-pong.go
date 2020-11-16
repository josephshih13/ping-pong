package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createtable() {
	_, err := conn.Exec(context.Background(), "CREATE TABLE ping (ping_count INT NOT NULL);INSERT INTO ping (ping_count)VALUES(0);")
	check(err)
}

func check_table_exist() bool {
	var exist bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS (SELECT FROM information_schema.tables WHERE  table_name = 'ping');").Scan(&exist)
	check(err)
	return exist
}

func read_number() int {
	if !check_table_exist() {
		createtable()
		return 0
	}
	var num int
	err := conn.QueryRow(context.Background(), "SELECT * from ping;").Scan(&num)
	check(err)
	return num
}

func add_number() {
	num := read_number()
	num = num + 1
	_, err := conn.Exec(context.Background(), "UPDATE ping SET ping_count = $1;", num)
	check(err)
}

func getCnt(c echo.Context) error {
	cnt := read_number()
	ret := fmt.Sprintf("Pong %d\n", cnt)
	add_number()
	return c.String(http.StatusOK, ret)
}

func getCnt2(c echo.Context) error {
	cnt := read_number()
	ret := fmt.Sprintf("Pong %d\n", cnt)
	return c.String(http.StatusOK, ret)
}

func main() {

	var err error
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		db_url = "postgresql://postgres:test1234@joseph-public.chld9kh33qyg.us-east-1.rds.amazonaws.com:5432/postgres"
	}
	conn, err = pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

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
