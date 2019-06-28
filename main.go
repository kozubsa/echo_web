package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	e := echo.New()
	e.GET("/", get)
	e.POST("/", post)
	e.Logger.Fatal(e.Start(":80"))
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}

// Handler
func post(c echo.Context) error {
	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}