package service

import (
	"github.com/ngenohkevin/pcbook/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

}

func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	laptopServer := NewLaptopServer(NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0") //random available port
	require.NoError(t, err)

	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Printf("cannot create server")
		}
	}()

	return laptopServer, listener.Addr().String()
}
