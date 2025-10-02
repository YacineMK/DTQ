package grpcclients

import (
	trippb "github.com/YacineMK/DTQ/shared/proto/trip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TripService struct {
	Client trippb.TripServiceClient
	Conn   *grpc.ClientConn
}

func NewTripServiceClient(addr string) (*TripService, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := trippb.NewTripServiceClient(conn)
	return &TripService{Client: client, Conn: conn}, nil
}
