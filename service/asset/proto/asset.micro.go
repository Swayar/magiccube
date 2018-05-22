// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: asset.proto

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	asset.proto

It has these top-level messages:
	PushTxRequest
	GetFileUploadURLRequest
	GetFileUploadURLResponse
	RegisterFileRequest
	RegisterFileResponse
	QueryUploadedDataResponse
	QueryUploadedData
	QueryUploadedRow
	RegisterRequest
	RegisterResponse
	QueryAllAssetRequest
	QueryPara
	QueryAllAssetResponse
	ModifyRequest
	ModifyResponse
	GetFileUploadStatRequest
	GetFileUploadStatResponse
	GetDownLoadURLRequest
	GetDownLoadURLResponse
	QueryByIDRequest
	GetUserPurchaseAssetListRequest
	GetUserPurchaseAssetListResponse
	QueryPurchaseData
	QueryPurchaseRow
	PreSaleNoticeRequest
	PreSaleNoticeResponse
	QueryRequest
	QueryResponse
	QueryData
	AssetData
	QueryMyNoticeRequest
	QueryMyNoticeResponse
	QueryNoticeData
	QueryNoticeRow
*/
package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Asset service

type AssetClient interface {
	//    rpc GetFileUploadURL (GetFileUploadURLRequest) returns (GetFileUploadURLResponse) {
	//    };
	//    rpc GetFileUploadStat (GetFileUploadStatRequest) returns (GetFileUploadStatResponse) {
	//    };
	RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...client.CallOption) (*RegisterFileResponse, error)
	QueryUploadedData(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryUploadedDataResponse, error)
	//    rpc GetDownLoadURL (GetDownLoadURLRequest) returns (GetDownLoadURLResponse) {
	//    };
	RegisterAsset(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	QueryAsset(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
	PreSaleNotice(ctx context.Context, in *PushTxRequest, opts ...client.CallOption) (*PreSaleNoticeResponse, error)
	QueryMyNotice(ctx context.Context, in *QueryMyNoticeRequest, opts ...client.CallOption) (*QueryMyNoticeResponse, error)
	QueryMyPreSale(ctx context.Context, in *QueryMyNoticeRequest, opts ...client.CallOption) (*QueryMyNoticeResponse, error)
}

type assetClient struct {
	c           client.Client
	serviceName string
}

func NewAssetClient(serviceName string, c client.Client) AssetClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "asset"
	}
	return &assetClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *assetClient) RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...client.CallOption) (*RegisterFileResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.RegisterFile", in)
	out := new(RegisterFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryUploadedData(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryUploadedDataResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryUploadedData", in)
	out := new(QueryUploadedDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) RegisterAsset(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.RegisterAsset", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryAsset(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryAsset", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) PreSaleNotice(ctx context.Context, in *PushTxRequest, opts ...client.CallOption) (*PreSaleNoticeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.PreSaleNotice", in)
	out := new(PreSaleNoticeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryMyNotice(ctx context.Context, in *QueryMyNoticeRequest, opts ...client.CallOption) (*QueryMyNoticeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryMyNotice", in)
	out := new(QueryMyNoticeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryMyPreSale(ctx context.Context, in *QueryMyNoticeRequest, opts ...client.CallOption) (*QueryMyNoticeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryMyPreSale", in)
	out := new(QueryMyNoticeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Asset service

type AssetHandler interface {
	//    rpc GetFileUploadURL (GetFileUploadURLRequest) returns (GetFileUploadURLResponse) {
	//    };
	//    rpc GetFileUploadStat (GetFileUploadStatRequest) returns (GetFileUploadStatResponse) {
	//    };
	RegisterFile(context.Context, *RegisterFileRequest, *RegisterFileResponse) error
	QueryUploadedData(context.Context, *QueryRequest, *QueryUploadedDataResponse) error
	//    rpc GetDownLoadURL (GetDownLoadURLRequest) returns (GetDownLoadURLResponse) {
	//    };
	RegisterAsset(context.Context, *RegisterRequest, *RegisterResponse) error
	QueryAsset(context.Context, *QueryRequest, *QueryResponse) error
	PreSaleNotice(context.Context, *PushTxRequest, *PreSaleNoticeResponse) error
	QueryMyNotice(context.Context, *QueryMyNoticeRequest, *QueryMyNoticeResponse) error
	QueryMyPreSale(context.Context, *QueryMyNoticeRequest, *QueryMyNoticeResponse) error
}

func RegisterAssetHandler(s server.Server, hdlr AssetHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Asset{hdlr}, opts...))
}

type Asset struct {
	AssetHandler
}

func (h *Asset) RegisterFile(ctx context.Context, in *RegisterFileRequest, out *RegisterFileResponse) error {
	return h.AssetHandler.RegisterFile(ctx, in, out)
}

func (h *Asset) QueryUploadedData(ctx context.Context, in *QueryRequest, out *QueryUploadedDataResponse) error {
	return h.AssetHandler.QueryUploadedData(ctx, in, out)
}

func (h *Asset) RegisterAsset(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AssetHandler.RegisterAsset(ctx, in, out)
}

func (h *Asset) QueryAsset(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.AssetHandler.QueryAsset(ctx, in, out)
}

func (h *Asset) PreSaleNotice(ctx context.Context, in *PushTxRequest, out *PreSaleNoticeResponse) error {
	return h.AssetHandler.PreSaleNotice(ctx, in, out)
}

func (h *Asset) QueryMyNotice(ctx context.Context, in *QueryMyNoticeRequest, out *QueryMyNoticeResponse) error {
	return h.AssetHandler.QueryMyNotice(ctx, in, out)
}

func (h *Asset) QueryMyPreSale(ctx context.Context, in *QueryMyNoticeRequest, out *QueryMyNoticeResponse) error {
	return h.AssetHandler.QueryMyPreSale(ctx, in, out)
}
