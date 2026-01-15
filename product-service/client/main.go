package main

import (
	"context"
	"log"
	productpb "product-service/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := productpb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetProduct(ctx, &productpb.GetProductRequest{Id: 1})
	if err != nil {
		log.Fatalf("GetProduct failed: %v", err)
	}

	log.Printf("Got product: %+v", res.GetProduct())
}
