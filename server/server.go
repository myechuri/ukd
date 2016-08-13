package server

import (
        "os/exec"
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
        grpclog.Printf("StartUK request: name: %s, Image: %s", request.Name, request.Location)
        cmdName := "qemu-system-x86_64"
        args := []string{
        "-m", "2G",
        "-smp", "4",
        "-vnc", ":1",
        "-gdb", "tcp::1234,server,nowait",
        "-device", "virtio-blk-pci,id=blk0,bootindex=0,drive=hd0,scsi=off",
        "-drive", "file=/root/osv/build/debug/usr.img,if=none,id=hd0,cache=none,aio=native",
        "-netdev", "tap,id=hn0,script=scripts/qemu-ifup.sh,vhost=on",
        "-device", "virtio-net-pci,netdev=hn0,id=nic0",
        "-redir", "tcp:2222::22",
        "-device", "virtio-rng-pci",
        "-enable-kvm",
        "-cpu", "host,+x2apic",
        "-chardev", "stdio,mux=on,id=stdio,signal=off",
        "-mon", "chardev=stdio,mode=readline,default",
        "-device", "isa-serial,chardev=stdio", }
        cmd := exec.Command(cmdName, args...)
        if err := cmd.Start(); err != nil {
            grpclog.Fatalf(err)
        }
	reply := api.StartReply{
		Success: false,
		Ip:      "0.0.0.0",
		Reason:  "Not yet implemented"}
	grpclog.Printf("Start request")
	return &reply, nil
}

func (s ukdServer) StopUK(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
        grpclog.Printf("StopUK request: name: %s", request.Name)
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
