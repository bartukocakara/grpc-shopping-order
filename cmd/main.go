package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bartukocakara/grpc-shopping-order/pkg/client"
	"github.com/bartukocakara/grpc-shopping-order/pkg/config"
	"github.com/bartukocakara/grpc-shopping-order/pkg/db"
	"github.com/bartukocakara/grpc-shopping-order/pkg/pb"
	"github.com/bartukocakara/grpc-shopping-order/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := service.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
