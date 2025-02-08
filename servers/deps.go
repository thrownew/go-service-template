package servers

import (
	"pupa/deps"

	"go.uber.org/dig"
)

const (
	GRPCServiceGroup = "grpc_service"
)

type (
	serverIn struct {
		dig.In

		Services []GRPCService `group:"grpc_service"`
	}
)

func Provide() deps.Provider {
	return deps.ProvideAll(
		deps.Provide(func(in serverIn) *GRPCServer {
			return NewServer(in.Services...)
		}),
	)
}

func GRPCServiceAdapter[T GRPCService](opts ...dig.ProvideOption) deps.Provider {
	return deps.Provide(func(s T) GRPCService {
		return s

	}, append(opts, dig.Group(GRPCServiceGroup))...)
}
