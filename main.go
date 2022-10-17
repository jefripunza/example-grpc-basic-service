package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"basic-service/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedBasicServiceServer
}

func main() {
	godotenv.Load(".env")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("GRPC_PORT")))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterBasicServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Printf("gRPC Started at http://localhost:%v ...\n", os.Getenv("GRPC_PORT"))
	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}

func (s *server) Add(_ context.Context, request *proto.RequestBasic) (*proto.ResponseBasic, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b
	fmt.Printf("%v + %v = %v\n", a, b, result)

	return &proto.ResponseBasic{Result: result}, nil
}

func (s *server) Multiply(_ context.Context, request *proto.RequestBasic) (*proto.ResponseBasic, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b
	fmt.Printf("%v * %v = %v\n", a, b, result)

	return &proto.ResponseBasic{Result: result}, nil
}
