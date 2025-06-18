/*
Copyright © 2022 theo-commits <me@tlindsey.cloud>
This file is part of CLI application cli-tool
*/
package main

import "pluvia/cmd"

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
