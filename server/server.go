package server

import (
	"fmt"

	pb "github.com/myechuri/ukd/server/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type ukdServer struct {
}

func (s *ukdServer) StartUK(context context.Context, request *pb.StartRequest) (*pb.StartResponse, error) {
	grpclog.Fatalf("Start request")
}

func (s *ukdServer) StopUK(context context.Context, request *pb.StopRequest) (*pb.StopResponse, error) {
	grpclog.Fatalf("Stop request")
}

func NewServer() *pb.ukdServer {
        s := new(ukdServer)
	return s
}
