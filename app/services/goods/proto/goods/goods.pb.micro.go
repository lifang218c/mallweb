// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: goods.proto

package goods

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GoodsService service

func NewGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "GoodsService.FreshGoodsIndex",
			Path:    []string{"/api/v1/goods/freshGoodsIndex"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "GoodsService.GetGoodsDetail",
			Path:    []string{"/api/v1/goods/goodsDetail/{goods_sku_id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "GoodsService.SearchGoods",
			Path:    []string{"/api/v1/goods/searchGoods"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for GoodsService service

type GoodsService interface {
	// 生鲜模块首页商品分类、轮播图、促销商品展示
	FreshGoodsIndex(ctx context.Context, in *FreshGoodsIndexRequest, opts ...client.CallOption) (*FreshGoodsIndexResponse, error)
	// 获取商品详情
	GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, opts ...client.CallOption) (*GoodsDetailResponse, error)
	// 搜索商品
	SearchGoods(ctx context.Context, in *SearchGoodsRequest, opts ...client.CallOption) (*SearchGoodsResponse, error)
}

type goodsService struct {
	c    client.Client
	name string
}

func NewGoodsService(name string, c client.Client) GoodsService {
	return &goodsService{
		c:    c,
		name: name,
	}
}

func (c *goodsService) FreshGoodsIndex(ctx context.Context, in *FreshGoodsIndexRequest, opts ...client.CallOption) (*FreshGoodsIndexResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.FreshGoodsIndex", in)
	out := new(FreshGoodsIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, opts ...client.CallOption) (*GoodsDetailResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.GetGoodsDetail", in)
	out := new(GoodsDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) SearchGoods(ctx context.Context, in *SearchGoodsRequest, opts ...client.CallOption) (*SearchGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.SearchGoods", in)
	out := new(SearchGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoodsService service

type GoodsServiceHandler interface {
	// 生鲜模块首页商品分类、轮播图、促销商品展示
	FreshGoodsIndex(context.Context, *FreshGoodsIndexRequest, *FreshGoodsIndexResponse) error
	// 获取商品详情
	GetGoodsDetail(context.Context, *GoodsDetailRequest, *GoodsDetailResponse) error
	// 搜索商品
	SearchGoods(context.Context, *SearchGoodsRequest, *SearchGoodsResponse) error
}

func RegisterGoodsServiceHandler(s server.Server, hdlr GoodsServiceHandler, opts ...server.HandlerOption) error {
	type goodsService interface {
		FreshGoodsIndex(ctx context.Context, in *FreshGoodsIndexRequest, out *FreshGoodsIndexResponse) error
		GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, out *GoodsDetailResponse) error
		SearchGoods(ctx context.Context, in *SearchGoodsRequest, out *SearchGoodsResponse) error
	}
	type GoodsService struct {
		goodsService
	}
	h := &goodsServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GoodsService.FreshGoodsIndex",
		Path:    []string{"/api/v1/goods/freshGoodsIndex"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GoodsService.GetGoodsDetail",
		Path:    []string{"/api/v1/goods/goodsDetail/{goods_sku_id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GoodsService.SearchGoods",
		Path:    []string{"/api/v1/goods/searchGoods"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&GoodsService{h}, opts...))
}

type goodsServiceHandler struct {
	GoodsServiceHandler
}

func (h *goodsServiceHandler) FreshGoodsIndex(ctx context.Context, in *FreshGoodsIndexRequest, out *FreshGoodsIndexResponse) error {
	return h.GoodsServiceHandler.FreshGoodsIndex(ctx, in, out)
}

func (h *goodsServiceHandler) GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, out *GoodsDetailResponse) error {
	return h.GoodsServiceHandler.GetGoodsDetail(ctx, in, out)
}

func (h *goodsServiceHandler) SearchGoods(ctx context.Context, in *SearchGoodsRequest, out *SearchGoodsResponse) error {
	return h.GoodsServiceHandler.SearchGoods(ctx, in, out)
}
