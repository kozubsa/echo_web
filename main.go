package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	e := echo.New()
	e.GET("/", get)
	e.POST("/", post)
	e.Logger.Fatal(e.Start(":8080"))
}

func get(c echo.Context) error {
	log.Printf("\n\tMethod:\t%v\n\tHeader:\t%v\n\tQuery:\t%v\n\n",
		c.Request().Method, c.Request().Header, c.Request().URL.RawQuery)

	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}

// Handler
func post(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	r, err := url.ParseQuery(string(body))
	if err != nil {
		return err
	}

	log.Println(r)

	log.Printf("\n\tMethod:\t%v\n\tHeader:\t%v\n\tQuery:\t%v\n\tBody:\t%v\n\n",
		c.Request().Method, c.Request().Header, c.Request().URL.RawQuery, string(body))

	var Response = struct {
		StatusCode string `json:"status_code"`
		RequestBody string `json:"request_body"`
	}{
		StatusCode: "Ok",
		RequestBody: string(body),
	}

	return c.JSON(http.StatusOK, Response)
}
