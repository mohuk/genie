package main

import (
	"log"
	"os"

	"github.com/mohuk/genie/bin/cmd"
	"github.com/urfave/cli"
)

const (
	// Version ...
	Version = "0.1.0"
)

func main() {

	genie := cli.NewApp()
	genie.Name = "genie"
	genie.Description = "automated form generator from databases"
	genie.Version = Version
	genie.Commands = []cli.Command{
		cmd.Serve,
	}
	if err := genie.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
