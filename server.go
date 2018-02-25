package main

import (
	"log"
	"net"

	pb "github.com/Nikkison/pp-reservation/pp-reservation"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

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

// func setReservation(reservation_id int32) ()

func main() {
	lis, err := net.Listen("tcp", port)
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
