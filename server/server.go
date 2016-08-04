package server

import (
	"github.com/myechuri/ukd/server/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
)

type ukdServer struct {
}

func (s *ukdServer) StartUK(context context.Context, request *api.StartRequest) (*api.StartReply, error) {
        reply := api.StartReply{
                Success: true,
                Ip: "10.0.0.4"}
	grpclog.Fatalf("Start request")
        return &reply, nil
}

func (s *ukdServer) StopUK(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
        reply := api.StopReply{
                Success: true}
	grpclog.Fatalf("Stop request")
        return &reply, nil

}

func NewServer() *api.UkdServer {
        s := new(api.UkdServer)
	return s
}
