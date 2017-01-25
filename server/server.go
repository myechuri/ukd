package server

import (
	"bufio"
	"bytes"
	"github.com/myechuri/ukd/server/api"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

type version_type struct {
	Major int32
	Minor string
}

const (
	X86_64     = "x86_64"
	ARMv71     = "armv7l"
	IMAGE_PATH = "/var/lib/ukd/images"
	LOG_PATH   = "/var/log/ukd"
)

type PlatformRuntimeInfo struct {
	platform string
}

type AppRuntimeInfo struct {
	Process     *os.Process
	Image       string
	Ip          string
	Incarnation uuid.UUID
}

type ukdServer struct {
	Version         version_type
	PlatformRuntime PlatformRuntimeInfo
	AppRuntime      map[string]*AppRuntimeInfo
}

var arch string

func (s ukdServer) GetVersion(context context.Context, request *api.VersionRequest) (*api.VersionReply, error) {
	reply := api.VersionReply{
		Major: s.Version.Major,
		Minor: s.Version.Minor}
	grpclog.Printf("Version request")
	return &reply, nil
}

func ComposeQemuX86_64Command(name string, location string, log string) (string, []string, error) {
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
	logFileArg := "file:" + log

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
		"-serial", logFileArg}

	return cmdName, args, nil
}

func ComposeQemuAarch64Command(name string, location string, log string) (string, []string, error) {

	// Use user networking for now.
	// TODO: add tap netdev.
	cmdName := "qemu-system-aarch64"
	logFileArg := "file:" + log
	args := []string{
		"-machine", "virt",
		"-cpu", "cortex-a57",
		"-kernel", location,
		"--nographic",
		"-serial", logFileArg}

	return cmdName, args, nil

}

func getAppIP(log string) (string, error) {

	time.Sleep(time.Second)
	logFileExists := false
	for !(logFileExists) {
		_, err := os.Stat(log)
		if err == nil {
			logFileExists = true
		}
		time.Sleep(time.Second)
	}

	file, err := os.Open(log)
	defer file.Close()

	// TODO: Handle app restart case where there exists log file content from previous run.
	r := bufio.NewReader(file)
	matched := false
	var line []byte
	for !(matched) {
		line, _, err = r.ReadLine()
		if err != nil {
			grpclog.Printf("Error %s reading %s, sleeping 5 seconds for log content to be available", err, log)
			if err.Error() == "EOF" {
				time.Sleep(5 * time.Second)
			} else {
				// TODO: Handle read errors.
			}
		} else {
			matched, _ = regexp.MatchString("eth0:.*", string(line))
		}
	}
	ip := strings.Fields(string(line))[1]

	return ip, err
}

func StartQemu(s ukdServer, name string, location string) (*api.StartReply, error) {

	var cmdName string
	var args []string

	// Generate a uuid for in-use application incarnation.
	appUuid := uuid.NewV4()

	logLocation := getLogLocation(name, appUuid)
	// TODO: Handle error returned by ComposeQemu*
	if arch == X86_64 {
		cmdName, args, _ = ComposeQemuX86_64Command(name, location, logLocation)
	} else if arch == ARMv71 {
		// TODO: Add log location.
		cmdName, args, _ = ComposeQemuAarch64Command(name, location, logLocation)
	} else {
		reply := &api.StartReply{
			Success: false,
			Ip:      "",
			Info:    "Platform " + arch + " is not yet supported."}
		return reply, nil
	}

	cmd := exec.Command(cmdName, args...)

	// Disable Glibc's per-thread arena to limit qemu virtual memory.
	// [ References:
	// 1. https://siddhesh.in/posts/malloc-per-thread-arenas-in-glibc.html
	// 2. https://devcenter.heroku.com/articles/tuning-glibc-memory-behavior ]
	cmd.Env = []string{"MALLOC_ARENA_MAX=1"}
	cmd.Start()

	// TODO: handle error path from retriving IP.
	ip, _ := getAppIP(logLocation)

	runtime := &AppRuntimeInfo{Process: cmd.Process,
		Image:       location,
		Ip:          ip,
		Incarnation: appUuid}
	s.AppRuntime[name] = runtime

	reply := api.StartReply{
		Success: true, // TODO: gather err from previous steps
		Ip:      ip,
		Info:    "Successful start"}
	return &reply, nil
}

