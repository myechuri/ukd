package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func getLog(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

	logRequest := &api.LogRequest{
		Name: ukName,
	}
	reply, _ := client.GetLog(context.Background(), logRequest)
	log.Printf("Unikernel application log retrived: %t, Info: %s",
		reply.Success, reply.Info)
	log.Printf("Unikernel application log:\n%s", string(reply.LogContent))
}

func LogCommand() *cobra.Command {

	var logCommand = &cobra.Command{
		Use:   "log",
		Short: "Get Log of a Unikernel Application",
		Long:  `Get Log of a unikernel application with given name`,
		Run: func(cmd *cobra.Command, args []string) {
			getLog(cmd, args)
		},
	}
	logCommand.Flags().StringVar(&ukName, "name", "", "name of the application")
	return logCommand
}
