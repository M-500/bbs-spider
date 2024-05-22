// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: rpc/search.proto

package search

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchSvrClient is the client API for SearchSvr apps.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchSvrClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchCompany(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchPeople(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchPeopleResponse, error)
	// 同步企业信息的接口
	SyncCompany(ctx context.Context, in *SyncCompanyRequest, opts ...grpc.CallOption) (*SyncCompanyResponse, error)
}

type searchSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchSvrClient(cc grpc.ClientConnInterface) SearchSvrClient {
	return &searchSvrClient{cc}
}

func (c *searchSvrClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/search.SearchSvr/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSvrClient) SearchCompany(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/search.SearchSvr/SearchCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSvrClient) SearchPeople(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchPeopleResponse, error) {
	out := new(SearchPeopleResponse)
	err := c.cc.Invoke(ctx, "/search.SearchSvr/SearchPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSvrClient) SyncCompany(ctx context.Context, in *SyncCompanyRequest, opts ...grpc.CallOption) (*SyncCompanyResponse, error) {
	out := new(SyncCompanyResponse)
	err := c.cc.Invoke(ctx, "/search.SearchSvr/SyncCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchSvrServer is the server API for SearchSvr apps.
// All implementations must embed UnimplementedSearchSvrServer
// for forward compatibility
type SearchSvrServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	SearchCompany(context.Context, *SearchRequest) (*SearchResponse, error)
	SearchPeople(context.Context, *SearchRequest) (*SearchPeopleResponse, error)
	// 同步企业信息的接口
	SyncCompany(context.Context, *SyncCompanyRequest) (*SyncCompanyResponse, error)
	mustEmbedUnimplementedSearchSvrServer()
}

// UnimplementedSearchSvrServer must be embedded to have forward compatible implementations.
type UnimplementedSearchSvrServer struct {
}

func (UnimplementedSearchSvrServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchSvrServer) SearchCompany(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCompany not implemented")
}
func (UnimplementedSearchSvrServer) SearchPeople(context.Context, *SearchRequest) (*SearchPeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPeople not implemented")
}
func (UnimplementedSearchSvrServer) SyncCompany(context.Context, *SyncCompanyRequest) (*SyncCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncCompany not implemented")
}
func (UnimplementedSearchSvrServer) mustEmbedUnimplementedSearchSvrServer() {}

// UnsafeSearchSvrServer may be embedded to opt out of forward compatibility for this apps.
// Use of this interface is not recommended, as added methods to SearchSvrServer will
// result in compilation errors.
type UnsafeSearchSvrServer interface {
	mustEmbedUnimplementedSearchSvrServer()
}

func RegisterSearchSvrServer(s grpc.ServiceRegistrar, srv SearchSvrServer) {
	s.RegisterService(&SearchSvr_ServiceDesc, srv)
}

func _SearchSvr_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchSvrServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchSvr/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchSvrServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchSvr_SearchCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchSvrServer).SearchCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchSvr/SearchCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchSvrServer).SearchCompany(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchSvr_SearchPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchSvrServer).SearchPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchSvr/SearchPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchSvrServer).SearchPeople(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchSvr_SyncCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchSvrServer).SyncCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchSvr/SyncCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchSvrServer).SyncCompany(ctx, req.(*SyncCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchSvr_ServiceDesc is the grpc.ServiceDesc for SearchSvr apps.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchSvr",
	HandlerType: (*SearchSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchSvr_Search_Handler,
		},
		{
			MethodName: "SearchCompany",
			Handler:    _SearchSvr_SearchCompany_Handler,
		},
		{
			MethodName: "SearchPeople",
			Handler:    _SearchSvr_SearchPeople_Handler,
		},
		{
			MethodName: "SyncCompany",
			Handler:    _SearchSvr_SyncCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/search.proto",
}