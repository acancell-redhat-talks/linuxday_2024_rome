// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/recommendationservice.proto

package hipstershop

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RecommendationService service

func NewRecommendationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RecommendationService service

type RecommendationService interface {
	ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error)
}

type recommendationService struct {
	c    client.Client
	name string
}

func NewRecommendationService(name string, c client.Client) RecommendationService {
	return &recommendationService{
		c:    c,
		name: name,
	}
}

func (c *recommendationService) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error) {
	req := c.c.NewRequest(c.name, "RecommendationService.ListRecommendations", in)
	out := new(ListRecommendationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecommendationService service

type RecommendationServiceHandler interface {
	ListRecommendations(context.Context, *ListRecommendationsRequest, *ListRecommendationsResponse) error
}

func RegisterRecommendationServiceHandler(s server.Server, hdlr RecommendationServiceHandler, opts ...server.HandlerOption) error {
	type recommendationService interface {
		ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, out *ListRecommendationsResponse) error
	}
	type RecommendationService struct {
		recommendationService
	}
	h := &recommendationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RecommendationService{h}, opts...))
}

type recommendationServiceHandler struct {
	RecommendationServiceHandler
}

func (h *recommendationServiceHandler) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, out *ListRecommendationsResponse) error {
	return h.RecommendationServiceHandler.ListRecommendations(ctx, in, out)
}

// Api Endpoints for ProductCatalogService service

func NewProductCatalogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductCatalogService service

type ProductCatalogService interface {
	ListProducts(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error)
}

type productCatalogService struct {
	c    client.Client
	name string
}

func NewProductCatalogService(name string, c client.Client) ProductCatalogService {
	return &productCatalogService{
		c:    c,
		name: name,
	}
}

func (c *productCatalogService) ListProducts(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListProductsResponse, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.ListProducts", in)
	out := new(ListProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.GetProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.SearchProducts", in)
	out := new(SearchProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductCatalogService service

type ProductCatalogServiceHandler interface {
	ListProducts(context.Context, *Empty, *ListProductsResponse) error
	GetProduct(context.Context, *GetProductRequest, *Product) error
	SearchProducts(context.Context, *SearchProductsRequest, *SearchProductsResponse) error
}

func RegisterProductCatalogServiceHandler(s server.Server, hdlr ProductCatalogServiceHandler, opts ...server.HandlerOption) error {
	type productCatalogService interface {
		ListProducts(ctx context.Context, in *Empty, out *ListProductsResponse) error
		GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error
		SearchProducts(ctx context.Context, in *SearchProductsRequest, out *SearchProductsResponse) error
	}
	type ProductCatalogService struct {
		productCatalogService
	}
	h := &productCatalogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductCatalogService{h}, opts...))
}

type productCatalogServiceHandler struct {
	ProductCatalogServiceHandler
}

func (h *productCatalogServiceHandler) ListProducts(ctx context.Context, in *Empty, out *ListProductsResponse) error {
	return h.ProductCatalogServiceHandler.ListProducts(ctx, in, out)
}

func (h *productCatalogServiceHandler) GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error {
	return h.ProductCatalogServiceHandler.GetProduct(ctx, in, out)
}

func (h *productCatalogServiceHandler) SearchProducts(ctx context.Context, in *SearchProductsRequest, out *SearchProductsResponse) error {
	return h.ProductCatalogServiceHandler.SearchProducts(ctx, in, out)
}
