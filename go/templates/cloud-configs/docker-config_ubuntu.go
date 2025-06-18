package cloudconfigs

type DockerCloudConfigUbuntu struct{}

func WithDockerCloudConfigUbuntu() *DockerCloudConfigUbuntu {
	return &DockerCloudConfigUbuntu{}
}

func (d *DockerCloudConfigUbuntu) Name() string {
	return "docker"
}

func (d *DockerCloudConfigUbuntu) Packages() []string {
	return []string{"apt-transport-https", "ca-certificates", "curl", "software-properties-common"}
}

func (d *DockerCloudConfigUbuntu) RunCmds() []string {
	return []string{
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
		"sudo apt-get update",
		"sudo apt-get install -y docker-ce docker-ce-cli containerd.io",
		"sudo systemctl start docker",
		"sudo systemctl enable docker",
	}
}

var _ CloudConfigPiece = (*DockerCloudConfigUbuntu)(nil)
