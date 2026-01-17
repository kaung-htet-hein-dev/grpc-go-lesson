package server

import (
	"log"
	"net"
	"product-service/internal/usecase"
	productpb "product-service/proto"

	grpcHandler "product-service/internal/delivery/grpc"

	"google.golang.org/grpc"
)

func Execute() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	productUsecase := usecase.NewProductUsecase()
	productHandler := grpcHandler.NewProductHandler(productUsecase)

	productpb.RegisterProductServiceServer(grpcServer, productHandler)

	log.Println("Product Service running on :50051")
	grpcServer.Serve(lis)
}
