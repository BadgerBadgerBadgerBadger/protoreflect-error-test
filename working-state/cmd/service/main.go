package main

import (
	"context"
	"net"

	"github.com/BadgerBadgerBadgerBadger/goplay/pkg/util"
	"github.com/BadgerBadgerBadgerBadger/protoreflect-error-test/proto/lang_go/dependencies/secondary"
	"github.com/BadgerBadgerBadgerBadger/protoreflect-error-test/proto/lang_go/primary"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:8000")
	util.Must(err)

	grpcServer := googleGrpc.NewServer()

	reflection.Register(grpcServer)
	primary.RegisterMainServer(grpcServer, mainService{})

	err = grpcServer.Serve(lis)
	util.Must(err)
}

type mainService struct {
	primary.UnimplementedMainServer
}

func (p mainService) GetPrimary(ctx context.Context, request *primary.PrimaryRequest) (*primary.PrimaryResponse, error) {
	return &primary.PrimaryResponse{
		PrimaryResponse: &secondary.SecondaryMessage{
			Value: "Hello from primary",
		},
	}, nil
}
