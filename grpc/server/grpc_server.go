package grpc

import (
	"context"
	pb "golang-examples/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

func StartGrpcServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &LocalHelloServer{})
	log.Println("Started grpc server at port: ", port)
	grpcServer.Serve(lis)
}

type LocalHelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *LocalHelloServer) GetHello(ctx context.Context, empty *pb.Empty) (*pb.Hello, error) {
	return &pb.Hello{Hello: "Hello from grpc server"}, nil
}
