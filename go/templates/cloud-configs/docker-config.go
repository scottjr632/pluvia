package cloudconfigs

type DockerCloudConfig struct{}

type DockerConfigType string

const (
	DockerConfigTypeUbuntu DockerConfigType = "ubuntu"
	DockerConfigTypeAmzLnx DockerConfigType = "amzlnx"
)

func WithDockerCloudConfig(name DockerConfigType) CloudConfigPiece {
	if name == DockerConfigTypeUbuntu {
		return &DockerCloudConfigUbuntu{}
	}
	if name == DockerConfigTypeAmzLnx {
		return &DockerCloudConfigAmzLnx{}
	}
	panic("Unknown docker config type: " + name)
}
