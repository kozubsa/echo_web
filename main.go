package main

import (
	"math/rand"
	"net/http"
	"time"

	pb "github.com/introphin/proto-advertiser-usecase"
	"github.com/labstack/echo"
)

type Repository struct {
	adv pb.AdvertiserUsecaseServiceClient
}

type Postback struct {
	Id            int32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Uuid          string
	UserId        int32
	Stage         int32
	ClickId       string
	Payout        float32
	AccessToken   string
	TransactionId string
	Status        string
}

type Client struct {
	Uuid        string `json:"uuid" form:"uuid" query:"uuid"`
	Tuuid       string `json:"test_order_uuid" form:"test_order_uuid" query:"test_order_uuid"`
	OfferHash   string `json:"offer_hash" form:"offer_hash" query:"offer_hash"`
	ClientName  string `json:"client_name" form:"client_name" query:"client_name"`
	ClientPhone string `json:"client_phone" form:"client_phone" query:"client_phone"`
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	e := echo.New()
	e.GET("/", get)
	e.POST("/", post)
	e.Logger.Fatal(e.Start(":8080"))
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}

// Handler
func post(c echo.Context) error {
	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}