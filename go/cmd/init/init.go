package initcmd

import (
	"context"

	plCtx "github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/engine"
	"github.com/pluvia/pluvia/utils"
	"github.com/urfave/cli/v3"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new Pluvia project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "[required] The name of the project",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "region",
			Aliases: []string{"r"},
			Usage:   "The region to deploy to",
			Value:   "us-east-1",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		name := c.String("name")
		utils.MustCond(len(name) > 0, "Name must be provided")

		region := c.String("region")

		plc := plCtx.New(plCtx.WithContext(ctx))
		eng := engine.NewWithResult(plc, name, region)
		return eng.UnrapErr()
	},
}
