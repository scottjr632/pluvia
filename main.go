package main

import (
	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/engine"
	"github.com/pluvia/pluvia/templates/box"
)

func main() {
	b := box.New("cool-box", "ami-0c55b159cbfafe1f0", "t2.micro", box.WithIncludeSSH()).Must()

	b.Attach(box.AttachWithDockerStrategy(
		"./DOCKERFILE",
	))

	ctx := context.New()
	eng := engine.NewWithResult(ctx, "pluvia-demo", "us-east-1").Must()

	eng.RunWithResult(ctx, b).Must()

	eng.AttachWithResult(ctx, b).Must()
}
