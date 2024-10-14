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

	ctx := context.New(nil)
	eng, err := engine.New(*ctx, "my-aws-project", "us-east-1")
	if err != nil {
		panic(err)
	}

	if err := eng.Run(*ctx, b); err != nil {
		panic(err)
	}

	if err := eng.Attach(*ctx, b); err != nil {
		panic(err)
	}
}
