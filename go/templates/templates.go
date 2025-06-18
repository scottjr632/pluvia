package templates

import (
	"github.com/pluvia/pluvia/context"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pluvia/pluvia/templates/strategies"
)

type ContextWithPulumi struct {
	context.Context
	PL *pulumi.Context
}

type Attachable[T any] interface {
	RunAttachable
	Attach(strategy strategies.StrategyFn[T]) error
}

type RunAttachable interface {
	Template
	Run(ctx context.Context) error
}

type Template interface {
	Create(ctx *ContextWithPulumi) error
}
