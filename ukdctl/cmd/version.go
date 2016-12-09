package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func getServerVersion(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	versionRequest := &api.VersionRequest{}
	reply, err := client.GetVersion(context.Background(), versionRequest)
	if err != nil {
		log.Fatalf("could not gather grpc server version: %v", err)
	}
	log.Printf("Ukd server version: %d.%s", reply.Major, reply.Minor)
}

func VersionCommand() *cobra.Command {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get server version",
		Long:  `Get grpc server version`,
		Run: func(cmd *cobra.Command, args []string) {
			getServerVersion(cmd, args)
		},
	}
	return versionCmd
}
