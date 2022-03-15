package model

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	db *gorm.DB
	rc *redis.Client
	nc *nats.Conn
}

var ProviderSet = wire.NewSet(NewModel)

func NewModel() (*Model, error) {
	//mysql
	dbType := config.Get("database", "type").String("mysql")
	dbAddr := config.Get("database", "DB").String("none")
	open, err := sql.Open(dbType, dbAddr)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: open}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	//redis
	redisAddr := config.Get("redis", "addr").String("127.0.0.1:6379")
	redisPwd := config.Get("redis", "pwd").String("")
	redisDB := config.Get("redis", "db").Int(0)
	rc := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPwd,
		DB:       redisDB,
	})
	//nats
	var opts = make([]nats.Option, 0)
	var name = "micro-wire"
	opts = append(opts, nats.Name(name))                 //客户端名称
	opts = append(opts, nats.ReconnectWait(time.Second)) //设置等待重连时长
	opts = append(opts, nats.MaxReconnects(-1))          //设置重连最大次数,负数表示无限次
	//处理重连程序
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		logger.Debugf("重新连接:[%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		logger.Error("Exiting, no servers available")
	}))
	nc, err := nats.Connect(config.Get("nats", "addr").String("nats://127.0.0.1:4222"), opts...)
	if err != nil {
		return nil, err
	}
	return &Model{
		db: db,
		rc: rc,
		nc: nc,
	}, nil
}
