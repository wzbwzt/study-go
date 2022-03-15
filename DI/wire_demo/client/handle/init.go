package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var ProviderSet = wire.NewSet(NewWeb, NewService, NewOptions)

func NewWeb() (*gin.Engine, error) {
	r := gin.Default()
	return r, nil
}

func NewService() (micro.Service, error) {
	var err error
	// etcd
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = config.Get("etcd").StringSlice(nil)
	})
	service := micro.NewService(
		micro.Name(config.Get("name", "client").String("micro-wire-client")),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	service.Init()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return service, nil
}

func NewOptions(r *gin.Engine, s micro.Service) (*Options, error) {
	return &Options{
		Svr: s,
		Web: r,
	}, nil
}
