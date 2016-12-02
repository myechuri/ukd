package main

import (
	"github.com/myechuri/ukd/ukdctl/cmd"
	"github.com/spf13/cobra"
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
)

func main() {

	rootCmd.PersistentFlags().StringVar(&serverAddress, "server-endpoint", defaultServer, "server IP and Port ('ip:port') to connect to")

	rootCmd.AddCommand(cmd.VersionCommand(),
		cmd.StartCommand(),
		cmd.StopCommand(),
		cmd.UpdateImageCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
