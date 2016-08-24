package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	ukName        string
	imageLocation string
	serverAddress string
)

func startUK(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	startRequest := &api.StartRequest{
		Name:     ukName,
		Location: imageLocation,
	}
	reply, _ := client.StartUK(context.Background(), startRequest)
	log.Printf("Application unikernel started: %t, IP: %s, Info: %s",
		reply.Success, reply.Ip, reply.Info)
}

func StartCommand() *cobra.Command {

	var startCmd = &cobra.Command{
		Use:   "startUK [name] [image location]",
		Short: "Start a Unikernel",
		Long:  `Start a unikernel with a given name and image location`,
		Run: func(cmd *cobra.Command, args []string) {
			startUK(cmd, args)
		},
	}
	startCmd.Flags().StringVar(&ukName, "name", "", "name of the application")
	startCmd.Flags().StringVar(&imageLocation, "image-location", "", "location of the application image")
	return startCmd
}
