package engine

import (
	"fmt"
	"os"

	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/result"
	"github.com/pluvia/pluvia/templates"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
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
  
  // This and other settings need to be read-in from an existing config file
	// then fall back to env var, and then error out instead of dictating a default imo.
	// It's just a bit easier for *me* to have this stored separately until then atm
	pluviaRoot := os.Getenv("PWD")
	localBackendURL := "file://" + pluviaRoot
	stackName := "pluvia-demo"

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

func NewWithResult(
	ctx context.Context,
	projectName string,
	region string,
) result.Result[*Engine] {
	res, err := New(ctx, projectName, region)
	return result.NewResult(res, err)
}

func (engine *Engine) Run(ctx context.Context, tmpls ...templates.Template) error {
	engine.st.Workspace().SetProgram(func(pl *pulumi.Context) error {
		// t just makes me think of tests and I hate it here lmao 
		for _, t := range tmpls {
			ctxWithPulumi := templates.ContextWithPulumi{Context: ctx, PL: pl}
			if err := t.Create(&ctxWithPulumi); err != nil {
				return err
			}
		}
		return nil
	})

	// res, err := engine.st.Preview(ctx.Ctx())
	// if err != nil {
	// 	fmt.Printf("Failed to preview stack: %v\n", err)
	// 	return err
	// }

	// fmt.Printf("Preview completed! Resources:\n%v\n", res.StdOut)

	// pulumi.Run(engine.st.Workspace().Program())
	_, err := engine.st.Up(ctx.Ctx(), optup.ProgressStreams(os.Stdout))
	// fmt.Println(res.StdOut)

	return err
}

func (engine *Engine) RunWithResult(ctx context.Context, tmpls ...templates.Template) result.Failable {
	ctx.Log().Debug("Running templates")

	err := engine.Run(ctx, tmpls...)
	if err != nil {
		ctx.Log().Error(err.Error())
	}

	return result.NewFailable(err)
}

func (engine *Engine) Attach(ctx context.Context, strats ...templates.RunAttachable) error {
	ctx.Log().Debug("Running attachments")

	for _, t := range strats {
		if err := t.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (engine *Engine) AttachWithResult(ctx context.Context, strats ...templates.RunAttachable) result.Failable {
	err := engine.Attach(ctx, strats...)
	return result.NewFailable(err)
}

func (engine *Engine) Destroy(ctx context.Context) error {
	_, err := engine.st.Destroy(ctx.Ctx(), optdestroy.ProgressStreams(os.Stdout))
	return err
}

func (engine *Engine) DestroyWithResult(ctx context.Context) result.Failable {
	err := engine.Destroy(ctx)
	return result.NewFailable(err)
}
