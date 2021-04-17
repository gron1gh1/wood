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

type WoodDB struct {
	db *sql.DB
}

func (wood *WoodDB) Open(host string, port int, user string, password string, dbname string) error {
	var err error
	wood.db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	return err
}

func (wood *WoodDB) Close() {
	wood.db.Close()
}

func (wood *WoodDB) Login(username string, password string) bool {
	var id int
	_ = wood.db.QueryRow("SELECT id FROM account WHERE username=$1 AND password=$2",
		username,
		password).Scan(&id)
	if id == 0 {
		return false
	}
	return true
}

func main() {
	woodDB := WoodDB{}
	err := woodDB.Open("localhost", 5432, "wood", os.Args[1], "wood")
	defer woodDB.Close()
	if err != nil {
		panic(err)
	}
	if woodDB.Login("gron1gh1", "") == true {
		fmt.Printf("로그인 성공")
	} else {
		fmt.Printf("로그인 실패")
	}
	// rows, err := woodDB.db.Query("SELECT id,username FROM account")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// var id int
	// var username string
	// for rows.Next() {
	// 	err := rows.Scan(&id, &username)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(id, username)
	// }

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
