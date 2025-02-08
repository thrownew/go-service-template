package deps

import (
	"fmt"

	"go.uber.org/dig"
)

type (
	Provider func(*dependencies)

	dependency struct {
		constructor any
		opts        []dig.ProvideOption
	}

	dependencies struct {
		list []dependency
	}
)

func NewContainer(providers ...Provider) (*dig.Container, error) {
	var d dependencies
	for _, p := range providers {
		p(&d)
	}
	c := dig.New()
	for _, p := range d.list {
		if err := c.Provide(p.constructor, p.opts...); err != nil {
			return nil, fmt.Errorf("provide: %w", err)
		}
	}
	return c, nil
}

func Provide(constructor any, opts ...dig.ProvideOption) Provider {
	return func(d *dependencies) {
		d.list = append(d.list, dependency{
			constructor: constructor,
			opts:        opts,
		})
	}
}

func ProvideAll(providers ...Provider) Provider {
	return func(d *dependencies) {
		for _, p := range providers {
			p(d)
		}
	}
}
