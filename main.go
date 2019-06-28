package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"google.golang.org/grpc"

	pb "github.com/introphin/proto-advertiser-usecase"
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
	m := Client{}
	if err := c.Bind(&m); err != nil {
		log.Println(err)
		return err
	}

	uuid := m.Uuid
	if m.Uuid == "" {
		uuid = m.Tuuid
	}

	log.Println("get order uuid:", uuid)

	go sendPostback(uuid)

	return c.String(http.StatusOK, fmt.Sprintf(`{"status_code": "Ok", "uuid": %v}`, uuid))
}

// Handler
func post(c echo.Context) error {
	return c.String(http.StatusOK, `{"status_code": "Ok"}`)
}

func sendPostback(uuid string) {
	rand.Seed(time.Now().UTC().UnixNano())
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	r := Repository{
		adv: pb.NewAdvertiserUsecaseServiceClient(conn),
	}

	postback, err := r.CreatePostback(uuid)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("postback %#v\r\n", postback)
}

func (r Repository) CreatePostback(uuid string) (Postback, error) {
	if uuid == "" {
		log.Println("uuid empty")
		log.Println("uuid empty")
		log.Println("uuid empty")
		log.Println("uuid empty")
		log.Println("uuid empty")
		log.Println("uuid empty")
	}
	request := &pb.CreatePostbackRequest{
		UserId:        int32(rand.Intn(999)),
		ClickId:       uuid,
		Payout:        rand.Float32(),
		AccessToken:   "AccessToken",
		TransactionId: "TransactionId",
		Status:        "enable",
	}
	pbPostback, err := r.adv.CreatePostback(context.Background(), request)
	if err != nil {
		return Postback{}, err
	}

	postback := Postback{
		Id:            pbPostback.Id,
		Uuid:          pbPostback.Uuid,
		UserId:        pbPostback.UserId,
		Stage:         pbPostback.Stage,
		ClickId:       pbPostback.ClickId,
		Payout:        pbPostback.Payout,
		AccessToken:   pbPostback.AccessToken,
		TransactionId: pbPostback.TransactionId,
		Status:        pbPostback.Status,
	}
	if pbPostback.CreatedAt != nil {
		ca := time.Unix(pbPostback.CreatedAt.Seconds, 0).UTC()
		postback.CreatedAt = &ca
	}
	if pbPostback.UpdatedAt != nil {
		ua := time.Unix(pbPostback.UpdatedAt.Seconds, 0).UTC()
		postback.UpdatedAt = &ua
	}

	return postback, nil
}
