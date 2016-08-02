package main

import (
    "google.golang.org/grpc"
    "github.com/myechuri/ukd/proto"
    "fmt"
    "os"
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "ukd"
  app.Usage = "ukd usage"
  app.Action = func(c *cli.Context) error {
    // Create grpc server for ukd
    return nil
  }

  app.Run(os.Args)
}
