package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

type Claims struct {
	UserNo int
	jwt.StandardClaims
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

func (wood *WoodDB) GetJwtToken() (string, error) {
	timeOut := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeOut.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("hello"))
	if err != nil {
		return "", fmt.Errorf("token signed Error")
	} else {
		return tokenString, nil
	}
}

func main() {
	woodDB := WoodDB{}
	err := woodDB.Open("localhost", 5432, "wood", os.Args[1], "wood")
	defer woodDB.Close()
	if err != nil {
		panic(err)
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
		username := c.FormValue("username")
		password := c.FormValue("password")
		if woodDB.Login(username, password) == true {
			jwtToken, _ := woodDB.GetJwtToken()
			c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
			c.Response().Header().Set("Last-Modified", time.Now().String())
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("Expires", "-1")
			c.SetCookie(&http.Cookie{
				Name:   "access-token",
				Value:  jwtToken,
				MaxAge: 1800,
			})
			mapD := map[string]interface{}{"isOK": true}
			return c.JSON(http.StatusOK, mapD)

		} else {
			mapD := map[string]interface{}{"isOK": false}
			return c.JSON(http.StatusOK, mapD)
		}

	})

	e.POST("/logout", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":3001"))

}
