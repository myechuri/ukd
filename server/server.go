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
	grpclog.Printf("Version request")
	return &reply, nil
}

func (s ukdServer) StartUK(context context.Context, request *api.StartRequest) (*api.StartReply, error) {
	reply := api.StartReply{
		Success: false,
		Ip:      "0.0.0.0",
		Reason:  "Not yet implemented"}
	grpclog.Printf("Start request")
	return &reply, nil
}

func (s ukdServer) StopUK(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
	reply := api.StopReply{
		Success: false,
		Reason:  "Not yet implemented"}
	grpclog.Printf("Stop request")
	return &reply, nil

}

func NewServer() *ukdServer {
	s := &ukdServer{Version: version_type{Major: 0, Minor: 1}}
	return s
}
