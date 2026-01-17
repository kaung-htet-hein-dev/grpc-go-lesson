package main

import (
	"context"
	"log"
	"net"
	orderpb "order-service/proto"
	productpb "order-service/proto/product"

	"google.golang.org/grpc"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	productClient productpb.ProductServiceClient
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (
	*orderpb.CreateOrderResponse,
	error) {

	productRes, err := s.productClient.GetProduct(ctx, &productpb.GetProductRequest{Id: req.ProductId})

	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		Order: &orderpb.Order{
			Id:         1,
			ProductId:  req.ProductId,
			Quantity:   req.Quantity,
			TotalPrice: productRes.Product.Price * float64(req.Quantity),
		},
	}, nil
}

func main() {
	// connect to product service

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}

	productClient := productpb.NewProductServiceClient(conn)

	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(grpcServer, &OrderServer{
		productClient: productClient,
	})

	log.Println("Order Service is running on port :50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
