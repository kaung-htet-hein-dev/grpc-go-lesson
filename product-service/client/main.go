package main

import (
	"context"
	"log"
	productpb "product-service/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

	res, err := client.GetProduct(ctx, &productpb.GetProductRequest{Id: 0})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("gRPC error: %v, message: %v", st.Code(), st.Message())
		} else {
			log.Fatalf("unknown error: %v", err)
		}
	}

	log.Printf("Got product: %+v", res.GetProduct())
}
