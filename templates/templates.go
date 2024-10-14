package templates

import (
	"github.com/pluvia/pluvia/context"

	"github.com/pluvia/pluvia/templates/strategies"
)

type Attachable[T any] interface {
	RunAttachable
	Attach(strategy strategies.StrategyFn[T]) error
}

type RunAttachable interface {
	Template
	Run(ctx context.Context) error
}

type Template interface {
	Create(ctx *context.Context) error
}
