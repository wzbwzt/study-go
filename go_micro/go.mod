module go_micro_demo

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
)
