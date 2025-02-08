package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"pupa/databases"
	"pupa/deps"
	"pupa/logs"
	"pupa/servers"
)

type (
	Box struct {
		providers []deps.Provider
	}

	BoxOption func(*Box)
)

func BoxProvider(providers ...deps.Provider) BoxOption {
	return func(b *Box) {
		b.providers = append(b.providers, providers...)
	}
}

// NewBox constructor for empty box
func NewBox(opts ...BoxOption) *Box {
	b := &Box{
		providers: []deps.Provider{
			deps.Provide(func() context.Context {
				return context.Background()
			}),
			servers.GRPCClientsProvider("127.0.0.1:8080"),
			logs.Provide(),
			databases.Provide(),
		},
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Run invoke dependencies from container and run function
func (t *Box) Run(f any) error {
	c, err := deps.NewContainer(t.providers...)
	if err != nil {
		return fmt.Errorf("container: %w", err)
	}
	if err = c.Invoke(f); err != nil {
		return fmt.Errorf("invoke: %w", err)
	}
	return nil
}

// RunT invoke dependencies from container and run test function
func (t *Box) RunT(tt *testing.T, name string, f any) {
	tt.Run(name, func(tt *testing.T) {
		c, err := deps.NewContainer(deps.ProvideAll(
			deps.ProvideAll(t.providers...),
			deps.Provide(func() *testing.T { return tt }),
		))
		require.NoError(tt, err)
		require.NoError(tt, c.Invoke(f))
	})
}

// RunB invoke dependencies from container and run benchmarks function
func (t *Box) RunB(tb *testing.B, name string, f any) {
	tb.Run(name, func(tb *testing.B) {
		c, err := deps.NewContainer(deps.ProvideAll(
			deps.ProvideAll(t.providers...),
			deps.Provide(func() *testing.B { return tb }),
		))
		require.NoError(tb, err)
		require.NoError(tb, c.Invoke(f))
	})
}
