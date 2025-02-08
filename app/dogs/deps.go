package dogs

import (
	"pupa/deps"
	"pupa/servers"
)

func Provide() deps.Provider {
	return deps.ProvideAll(
		deps.Provide(NewServer),
		servers.GRPCServiceAdapter[*Server](),
		deps.Provide(NewRepository),
	)
}
