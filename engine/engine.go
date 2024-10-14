package engine

import (
	"fmt"
	"os"

	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/templates"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Engine struct {
	st auto.Stack
}

func setConfigPassphrase() {
	if os.Getenv("PULUMI_CONFIG_PASSPHRASE") == "" {
		os.Setenv("PULUMI_CONFIG_PASSPHRASE", "your-secure-passphrase")
	}
}

func New(
	ctx context.Context,
	projectName string,
	region string,
) (*Engine, error) {
	setConfigPassphrase()

	localBackendURL := "file://~/.pulumi"
	stackName := "dev"

	project := workspace.Project{
		Name:    tokens.PackageName(projectName),
		Runtime: workspace.NewProjectRuntimeInfo("go", nil),
		Backend: &workspace.ProjectBackend{
			URL: localBackendURL,
		},
	}

	ws, err := auto.NewLocalWorkspace(ctx.Ctx(), auto.Project(project))
	if err != nil {
		fmt.Printf("Failed to create local workspace: %v\n", err)
		return nil, err
	}

	stack, err := auto.UpsertStack(ctx.Ctx(), stackName, ws)
	if err != nil {
		fmt.Printf("Failed to create stack: %v\n", err)
		return nil, err
	}

	err = stack.SetConfig(ctx.Ctx(), "aws:region", auto.ConfigValue{Value: region})
	if err != nil {
		fmt.Printf("Failed to set AWS region: %v\n", err)
		return nil, err
	}

	engine := &Engine{stack}
	return engine, nil
}

func (engine *Engine) Run(ctx context.Context, tmpls ...templates.Template) error {
	engine.st.Workspace().SetProgram(func(pl *pulumi.Context) error {
		for _, t := range tmpls {
			if err := t.Create(&ctx); err != nil {
				return err
			}
		}
		return nil
	})

	res, err := engine.st.Preview(ctx.Ctx())
	if err != nil {
		fmt.Printf("Failed to preview stack: %v\n", err)
		return err
	}

	fmt.Printf("Preview completed! Resources:\n%v\n", res.StdOut)

	return nil
}

func (engine *Engine) Attach(ctx context.Context, strats ...templates.RunAttachable) error {
	for _, t := range strats {
		if err := t.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}
