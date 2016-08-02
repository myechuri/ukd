package main

import (
	"fmt"
	"net"

	"fmt"
	"github.com/myechuri/ukd/server"
	"github.com/myechuri/ukd/server/proto"
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
		err = startUkdServer(c)

		return err
	}
	app.Run(os.Args)
}

func startUKdServer(c *cli.Context) error {
	protocol = c.String("protocol")
	port = c.String("port")

	// TODO: Validate protocol.
	lis, err := net.Listen(protocol, fmt.Sprintf(":%d", *port))
	grpcServer := grpc.NewServer()
	s := server.NewServer()
	pb.RegisterUKdServer(grpcServer, s)

	// TODO: TLS.
	grpcServer.Serve(lis)
}
