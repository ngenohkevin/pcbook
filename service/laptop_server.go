package service

import (
	"context"
	"github.com/ngenohkevin/pcbook/pb"
)

type LaptopServer struct {
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {

}
