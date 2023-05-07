package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	product "grpcdemov2/proto/product"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := product.NewProductServiceClient(conn)
	response, err := client.GetProduct(context.Background(), &product.GetProductRequest{Id: "aaa"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