func getLogLocation(name string, incarnation uuid.UUID) string {
	logLocation := LOG_PATH + "/" + name + "-" + incarnation.String() + ".log"
	return logLocation
}

func (s ukdServer) GetLog(context context.Context, request *api.LogRequest) (*api.LogReply, error) {
	grpclog.Printf("Log request: name: %s", request.Name)

	if s.AppRuntime[request.Name] == nil {
		reply := api.LogReply{
			Success:    true,
			LogContent: nil,
			Info:       "Application (" + request.Name + ") is currently stopped. No log to report."}
		return &reply, nil
	}
	incarnation := s.AppRuntime[request.Name].Incarnation
	logLocation := getLogLocation(request.Name, incarnation)

	_, err := os.Stat(logLocation)
	if err != nil {
		reply := api.LogReply{
			Success:    false,
			LogContent: nil,
			Info:       request.Name + "does not have a log at " + logLocation + "on the compute node."}
		return &reply, nil
	}
	logContent, err := ioutil.ReadFile(logLocation)
	if err != nil {
		reply := api.LogReply{
			Success:    false,
			LogContent: nil,
			Info:       "Failed to read content from " + logLocation + ", error: " + err.Error()}
		return &reply, nil
	}
	grpclog.Printf("Sending log content for %s from %s", request.Name, logLocation)
	reply := api.LogReply{
		Success:    true,
		LogContent: logContent,
		Info:       ""}
	return &reply, nil
}

