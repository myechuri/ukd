package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func stopUK(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	stopRequest := &api.StopRequest{
		Name: "test app",
	}
	reply, _ := client.StopUK(context.Background(), stopRequest)
	log.Printf("Application unikernel stopped: %t, Info: %s",
		reply.Success, reply.Info)
}

func StopCommand() *cobra.Command {

	var stopCommand = &cobra.Command{
		Use:   "stopUK",
		Short: "Stop a Unikernel",
		Long:  `Stop a unikernel with given name`,
		Run: func(cmd *cobra.Command, args []string) {
			stopUK(cmd, args)
		},
	}
	return stopCommand
}
