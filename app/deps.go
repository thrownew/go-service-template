package app

import (
	"pupa/app/dogs"
	"pupa/deps"
)

func Provide() deps.Provider {
	return deps.ProvideAll(
		dogs.Provide(),
	)
}
