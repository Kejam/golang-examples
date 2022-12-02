package grpc_client

import (
	"context"
	pb "golang-examples/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func DialGrpcServer(port string) string {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	dial, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer dial.Close()
	client := pb.NewHelloServiceClient(dial)
	request := &pb.Empty{}
	hello, err := client.GetHello(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return hello.Hello
}
