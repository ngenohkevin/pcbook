package main

import (
	"context"
	"flag"
	"github.com/ngenohkevin/pcbook/pb"
	"github.com/ngenohkevin/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	//set timeout

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exist")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)
}
