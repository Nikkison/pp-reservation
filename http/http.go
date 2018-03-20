package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/Nikkison/pp-reservation/pp-reservation"
	"google.golang.org/grpc"
)

const (
	// 予約APIサーバ
	reservationServerAddress = "localhost:50051"
	httpServer               = "localhost:50080"
)

var resvConn pb.ReservationClient

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	querys := r.URL.Query()
	log.Print("querys:", querys)
	w.Write([]byte("hello"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	keys := r.URL.Query()
	log.Print(keys)
	w.Write([]byte("hello"))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data *pb.ReservationResponse `json:"data"`
	}

	querys := r.URL.Query()
	log.Print("querys:", querys)

	resv, err := resvConn.CreateReservation(context.Background(), &pb.CreateReservationRequest{
		SubscriberName: "karino",
		VisitorName:    "alien",
		RoomId:         1,
		TimeZone:       "10:00~20:00",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	json, _ := json.Marshal(&Response{Data: resv})

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}

func main() {
	// 1. sudo apt-get -y install redis-server
	// 2. sudo service redis startd
	// 3. redis-cli
	// 3.1 go get github.com/garyburd/redigo/redis
	// 4. keys *

	// Set up a connection to the server.
	conn, err := grpc.Dial(reservationServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	resvConn = pb.NewReservationClient(conn)

	// 一覧表示Httpレスポンス
	http.HandleFunc("/", rootHandler)
	// 予約詳細jsonレスポンス
	http.HandleFunc("/get", getHandler)
	// 予約作成（jsonレスポンス）
	http.HandleFunc("/create", createHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
