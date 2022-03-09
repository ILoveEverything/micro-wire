//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wu-xie-888/micro-wire/client/handle"
)

func initApp() (*handle.Options, error) {
	panic(wire.Build(handle.ProviderSet))
}
