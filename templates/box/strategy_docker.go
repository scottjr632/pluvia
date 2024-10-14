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
		return nil
	}
}

func (s *dockerStrategy) Run(ctx context.Context) error {
	s.box.instance.PublicIp.ApplyT(func(value string) {
		cmdErr := exec.Command("DOCKER_HOST=ssh://ec2-user@"+value, "docker", "build").Run()
		if cmdErr != nil {
			panic(cmdErr)
		}
	})
	return nil
}

var _ strategies.Strategy = (*dockerStrategy)(nil)
