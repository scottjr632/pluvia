package strategies

import "github.com/pluvia/pluvia/context"

type StrategyFn[T any] func(T) Strategy

type Strategy interface {
	Run(ctx context.Context) error
}
