package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/Nikkison/pp-reservation/pp-reservation"
	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	serverPort              = ":50051"
	redisPort               = ":6379"
	pepperReservationPrefix = "ppresv"
)

var redisConn redis.Conn

func getResvID() int32 {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return r.Int31()
}

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

func redisHmset(in &pb.CreateReservationRequest) {
	resvID := getResvID()
	key := pepperReservationPrefix + strconv.Itoa(int(resvID))
	redisConn.Do("HMSET", key, "reservation_id", resvID, "subscriber_name", in.SubscriberName, "visitor_name", in.VisitorName, "room_id", in.RoomId, "time_zone", in.TimeZone)
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
	log.Print("crr:", in)
	redisHmset(in)
	return &pb.ReservationResponse{
		ReservationId:  1,
		SubscriberName: "karino",
		VisitorName:    "ookubo",
		RoomId:         1,
		TimeZone:       "9:00 ~ 10:00"}, nil
}

func main() {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReservationServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// redis connect
	redisConn := redisConnection()
	defer redisConn.Close()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
