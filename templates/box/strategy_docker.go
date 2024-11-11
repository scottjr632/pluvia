package box

import (
	"os/exec"

	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/templates/strategies"
)

type dockerStrategy struct {
	box *Box
}

func AttachWithDockerStrategy(
	imagePath string,
) func(b *Box) strategies.Strategy {
	return func(b *Box) strategies.Strategy {
		if !b.includeSSH {
			panic("Cannot attach docker strategy without SSH")
		}
		return &dockerStrategy{b}
	}
}

func (s *dockerStrategy) Run(ctx context.Context) error {
	s.box.instance.PublicIp.ApplyT(func(value string) string {
		dockerHost := "DOCKER_HOST=ssh://ec2-user@" + value
		ctx.Log().Debug("Building docker image on remote machine with " + dockerHost)

		cmdErr := exec.Command(dockerHost, "docker", "build").Run()
		if cmdErr != nil {
			ctx.Log().Error("Failed to build docker image on remote machine with "+dockerHost, cmdErr.Error())
		}
		return value
	})
	return nil
}

var _ strategies.Strategy = (*dockerStrategy)(nil)
