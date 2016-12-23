package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func status(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	statusRequest := &api.StatusRequest{
		Name: ukName,
	}
	reply, _ := client.Status(context.Background(), statusRequest)
	log.Printf("Application unikernel status check: %t, status: %s, Info: %s",
		reply.Success, reply.Status, reply.Info)
}

func StatusCommand() *cobra.Command {

	var statusCommand = &cobra.Command{
		Use:   "status",
		Short: "Status a Unikernel",
		Long:  `Status a unikernel with given name`,
		Run: func(cmd *cobra.Command, args []string) {
			status(cmd, args)
		},
	}
	statusCommand.Flags().StringVar(&ukName, "name", "", "name of the application")
	return statusCommand
}
