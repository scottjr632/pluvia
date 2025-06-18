package main

import (
	"context"
	"log"
	"os"

	"github.com/pluvia/pluvia/cmds/initcmd"
	"github.com/urfave/cli/v3"
)

var cmd = &cli.Command{
	Name:  "pluvia",
	Usage: "A tool for creating and managing infrastructure and services together",
	Commands: []*cli.Command{
		initcmd.InitCmd,
	},
}

func main() {
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
