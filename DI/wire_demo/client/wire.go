//go:build wireinject
// +build wireinject

package main

import (
	"wireDemo/client/handle"

	"github.com/google/wire"
)

func initApp() (*handle.Options, error) {
	panic(wire.Build(handle.ProviderSet))
}
