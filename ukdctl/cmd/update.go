package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
        "io/ioutil"
)

var (
	baseImagePath string
	deltaFilePath string
)

func updateImage(cmd *cobra.Command, args []string) {
	// TODO: TLS
	serverAddress := cmd.InheritedFlags().Lookup("server-endpoint").Value.String()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUkdClient(conn)

        diff, err := ioutil.ReadFile(deltaFilePath)
	if err != nil {
   	    log.Fatalf("ReadFile: %s, error: %v", deltaFilePath, err)
    	}

	updateImageRequest := &api.UpdateImageRequest{
		Base:     baseImagePath,
		Diff:     diff,
	}
	reply, _ := client.UpdateImage(context.Background(), updateImageRequest)
	log.Printf("Unikernel image update: %t, Info: %s",
		reply.Success, reply.Info)
}

func UpdateImageCommand() *cobra.Command {

	var updateImageCmd = &cobra.Command{
		Use:   "update-image",
		Short: "Update a Unikernel Image",
		Long:  `Update a unikernel image to a new image`,
		Run: func(cmd *cobra.Command, args []string) {
			updateImage(cmd, args)
		},
	}
	updateImageCmd.Flags().StringVar(&baseImagePath, "baseImage", "", "fully qualified path of base image on destination")
	updateImageCmd.Flags().StringVar(&deltaFilePath, "deltaFile", "", "fully qualified path of delta of new image over baseImage")
	return updateImageCmd
}
