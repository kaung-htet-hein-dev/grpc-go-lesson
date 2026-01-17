package main

import (
	"context"
	"log"
	orderpb "order-service/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := orderpb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateOrder(ctx, &orderpb.CreateOrderRequest{
		ProductId: 1,
		Quantity:  2,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("Order created: %+v\n", res.Order)
}
