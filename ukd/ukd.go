package main

import (
	"fmt"
	"net"

	"github.com/myechuri/ukd/server"
	"github.com/myechuri/ukd/server/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"
)

const (
	defaultProtocol = "tcp"
	defaultPort     = 54545
)

var (
	protocol string
	port     int64
)

func main() {
	app := cli.NewApp()
	app.Name = "ukd"
	app.Usage = "Unikernel Node Runtime Server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "protocol",
			Value:       defaultProtocol,
			Usage:       "grpc server protocol",
			Destination: &protocol,
		},
		cli.Int64Flag{
			Name:        "port",
			Value:       defaultPort,
			Usage:       "grpc server port",
			Destination: &port,
		},
	}

	app.Action = func(c *cli.Context) error {
		// Create and start grpc server for ukd
		startUkdServer(protocol, port)
		return nil
	}
	app.Run(os.Args)
}

func startUkdServer(protocol string, port int64) error {
	// TODO: Validate protocol and port.
	lis, _ := net.Listen(protocol, fmt.Sprintf(":%d", port))
	grpcServer := grpc.NewServer()
	s := server.NewServer()
	api.RegisterUkdServer(grpcServer, *s)

	// TODO: TLS.
	grpcServer.Serve(lis)
	return nil
}
