package cloudconfigs

import (
	"testing"

	"github.com/pluvia/pluvia/utils/testutils"
)

func TestDockerCloudConfig(t *testing.T) {
	config := NewCloudConfigBuilder(WithDockerCloudConfig())
	built := config.Build()

	expected := `#cloud-config
packages:
  # docker
  - apt-transport-https
  - ca-certificates
  - curl
  - software-properties-common

runcmd:
  # docker
  - curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  - sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  - sudo apt-get update
  - sudo apt-get install -y docker-ce docker-ce-cli containerd.io
  - sudo systemctl start docker
  - sudo systemctl enable docker
`

	testutils.AssertStringEquals(t, expected, built)
}
