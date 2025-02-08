package databases

import "pupa/deps"

func Provide() deps.Provider {
	return deps.ProvideAll(
		deps.Provide(NewDB),
	)
}
