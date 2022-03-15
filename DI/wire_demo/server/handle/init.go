package handle

import (
	"wireDemo/api"
	"wireDemo/server/model"

	"github.com/google/wire"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Handle struct {
	micro.Service
	model *model.Model
}

var ProviderSet = wire.NewSet(NewHandle, NewService)

func NewService(handle *Handle) (micro.Service, error) {
	var err error
	// etcd
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = config.Get("etcd").StringSlice(nil)
	})
	service := micro.NewService(
		micro.Name(config.Get("name", "server").String("micro-wire-server")),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	service.Init()
	err = api.RegisterGreetHandler(service.Server(), handle)
	err = api.RegisterGoodbyeHandler(service.Server(), handle)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return service, nil
}

func NewHandle(model *model.Model) *Handle {
	return &Handle{model: model}
}
