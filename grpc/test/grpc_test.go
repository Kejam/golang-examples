package test

import (
	"github.com/stretchr/testify/assert"
	grpc_client "golang-examples/grpc/client"
	grpc "golang-examples/grpc/server"
	"testing"
)

func TestGrpcServer(t *testing.T) {
	go grpc.StartGrpcServer(":9000")
	responseMessage := grpc_client.DialGrpcServer(":9000")
	assert.Equal(t, "Hello from grpc server", responseMessage)
}

func TestMockStartGrcpServer(t *testing.T) {
	grpc.StartGrpcServer(":9000")
}
