package main

import (
        "log"
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
        "golang.org/x/net/context"
)

const (
	cliName        = "ukdctl"
	cliDescription = "Command line client for ukd."
)

var (
	rootCmd     = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"ukdctl"},
	}
)

func getServerVersion(client api.UkdClient) {
	versionRequest := &api.VersionRequest{}
	reply, err := client.GetVersion(context.Background(), versionRequest)
	if err != nil {
		log.Fatalf("could not gather grpc server version: %v", err)
	}
	log.Printf("Grpc server version: %d.%d", reply.Major, reply.Minor)

}

func main() {

	// TODO: Gather grpc server address from command line input.
	address := "localhost:55555"

	// TODO: TLS
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	var versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Get server version",
		Long:  `Get grpc server version`,
		Run: func(cmd *cobra.Command, args []string) {
			getServerVersion(client)
		},
	}
	rootCmd.AddCommand(versionCommand)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
