package options

type OptionFn[T any] func(T) T

// Apply applies the given options to the given value
func Apply[T any](val T, opts ...OptionFn[T]) {
	for _, opt := range opts {
		opt(val)
	}
}
