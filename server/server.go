package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Nikkison/pp-reservation/pp-reservation"
	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	serverPort = ":50051"
	redisPort  = ":6379"
)

func redisSet(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func redisConnection() redis.Conn {
	//redisに接続
	c, err := redis.Dial("tcp", redisPort)
	if err != nil {
		panic(err)
	}
	return c
}

type server struct{}

// 予約情報の詳細を実装
func (s *server) GetReservation(ctx context.Context, in *pb.ReservationRequest) (*pb.ReservationResponse, error) {
	return &pb.ReservationResponse{
		ReservationId:  in.ReservationId,
		SubscriberName: "karino",
		VisitorName:    "ookubo",
		RoomId:         1,
		TimeZone:       "9:00 ~ 10:00"}, nil
}

// 予約情報作成
func (s *server) CreateReservation(ctx context.Context, in *pb.CreateReservationRequest) (*pb.ReservationResponse, error) {
	return &pb.ReservationResponse{
		ReservationId:  1,
		SubscriberName: "karino",
		VisitorName:    "ookubo",
		RoomId:         1,
		TimeZone:       "9:00 ~ 10:00"}, nil
}

// func setReservation(reservation_id int32) ()

func main() {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReservationServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
