package cmd

import (
	"github.com/myechuri/ukd/server/api"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	oldImagePath      string
	baseSignaturePath string
	newImagePath      string
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

	// Step 1: Get signature of the current image on server.
	workDir, _ := ioutil.TempDir("", "ukdctl-update-image-")
	defer os.RemoveAll(workDir)
	var signature []byte
	imageSignatureRequest := &api.ImageSignatureRequest{
		Path: oldImagePath,
	}
	reply, _ := client.GetImageSignature(context.Background(),
		imageSignatureRequest)
	if !reply.Success {
		log.Printf("Failed to get signature for image %s, Info: %s",
			oldImagePath, reply.Info)
		return
	}
	log.Printf("Gathered signature of old image on ukd server")
	signature = reply.Signature
	signatureFile := workDir + "/signatureFile"
	f, err := os.Create(signatureFile)
	err = ioutil.WriteFile(signatureFile, signature, 0777)
	f.Close()

	// Step 2: Compute diff of new image using source signature.
	deltaFilePath := workDir + "/deltaFile"
	cmdName := "rdiff"
	args = []string{"delta", signatureFile, newImagePath, deltaFilePath}
	deltaCmd := exec.Command(cmdName, args...)
	err = deltaCmd.Run()
	if err != nil {
		info := "Failed to compute delta for " + newImagePath + " over " + oldImagePath + ", error: " + err.Error()
		log.Printf("Failed to update image. Info: %s", info)
		return
	}

	diff, err := ioutil.ReadFile(deltaFilePath)
	log.Printf("Calcuated diff of new image over old image: %dKB", len(diff)/1024)
	if err != nil {
		log.Fatalf("ReadFile: %s, error: %v", deltaFilePath, err)
	}

	updateImageRequest := &api.UpdateImageRequest{
		Base:    oldImagePath,
		Basesig: signature,
		Diff:    diff,
	}
	log.Printf("Transmitting diff over..")
	updateReply, _ := client.UpdateImage(context.Background(), updateImageRequest)
	log.Printf("Unikernel image update: %t, Info: %s",
		updateReply.Success, updateReply.Info)
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
	updateImageCmd.Flags().StringVar(&oldImagePath, "oldImage", "", "fully qualified path of the old image on ukd server compute node")
	updateImageCmd.Flags().StringVar(&newImagePath, "newImage", "", "fully qualified path of the new image on the compute node where ukdctl is run")
	// updateImageCmd.Flags().StringVar(&baseSignaturePath, "baseImageSignature", "", "fully qualified path of signature of base image that was used to compute diff")
	// updateImageCmd.Flags().StringVar(&deltaFilePath, "deltaFile", "", "fully qualified path of delta of new image over baseImage")
	return updateImageCmd
}
