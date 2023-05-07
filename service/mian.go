package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"grpcdemov2/proto/product"
	p "grpcdemov2/service/product"
	"log"
	"net"
	"net/http"
)

func main() {
	go startGRPCGateway()

	grpcServer := grpc.NewServer()
	product.RegisterProductServiceServer(grpcServer, new(p.ProductService))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(grpcServer.Serve(lis))
	log.Println("Server end")
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
		},
	}))
	err := product.RegisterProductServiceHandlerFromEndpoint(c, mux, ":1234", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cann't start grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8080", mux) // grpc gateway 的端口
	if err != nil {
		log.Fatalf("cann't listen and serve: %v", err)
	}
	log.Println("Gateway end")
}
