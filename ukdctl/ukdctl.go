package main

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	cliName        = "ukdctl"
	cliDescription = "Command line client for ukd."
	defaultServer  = "localhost:54545"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"ukdctl"},
	}
	serverAddress string // used by rootCmd
        ukName string        // used by startCmd, stopCmd
        imageLocation string // used by startCmd
)

func getServerVersion(client api.UkdClient) {
	versionRequest := &api.VersionRequest{}
	reply, err := client.GetVersion(context.Background(), versionRequest)
	if err != nil {
		log.Fatalf("could not gather grpc server version: %v", err)
	}
	log.Printf("Ukd server version: %d.%d", reply.Major, reply.Minor)

}

func startUK(client api.UkdClient, args []string) {
	startRequest := &api.StartRequest{
		Name:     ukName,
		Location: imageLocation,
	}
	reply, _ := client.StartUK(context.Background(), startRequest)
	log.Printf("Application unikernel started: %t, IP: %s, Reason: %s",
		reply.Success, reply.Ip, reply.Reason)
}

func stopUK(client api.UkdClient) {
	stopRequest := &api.StopRequest{
		Name: "test app",
	}
	reply, _ := client.StopUK(context.Background(), stopRequest)
	log.Printf("Application unikernel stopped: %t, Reason: %s",
		reply.Success, reply.Reason)
}

func main() {

	rootCmd.PersistentFlags().StringVar(&serverAddress, "server-endpoint", defaultServer, "server IP and Port ('ip:port') to connect to")

	// TODO: TLS
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get server version",
		Long:  `Get grpc server version`,
		Run: func(cmd *cobra.Command, args []string) {
			getServerVersion(client)
		},
	}
	rootCmd.AddCommand(versionCmd)

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
	rootCmd.AddCommand(startCmd)

	var stopCommand = &cobra.Command{
		Use:   "stopUK",
		Short: "Stop a Unikernel",
		Long:  `Stop a unikernel with given name`,
		Run: func(cmd *cobra.Command, args []string) {
			stopUK(client)
		},
	}
	rootCmd.AddCommand(stopCommand)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
