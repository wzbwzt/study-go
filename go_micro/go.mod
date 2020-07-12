module go_micro_demo

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/client/http v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/tcp v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/utp v0.0.0-20200119172437-4fe21aa238fd
	github.com/montanaflynn/stats v0.6.3
	github.com/nats-io/nats.go v1.9.1
)
