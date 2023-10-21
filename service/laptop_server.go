package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ngenohkevin/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
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

	//some heavy processing
	time.Sleep(6 * time.Second)
	//	save the laptop to store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save the laptop to the store: %v", err)
	}
	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{Id: laptop.Id}

	return res, nil
}
