package main

import (
	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/engine"
	"github.com/pluvia/pluvia/templates/box"
)

func main() {
	ctx := context.New()
	b := box.New(ctx, "cool-box", "ami-067d435ee698a3ff3", "t4g.small", box.WithIncludeSSH()).Must()

	b.Attach(box.AttachWithDockerStrategy(
		"./DOCKERFILE",
	))

	ctx := context.New()
	eng := engine.NewWithResult(ctx, "pluvia-demo", "us-east-1").Must()

	eng.RunWithResult(ctx, b).Must()

	eng.AttachWithResult(ctx, b).Must()
}
