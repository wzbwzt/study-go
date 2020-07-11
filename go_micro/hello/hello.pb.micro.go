// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for HelloInfo service

type HelloInfoService interface {
	Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error)
}

type helloInfoService struct {
	c    client.Client
	name string
}

func NewHelloInfoService(name string, c client.Client) HelloInfoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "helloinfo"
	}
	return &helloInfoService{
		c:    c,
		name: name,
	}
}

func (c *helloInfoService) Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error) {
	req := c.c.NewRequest(c.name, "HelloInfo.Info", in)
	out := new(InfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloInfo service

type HelloInfoHandler interface {
	Info(context.Context, *InfoRequest, *InfoResponse) error
}

func RegisterHelloInfoHandler(s server.Server, hdlr HelloInfoHandler, opts ...server.HandlerOption) error {
	type helloInfo interface {
		Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error
	}
	type HelloInfo struct {
		helloInfo
	}
	h := &helloInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloInfo{h}, opts...))
}

type helloInfoHandler struct {
	HelloInfoHandler
}

func (h *helloInfoHandler) Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error {
	return h.HelloInfoHandler.Info(ctx, in, out)
}
