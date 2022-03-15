//go:build wireinject
// +build wireinject

package main

import (
	"wireDemo/server/handle"
	"wireDemo/server/model"

	"github.com/google/wire"
	"github.com/micro/go-micro/v2"
)

func initApp() (micro.Service, error) {
	panic(wire.Build(handle.ProviderSet, model.ProviderSet))
}
