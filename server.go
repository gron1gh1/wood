package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "wood"
	DB_PASSWORD = "wood"
	DB_NAME     = "wood"
)

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("user=wood password=%s sslmode=disable", os.Args[1]))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,username FROM account")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int
	var username string
	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, username)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/logout", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":3001"))

}
