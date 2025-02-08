package servers

import (
	"fmt"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"pupa/deps"
	"pupa/servers/grpc/pupa"

	"go.uber.org/dig"
)

const (
	GRPCServiceGroup = "grpc_service"
)

type (
	servicesIn struct {
		dig.In

		Services []GRPCService `group:"grpc_service"`
	}

	clientsOut struct {
		dig.Out

		DogServiceClient pupa.DogServiceClient
	}
)

func Provide() deps.Provider {
	return deps.ProvideAll(
		deps.Provide(func(in servicesIn) *GRPCServer {
			return NewServer(in.Services...)
		}),
	)
}

func GRPCServiceAdapter[T GRPCService](opts ...dig.ProvideOption) deps.Provider {
	return deps.Provide(func(s T) GRPCService {
		return s

	}, append(opts, dig.Group(GRPCServiceGroup))...)
}

func GRPCClientsProvider(addr string) deps.Provider {
	return deps.Provide(func() (clientsOut, error) {
		conn, err := grpc.NewClient(
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return clientsOut{}, fmt.Errorf("grpc: client: %w", err)
		}
		return clientsOut{
			DogServiceClient: pupa.NewDogServiceClient(conn),
		}, nil
	})
}
