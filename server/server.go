package server

import (
	"github.com/myechuri/ukd/server/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
)

type version_type struct {
    Major int32
    Minor int32
}

type ukdServer struct {
    Version version_type
        
}

func (s ukdServer) GetVersion(context context.Context, request *api.VersionRequest) (*api.VersionReply, error) {
        reply := api.VersionReply{
                Major: s.Version.Major,
                Minor: s.Version.Minor}
	grpclog.Fatalf("Version request")
        return &reply, nil
}

func (s ukdServer) StartUK(context context.Context, request *api.StartRequest) (*api.StartReply, error) {
        reply := api.StartReply{
                Success: true,
                Ip: "10.0.0.4"}
	grpclog.Fatalf("Start request")
        return &reply, nil
}

func (s ukdServer) StopUK(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
        reply := api.StopReply{
                Success: true}
	grpclog.Fatalf("Stop request")
        return &reply, nil

}

func NewServer() *ukdServer {
        s := &ukdServer{}
	return s
}
