package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	ukName        string // used by startCmd
	imageLocation string // used by startCmd
)

func startUK(client api.UkdClient, args []string) {
	startRequest := &api.StartRequest{
		Name:     ukName,
		Location: imageLocation,
	}
	reply, _ := client.StartUK(context.Background(), startRequest)
	log.Printf("Application unikernel started: %t, IP: %s, Info: %s",
		reply.Success, reply.Ip, reply.Info)
}

func StartCommand(serverAddress string) *cobra.Command {

	// TODO: TLS
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	var startCmd = &cobra.Command{
		Use:   "startUK [name] [image location]",
		Short: "Start a Unikernel",
		Long:  `Start a unikernel with a given name and image location`,
		Run: func(cmd *cobra.Command, args []string) {
			startUK(client, args)
		},
	}
	startCmd.Flags().StringVar(&ukName, "name", "", "name of the application")
	startCmd.Flags().StringVar(&imageLocation, "image-location", "", "location of the application image")
        return startCmd
}