func (s ukdServer) Status(context context.Context, request *api.StatusRequest) (*api.StatusReply, error) {
	grpclog.Printf("Status request: name: %s", request.Name)

	if s.AppRuntime[request.Name] != nil {
		ip := s.AppRuntime[request.Name].Ip
		reply := api.StatusReply{
			Success: true,
			Status:  "RUNNING",
			Info:    "IP: " + ip}
		return &reply, nil
	} else {
		reply := api.StatusReply{
			Success: true,
			Status:  "STOPPED",
			Info:    ""}
		return &reply, nil
	}

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

	var reply *api.StartReply
	if request.Visor == "kvm-qemu" {
		reply, _ = StartQemu(s, request.Name, request.Location)
	} else {
		reply = &api.StartReply{
			Success: false,
			Ip:      "",
			Info:    "Requested hypervisor (" + request.Visor + ") is not yet supported."}
	}
	return reply, nil
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

func ComputeSignature(path string) (bool, []byte, string) {
	var success bool
	var info string
	var signature []byte

	if _, err := os.Stat(path); os.IsNotExist(err) {
		success = false
		info = "File not found: " + path
		return success, nil, info
	}

	workDir, _ := ioutil.TempDir("", "ukd-compute-signature-")
	defer os.RemoveAll(workDir)

	// Validate base image signature sent by client matches
	// server base image signature.
	serverImageSignature := workDir + "/serverSignature"
	cmdName := "rdiff"
	args := []string{"signature", path, serverImageSignature}
	cmd := exec.Command(cmdName, args...)
	err := cmd.Run()
	if err != nil {
		success = false
		info = "Failed to compute signature for " + path + ", error: " + err.Error()
		return success, nil, info
	}
	signature, err = ioutil.ReadFile(serverImageSignature)
	success = true
	info = "Successfully computed signature"

	// TODO: delete workDir
	return success, signature, info

}

func (s ukdServer) GetImageSignature(context context.Context, request *api.ImageSignatureRequest) (*api.ImageSignatureReply, error) {
	grpclog.Printf("Signature request: image=%s", request.Path)
	var success bool
	var info string
	var signature []byte

	// TODO: check that image is not currently in use.
	success, signature, info = ComputeSignature(request.Path)

	reply := api.ImageSignatureReply{
		Success:   success,
		Signature: signature,
		Info:      info}
	grpclog.Printf("GetImageSignature request")
	return &reply, nil
}

func ApplyDiff(base string, basesig []byte, diff []byte) (bool, string, string) {
	var success bool
	var info string

	workDir, _ := ioutil.TempDir("", "ukd-update-stage-")

	// Validate base image signature sent by client matches
	// server base image signature.
	serverImageSignature := workDir + "/serverSignature"
	cmdName := "rdiff"
	args := []string{"signature", base, serverImageSignature}
	cmd := exec.Command(cmdName, args...)
	err := cmd.Run()
	serverSignature, err := ioutil.ReadFile(serverImageSignature)
	defer os.RemoveAll(serverImageSignature)
	if !bytes.Equal(serverSignature, basesig) {
		success = false
		info = "Diff was generated for a different base image than " + base
		return success, "", info
	}

	// Write out diff to delta file.
	deltaFile := workDir + "/deltaFile"
	f, err := os.Create(deltaFile)
	if err != nil {
		grpclog.Printf("Failed to create temp file")
		success = false
		info = "Failed to create delta file " + deltaFile + ", error: " + err.Error()
		return success, "", info
	}
	err = ioutil.WriteFile(deltaFile, diff, 0700)
	f.Close()
	defer os.RemoveAll(deltaFile)

	updatedImagePath := workDir + "/newImage.img"
	cmdName = "rdiff"
	args = []string{"patch", base, deltaFile, updatedImagePath}
	cmd = exec.Command(cmdName, args...)
	err = cmd.Run()
	if err != nil {
		success = false
		info = "Failed to patch (" + base + " with " + deltaFile + ", error: " + err.Error()
	} else {
		success = true
		info = "Successfully staged patched image at " + updatedImagePath + ". Please validate the image before replacing master copy."
	}

	return success, updatedImagePath, info
}

func (s ukdServer) UpdateImage(context context.Context, request *api.UpdateImageRequest) (*api.UpdateImageReply, error) {
	grpclog.Printf("Update request: base=%s", request.Base)
	var success bool
	var info string

	// TODO: check that image is not currently in use.
	var newImagePath string
	success, newImagePath, info = ApplyDiff(request.Base, request.Basesig, request.Diff)

	if success {
		var newSignature []byte
		success, newSignature, info = ComputeSignature(newImagePath)
		if !bytes.Equal(newSignature, request.Newsig) {
			success = false
			info = "New image signature on destination does not match new image signature on source"
		} else {
			info = "Verified signature match for new Image on source and destination"
		}
	}

	reply := api.UpdateImageReply{
		Success:         success,
		StagedImagePath: newImagePath,
		Info:            info}
	grpclog.Printf("UpdateImage: success: %t, staged image path: %s, info: %s", success, newImagePath, info)
	return &reply, nil
}

func getPlatformRuntime() (PlatformRuntimeInfo, error) {

	cmdName := "arch"
	args := []string{}
	cmd := exec.Command(cmdName, args...)
	archBytes, err := cmd.Output()
	arch = strings.TrimSpace(string(archBytes))
	platformRuntimeInfo := PlatformRuntimeInfo{platform: arch}
	grpclog.Printf("Detected arch: %s on the system", arch)
	return platformRuntimeInfo, err
}

func NewServer() (*ukdServer, error) {

	var s *ukdServer
	platform, err := getPlatformRuntime()
	if err != nil {
		grpclog.Printf("Failed to detect runtime platform, error: %s", err.Error())
		return s, err
	}

	s = &ukdServer{Version: version_type{Major: 0, Minor: "01dev"},
		PlatformRuntime: platform,
		AppRuntime:      make(map[string]*AppRuntimeInfo)}

	// Image home.
	err = os.MkdirAll(IMAGE_PATH, 0700)
	if err != nil {
		grpclog.Printf("MkdirAll failed %q: %s", IMAGE_PATH, err)
	}

	// Log home.
	err = os.MkdirAll(LOG_PATH, 0700)
	if err != nil {
		grpclog.Printf("MkdirAll failed %q: %s", LOG_PATH, err)
	}

	return s, err
}
