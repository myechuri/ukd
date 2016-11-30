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

type RuntimeInfo struct {
        Process *os.Process
        Image string
}

type ukdServer struct {
	Version    version_type
        AppRuntime map[string]*RuntimeInfo
	// AppProcess map[string]*os.Process
}

func (s ukdServer) GetVersion(context context.Context, request *api.VersionRequest) (*api.VersionReply, error) {
	reply := api.VersionReply{
		Major: s.Version.Major,
		Minor: s.Version.Minor}
	grpclog.Printf("Version request")
	return &reply, nil
}

func StartQemu(s ukdServer, name string, location string) (*api.StartReply, error) {
	driveArg := "file=" + location + ",if=none,id=hd0,cache=none,aio=native"

	// Compose application-specific configuration.
	configRoot := "/var/lib/ukd"
	os.Mkdir(configRoot, 0777) // TODO: check error
	configRoot += "/" + name
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
        runtime := &RuntimeInfo{Process: cmd.Process,
                                Image: location}
	s.AppRuntime[name] = runtime

	reply := api.StartReply{
		Success: true, // TODO: gather err from previous steps
		Ip:      ip,
		Info:    "Successful start"}
	return &reply, nil
}

func (s ukdServer) Start(context context.Context, request *api.StartRequest) (*api.StartReply, error) {
	grpclog.Printf("Start request: name: %s, Image: %s", request.Name, request.Location)

	// Validate image exists.
	if _, err := os.Stat(request.Location); os.IsNotExist(err) {
		reply := api.StartReply{
			Success: false,
			Ip:      "",
			Info:    "Image " + request.Location + " does not exist, error: " + err.Error()}
		return &reply, nil
	}

	// Validate application name does not exist.
	if s.AppRuntime[request.Name] != nil {
		reply := api.StartReply{
			Success: false,
			Ip:      "",
			Info:    request.Name + " is already running. Please choose a different name for the application if you wish to start a second instance using the same image."}
		return &reply, nil
	}

	if request.Visor == "kvm-qemu" {
		reply, _ := StartQemu(s, request.Name, request.Location)
		return reply, nil
	} else {
		reply := api.StartReply{
			Success: false,
			Ip:      "",
			Info:    "Requested hypervisor (" + request.Visor + ") is not yet supported."}
		return &reply, nil
	}
}

func (s ukdServer) Stop(context context.Context, request *api.StopRequest) (*api.StopReply, error) {
	grpclog.Printf("Stop request: name: %s", request.Name)
	var success bool
	var info string
	runtime := s.AppRuntime[request.Name]
	if runtime == nil {
		success = true
		info = "App not found. Nothing to do."
	} else {
		runtime.Process.Signal(syscall.SIGTERM) // TODO: check error
		pstate, _ := runtime.Process.Wait()     // TODO: check err
		if pstate.Exited() {
			success = true
			info = "Successfully stopped Application (" + request.Name + ")"
			delete(s.AppRuntime, request.Name)
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

func ApplyDiff(base string, diff []byte) (bool, string) {
      var success bool
      var info string

      workDir, _ := ioutil.TempDir("", "ukd-update-")
      deltaFile := workDir + "/deltaFile"
      f, err := os.Create(deltaFile)
      if err != nil {
         grpclog.Printf("Failed to create temp file")
         success = false
         info = "Failed to create delta file " + deltaFile + ", error: " + err.Error()
         return success,info
      }
      err = ioutil.WriteFile(deltaFile, diff, 0700)
      f.Close()

      updatedImagePath := workDir + "/newImage.img"
      cmdName := "rdiff"
      args := []string{"patch", base, deltaFile, updatedImagePath}
      cmd := exec.Command(cmdName, args...)
      err = cmd.Run()
      if err != nil {
         success = false
         info = "Failed to patch (" + base + " with " + deltaFile + ", error: " + err.Error()
      } else {
         success = true
         info = "Successfully patched image at " + updatedImagePath
      }

      // defer os.RemoveAll(workDir)
      return success,info
}

func (s ukdServer) UpdateImage(context context.Context, request *api.UpdateImageRequest) (*api.UpdateImageReply, error) {
	grpclog.Printf("Update request: base=%s", request.Base)
	var success bool
	var info string

        // TODO: check that image is not currently in use.
        success, info = ApplyDiff(request.Base, request.Diff)

	reply := api.UpdateImageReply{
		Success: success,
		Info:    info}
	grpclog.Printf("Update image request")
	return &reply, nil
}

func NewServer() *ukdServer {
	s := &ukdServer{Version: version_type{Major: 0, Minor: 1},
		AppRuntime: make(map[string]*RuntimeInfo)}

        // Image home.
        imagePath := "/var/lib/ukd/images"
        err := os.MkdirAll("/var/lib/ukd/images", 0700)
	if err != nil {
            grpclog.Printf("MkdirAll failed %q: %s", imagePath, err)
	}

	return s
}
