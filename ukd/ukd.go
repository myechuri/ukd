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

func main() {
	app := cli.NewApp()
	app.Name = "ukd"
	app.Usage = "ukd usage"
	app.Action = func(c *cli.Context) error {
		// Create and start grpc server for ukd
		startUkdServer(c)
                return nil
	}
	app.Run(os.Args)
}

func startUkdServer(c *cli.Context) error {
        // TODO: Gather protocol and port from command line.
	// protocol := c.String("protocol")
	// port := c.String("port")
	protocol := "tcp"
	port := 55555

	// TODO: Validate protocol.
	lis, _ := net.Listen(protocol, fmt.Sprintf(":%d", port))
	grpcServer := grpc.NewServer()
	s := server.NewServer()
	api.RegisterUkdServer(grpcServer, *s)

	// TODO: TLS.
	grpcServer.Serve(lis)
        return nil
}
