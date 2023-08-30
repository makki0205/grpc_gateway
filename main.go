package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/makki0205/grpc_gateway/pkg/gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type Service struct {
}

func (s Service) mustEmbedUnimplementedTestServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	return &pb.HelloRes{Msg: "hello " + req.Name}, nil
}

func main() {
	CreateGRPCServer()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
func CreateGRPCServer() {
	// 1. 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()
	service := Service{}
	pb.RegisterTestServiceServer(s, &service)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
}
