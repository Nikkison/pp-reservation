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

// redis pool インスタンス
var pool = &redis.Pool{
	MaxIdle:   80,
	MaxActive: 12000, // max number of connections
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisPort)
		if err != nil {
			panic(err.Error())
		}
		return c, err
	},
}

func getResvID() int32 {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return r.Int31()
}

// func redisSet(key string, value string) {
// 	redisConn.Do("SET", key, value)
// }

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func redisHmset(in *pb.CreateReservationRequest) {
	redisConn := pool.Get()
	defer redisConn.Close()

	log.Print("hmset:", in)
	resvID := getResvID()

	key := pepperReservationPrefix + strconv.Itoa(int(resvID))
	log.Print("key:", key)
	_, err := redisConn.Do("GET", "KEY")
	fmt.Println(err)

	redisConn.Do("HMSET", key, "reservation_id", resvID, "subscriber_name", in.SubscriberName, "visitor_name", in.VisitorName, "room_id", in.RoomId, "time_zone", in.TimeZone)
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
	redisConn := pool.Get()
	defer redisConn.Close()

	test, _ := redisConn.Do("GET", "KEY")
	fmt.Println(test)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
