package main

import (
	"context"
	"flag"
	"github.com/Abdurazzoq789/pcbook/pb"
	"github.com/Abdurazzoq789/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	//laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	//set timeout

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//not a big deal
			log.Print("laptop already exist")
		} else {
			log.Fatal("cannot create laptop :", err)
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)
}