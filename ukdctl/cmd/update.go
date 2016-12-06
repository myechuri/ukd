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

	workDir, _ := ioutil.TempDir("", "ukdctl-update-image-")
	defer os.RemoveAll(workDir)

	// Step 1: Compute signature of new image.
	if _, err := os.Stat(newImagePath); os.IsNotExist(err) {
		log.Printf("%s not found, error: %s", newImagePath, err.Error())
		return
	}
	var newSignature []byte
	newSignatureFile := workDir + "/newSignatureFile"
	f, err := os.Create(newSignatureFile)
	cmdName := "rdiff"
	args = []string{"signature", newImagePath, newSignatureFile}
	rdiffCmd := exec.Command(cmdName, args...)
	err = rdiffCmd.Run()
	if err != nil {
		log.Printf("Failed to compute signature for %s, error: %s", newImagePath, err.Error())
		return
	}
	newSignature, err = ioutil.ReadFile(newSignatureFile)
	log.Printf("Computed new image signature")

	// Step 2: Get signature of the old image on destination.
	var baseSignature []byte
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
	log.Printf("Gathered signature of old image on destination")
	baseSignature = reply.Signature
	signatureFile := workDir + "/oldSignatureFile"
	f, err = os.Create(signatureFile)
	err = ioutil.WriteFile(signatureFile, baseSignature, 0777)
	f.Close()

	// Step 3: Compute diff of new image using source signature.
	deltaFilePath := workDir + "/deltaFile"
	cmdName = "rdiff"
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
		Basesig: baseSignature,
		Newsig:  newSignature,
		Diff:    diff,
	}
	log.Printf("Transmitting diff over..")
	updateReply, _ := client.UpdateImage(context.Background(), updateImageRequest)
	log.Printf("Unikernel image update: %t, new image path on destination: %s, Info: %s",
		updateReply.Success, updateReply.StagedImagePath, updateReply.Info)
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
	return updateImageCmd
}
