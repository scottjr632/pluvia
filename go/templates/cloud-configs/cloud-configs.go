package cloudconfigs

import (
	"strings"

	"github.com/pluvia/pluvia/utils"
)

const (
	cloudInit = "#cloud-config"
	indent    = "  "
)

type CloudConfigPiece interface {
	Name() string
	Packages() []string
	RunCmds() []string
}

type CloudConfigBuilder struct {
	pieces []CloudConfigPiece
}

func NewCloudConfigBuilder(initialPieces ...CloudConfigPiece) *CloudConfigBuilder {
	return &CloudConfigBuilder{pieces: initialPieces}
}

func (b *CloudConfigBuilder) Add(piece CloudConfigPiece) {
	b.pieces = append(b.pieces, piece)
}

func (b *CloudConfigBuilder) addPackages(builder *strings.Builder) {
	currPackages := utils.NewSet[string]()

	hasPackages := false
	for _, pc := range b.pieces {
		if len(pc.Packages()) > 0 {
			hasPackages = true
			break
		}
	}

	if !hasPackages {
		return
	}

	builder.WriteString("packages:\n")

	for i, pc := range b.pieces {
		builder.WriteString(indent + "# " + pc.Name() + "\n")
		for _, pkg := range pc.Packages() {
			if currPackages.Has(pkg) {
				continue
			}

			builder.WriteString(indent + "- " + pkg + "\n")
			currPackages.Add(pkg)
		}

		if i != len(b.pieces)-1 {
			builder.WriteString("\n")
		}
	}
}

func (b *CloudConfigBuilder) addRunCmds(builder *strings.Builder) {
	currCmds := utils.NewSet[string]()

	builder.WriteString("runcmd:\n")
	for i, pc := range b.pieces {
		builder.WriteString(indent + "# " + pc.Name() + "\n")
		for _, cmd := range pc.RunCmds() {
			if currCmds.Has(cmd) {
				continue
			}

			builder.WriteString(indent + "- " + cmd + "\n")
			currCmds.Add(cmd)
		}

		if i != len(b.pieces)-1 {
			builder.WriteString("\n")
		}
	}
}

func (b *CloudConfigBuilder) Build() string {
	var builder strings.Builder

	builder.WriteString(cloudInit)
	builder.WriteString("\n")

	b.addPackages(&builder)

	builder.WriteString("\n")

	b.addRunCmds(&builder)

	return builder.String()
}
