package server

import (
	"bufio"
	"github.com/myechuri/ukd/server/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

type version_type struct {
	Major int32
	Minor int32
}

type ukdServer struct {
	Version    version_type
	AppProcess map[string]*os.Process
}

func (s ukdServer) GetVersion(context context.Context, request *api.VersionRequest) (*api.VersionReply, error) {
	reply := api.VersionReply{
		Major: s.Version.Major,
		Minor: s.Version.Minor}
	grpclog.Printf("Version request")
	return &reply, nil
}

func (s ukdServer) Start(context context.Context, request *api.StartRequest) (*api.StartReply, error) {
	grpclog.Printf("Start request: name: %s, Image: %s", request.Name, request.Location)

	driveArg := "file=" + request.Location + ",if=none,id=hd0,cache=none,aio=native"

	// Compose application-specific configuration.
	configRoot := "/var/lib/ukd"
	os.Mkdir(configRoot, 0777) // TODO: check error
	configRoot += "/" + request.Name
	os.Mkdir(configRoot, 0777) // TODO: check error
	qemuIfupByteArray := []byte("#!/bin/sh\n" +
		"brctl stp virbr0 off\n" +
		"brctl addif virbr0 $1\n" +
		"ifconfig $1 up\n")
	configRoot += "/qemu-ifup.sh"
	ioutil.WriteFile(configRoot, qemuIfupByteArray, 0700) // TODO: check error
	netdevArg := "tap,id=hn0,script=" + configRoot + ",vhost=on"

	cmdName := "qemu-system-x86_64"
	args := []string{
		"-m", "2G",
		"-smp", "4",
		"-vnc", ":1",
		"-gdb", "tcp::1234,server,nowait",
		"-device", "virtio-blk-pci,id=blk0,bootindex=0,drive=hd0,scsi=off",
		"-drive", driveArg,
		"-netdev", netdevArg,
		"-device", "virtio-net-pci,netdev=hn0,id=nic0",
		"-redir", "tcp:2222::22",
		"-device", "virtio-rng-pci",
		"-enable-kvm",
		"-cpu", "host,+x2apic",
		"-chardev", "stdio,mux=on,id=stdio,signal=off",
		"-mon", "chardev=stdio,mode=readline,default",
		"-device", "isa-serial,chardev=stdio"}
	cmd := exec.Command(cmdName, args...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	r := bufio.NewReader(stdout)
	matched := false
	var line []byte
	for !(matched) {
		line, _, _ = r.ReadLine()
		matched, _ = regexp.MatchString("eth0:.*", string(line))
	}
	ip := strings.Fields(string(line))[1]
	s.AppProcess[request.Name] = cmd.Process

	reply := api.StartReply{
		Success: true, // TODO: gather err from previous steps
		Ip:      ip,
		Info:    "Successful start"}
	return &reply, nil
}

func (s ukdServer) Stop(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
	grpclog.Printf("Stop request: name: %s", request.Name)
	var success bool
	var info string
	process := s.AppProcess[request.Name]
	if process == nil {
		success = true
		info = "App not found. Nothing to do."
	} else {
		process.Signal(syscall.SIGTERM) // TODO: check error
		pstate, _ := process.Wait() // TODO: check err
                if pstate.Exited() {
		    success = true
		    info = "Successfully stopped Application"
                    delete(s.AppProcess, request.Name)
                } else {
                    success = false
                    info = pstate.String()
                }
	}
	reply := api.StopReply{
		Success: success,
		Info:    info}
	grpclog.Printf("Stop request")
	return &reply, nil

}

func NewServer() *ukdServer {
	s := &ukdServer{Version: version_type{Major: 0, Minor: 1},
		AppProcess: make(map[string]*os.Process)}
	return s
}
