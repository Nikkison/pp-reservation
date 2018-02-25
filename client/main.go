package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/Nikkison/pp-reservation/pp-reservation"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReservationClient(conn)

	// Contact the server and print out its response.
	id := 1
	if len(os.Args) > 1 {
		id, _ = strconv.Atoi(os.Args[1])
	}
	r, err := c.GetReservation(context.Background(), &pb.ReservationRequest{ReservationId: int32(id)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reservation: %+v", r)
}
