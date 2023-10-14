package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngenohkevin/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type LaptopServer struct {
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieve a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		//	check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "cannot generate anew laptop ID: %v", err)
			}
			laptop.Id = id.String()
		}
	}
}
