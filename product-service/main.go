package main

import (
	"context"
	"log"
	"net"
	productpb "product-service/proto"

	"google.golang.org/grpc"
)

type ProductServer struct {
	productpb.UnimplementedProductServiceServer
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (
	*productpb.CreateProductResponse,
	error) {
	product := &productpb.Product{
		Id:          1,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	return &productpb.CreateProductResponse{Product: product}, nil
}

func (s *ProductServer) GetProduct(
	ctx context.Context,
	req *productpb.GetProductRequest,
) (*productpb.GetProductResponse, error) {

	product := &productpb.Product{
		Id:          req.Id,
		Name:        "iPhone 16",
		Description: "Apple flagship phone",
		Price:       4999.99,
		Stock:       10,
	}

	return &productpb.GetProductResponse{
		Product: product,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	productpb.RegisterProductServiceServer(grpcServer, &ProductServer{})
	log.Println("Product gRPC server running on :50051")

	grpcServer.Serve(lis)
}
