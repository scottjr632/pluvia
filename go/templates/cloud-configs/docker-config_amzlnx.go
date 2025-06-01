package cloudconfigs

type DockerCloudConfigAmzLnx struct{}

func (d *DockerCloudConfigAmzLnx) Name() string {
	return "docker"
}

func (d *DockerCloudConfigAmzLnx) Packages() []string {
	return []string{}
}

func (d *DockerCloudConfigAmzLnx) RunCmds() []string {
	return []string{
		"echo 'Hello, World!' > /var/tmp/hello-world.txt",
		"sudo dnf update -y",
		"sudo dnf install docker -y",
		"sudo systemctl enable docker",
		"sudo usermod -aG docker $USER",
	}
}

var _ CloudConfigPiece = (*DockerCloudConfigAmzLnx)(nil)
