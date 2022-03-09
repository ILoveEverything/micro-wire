//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/micro/go-micro/v2"
	"github.com/wu-xie-888/micro-wire/server/handle"
	"github.com/wu-xie-888/micro-wire/server/model"
)

func initApp() (micro.Service, error) {
	panic(wire.Build(handle.ProviderSet, model.ProviderSet))
}
