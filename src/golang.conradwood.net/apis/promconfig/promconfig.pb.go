// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/promconfig/promconfig.proto
// DO NOT EDIT!

/*
Package promconfig is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/promconfig/promconfig.proto

It has these top-level messages:
	Target
	Reporter
	TargetList
	PercentAlert
	SeriesMatch
	Series
	SeriesList
	PercentAlertList
*/
package promconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import htmlserver "golang.conradwood.net/apis/htmlserver"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type AlertEffects int32

const (
	AlertEffects_NOBODY AlertEffects = 0
	AlertEffects_USERS  AlertEffects = 1
	AlertEffects_SYSOP  AlertEffects = 2
)

var AlertEffects_name = map[int32]string{
	0: "NOBODY",
	1: "USERS",
	2: "SYSOP",
}
var AlertEffects_value = map[string]int32{
	"NOBODY": 0,
	"USERS":  1,
	"SYSOP":  2,
}

func (x AlertEffects) String() string {
	return proto.EnumName(AlertEffects_name, int32(x))
}
func (AlertEffects) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Target struct {
	Name      string    `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Addresses []string  `protobuf:"bytes,2,rep,name=Addresses" json:"Addresses,omitempty"`
	Reporter  *Reporter `protobuf:"bytes,3,opt,name=Reporter" json:"Reporter,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Target) GetReporter() *Reporter {
	if m != nil {
		return m.Reporter
	}
	return nil
}

type Reporter struct {
	Reporter string `protobuf:"bytes,1,opt,name=Reporter" json:"Reporter,omitempty"`
}

func (m *Reporter) Reset()                    { *m = Reporter{} }
func (m *Reporter) String() string            { return proto.CompactTextString(m) }
func (*Reporter) ProtoMessage()               {}
func (*Reporter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reporter) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

type TargetList struct {
	Reporter *Reporter `protobuf:"bytes,1,opt,name=Reporter" json:"Reporter,omitempty"`
	Targets  []*Target `protobuf:"bytes,2,rep,name=Targets" json:"Targets,omitempty"`
}

func (m *TargetList) Reset()                    { *m = TargetList{} }
func (m *TargetList) String() string            { return proto.CompactTextString(m) }
func (*TargetList) ProtoMessage()               {}
func (*TargetList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TargetList) GetReporter() *Reporter {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func (m *TargetList) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

type PercentAlert struct {
	ID          uint64       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	TotalMetric string       `protobuf:"bytes,2,opt,name=TotalMetric" json:"TotalMetric,omitempty"`
	CountMetric string       `protobuf:"bytes,3,opt,name=CountMetric" json:"CountMetric,omitempty"`
	Effects     AlertEffects `protobuf:"varint,4,opt,name=Effects,enum=promconfig.AlertEffects" json:"Effects,omitempty"`
}

func (m *PercentAlert) Reset()                    { *m = PercentAlert{} }
func (m *PercentAlert) String() string            { return proto.CompactTextString(m) }
func (*PercentAlert) ProtoMessage()               {}
func (*PercentAlert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PercentAlert) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PercentAlert) GetTotalMetric() string {
	if m != nil {
		return m.TotalMetric
	}
	return ""
}

func (m *PercentAlert) GetCountMetric() string {
	if m != nil {
		return m.CountMetric
	}
	return ""
}

func (m *PercentAlert) GetEffects() AlertEffects {
	if m != nil {
		return m.Effects
	}
	return AlertEffects_NOBODY
}

type SeriesMatch struct {
	Prefix       []string `protobuf:"bytes,1,rep,name=Prefix" json:"Prefix,omitempty"`
	PartialMatch bool     `protobuf:"varint,2,opt,name=PartialMatch" json:"PartialMatch,omitempty"`
}

func (m *SeriesMatch) Reset()                    { *m = SeriesMatch{} }
func (m *SeriesMatch) String() string            { return proto.CompactTextString(m) }
func (*SeriesMatch) ProtoMessage()               {}
func (*SeriesMatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SeriesMatch) GetPrefix() []string {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *SeriesMatch) GetPartialMatch() bool {
	if m != nil {
		return m.PartialMatch
	}
	return false
}

type Series struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *Series) Reset()                    { *m = Series{} }
func (m *Series) String() string            { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()               {}
func (*Series) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Series) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SeriesList struct {
	Series []*Series `protobuf:"bytes,1,rep,name=Series" json:"Series,omitempty"`
}

func (m *SeriesList) Reset()                    { *m = SeriesList{} }
func (m *SeriesList) String() string            { return proto.CompactTextString(m) }
func (*SeriesList) ProtoMessage()               {}
func (*SeriesList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SeriesList) GetSeries() []*Series {
	if m != nil {
		return m.Series
	}
	return nil
}

type PercentAlertList struct {
	Alerts []*PercentAlert `protobuf:"bytes,1,rep,name=Alerts" json:"Alerts,omitempty"`
}

func (m *PercentAlertList) Reset()                    { *m = PercentAlertList{} }
func (m *PercentAlertList) String() string            { return proto.CompactTextString(m) }
func (*PercentAlertList) ProtoMessage()               {}
func (*PercentAlertList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PercentAlertList) GetAlerts() []*PercentAlert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

func init() {
	proto.RegisterType((*Target)(nil), "promconfig.Target")
	proto.RegisterType((*Reporter)(nil), "promconfig.Reporter")
	proto.RegisterType((*TargetList)(nil), "promconfig.TargetList")
	proto.RegisterType((*PercentAlert)(nil), "promconfig.PercentAlert")
	proto.RegisterType((*SeriesMatch)(nil), "promconfig.SeriesMatch")
	proto.RegisterType((*Series)(nil), "promconfig.Series")
	proto.RegisterType((*SeriesList)(nil), "promconfig.SeriesList")
	proto.RegisterType((*PercentAlertList)(nil), "promconfig.PercentAlertList")
	proto.RegisterEnum("promconfig.AlertEffects", AlertEffects_name, AlertEffects_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PromConfigService service

type PromConfigServiceClient interface {
	// query registry for new targets
	QueryForTargets(ctx context.Context, in *Reporter, opts ...grpc.CallOption) (*TargetList, error)
	// submit new targes, call this when we want to replace a bunch of new targets
	NewTargets(ctx context.Context, in *TargetList, opts ...grpc.CallOption) (*common.Void, error)
	// find series by partial name match
	FindSeries(ctx context.Context, in *SeriesMatch, opts ...grpc.CallOption) (*SeriesList, error)
	// get a list of all metrics
	GetSeries(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*SeriesList, error)
	// save or create a simple percentage-based alert. If ID is set it will overwrite existing one
	UpdatePercentageAlert(ctx context.Context, in *PercentAlert, opts ...grpc.CallOption) (*PercentAlert, error)
	// get all percentage alert
	GetAllPercentageAlerts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PercentAlertList, error)
	// renderer/editor
	HTMLRenderer(ctx context.Context, in *htmlserver.SnippetRequest, opts ...grpc.CallOption) (*htmlserver.SnippetResponse, error)
}

type promConfigServiceClient struct {
	cc *grpc.ClientConn
}

func NewPromConfigServiceClient(cc *grpc.ClientConn) PromConfigServiceClient {
	return &promConfigServiceClient{cc}
}

func (c *promConfigServiceClient) QueryForTargets(ctx context.Context, in *Reporter, opts ...grpc.CallOption) (*TargetList, error) {
	out := new(TargetList)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/QueryForTargets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) NewTargets(ctx context.Context, in *TargetList, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/NewTargets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) FindSeries(ctx context.Context, in *SeriesMatch, opts ...grpc.CallOption) (*SeriesList, error) {
	out := new(SeriesList)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/FindSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) GetSeries(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*SeriesList, error) {
	out := new(SeriesList)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/GetSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) UpdatePercentageAlert(ctx context.Context, in *PercentAlert, opts ...grpc.CallOption) (*PercentAlert, error) {
	out := new(PercentAlert)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/UpdatePercentageAlert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) GetAllPercentageAlerts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PercentAlertList, error) {
	out := new(PercentAlertList)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/GetAllPercentageAlerts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promConfigServiceClient) HTMLRenderer(ctx context.Context, in *htmlserver.SnippetRequest, opts ...grpc.CallOption) (*htmlserver.SnippetResponse, error) {
	out := new(htmlserver.SnippetResponse)
	err := grpc.Invoke(ctx, "/promconfig.PromConfigService/HTMLRenderer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PromConfigService service

type PromConfigServiceServer interface {
	// query registry for new targets
	QueryForTargets(context.Context, *Reporter) (*TargetList, error)
	// submit new targes, call this when we want to replace a bunch of new targets
	NewTargets(context.Context, *TargetList) (*common.Void, error)
	// find series by partial name match
	FindSeries(context.Context, *SeriesMatch) (*SeriesList, error)
	// get a list of all metrics
	GetSeries(context.Context, *common.Void) (*SeriesList, error)
	// save or create a simple percentage-based alert. If ID is set it will overwrite existing one
	UpdatePercentageAlert(context.Context, *PercentAlert) (*PercentAlert, error)
	// get all percentage alert
	GetAllPercentageAlerts(context.Context, *common.Void) (*PercentAlertList, error)
	// renderer/editor
	HTMLRenderer(context.Context, *htmlserver.SnippetRequest) (*htmlserver.SnippetResponse, error)
}

func RegisterPromConfigServiceServer(s *grpc.Server, srv PromConfigServiceServer) {
	s.RegisterService(&_PromConfigService_serviceDesc, srv)
}

func _PromConfigService_QueryForTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reporter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).QueryForTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/QueryForTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).QueryForTargets(ctx, req.(*Reporter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_NewTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).NewTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/NewTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).NewTargets(ctx, req.(*TargetList))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_FindSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeriesMatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).FindSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/FindSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).FindSeries(ctx, req.(*SeriesMatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_GetSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).GetSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/GetSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).GetSeries(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_UpdatePercentageAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PercentAlert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).UpdatePercentageAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/UpdatePercentageAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).UpdatePercentageAlert(ctx, req.(*PercentAlert))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_GetAllPercentageAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).GetAllPercentageAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/GetAllPercentageAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).GetAllPercentageAlerts(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromConfigService_HTMLRenderer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(htmlserver.SnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromConfigServiceServer).HTMLRenderer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promconfig.PromConfigService/HTMLRenderer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromConfigServiceServer).HTMLRenderer(ctx, req.(*htmlserver.SnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PromConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "promconfig.PromConfigService",
	HandlerType: (*PromConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryForTargets",
			Handler:    _PromConfigService_QueryForTargets_Handler,
		},
		{
			MethodName: "NewTargets",
			Handler:    _PromConfigService_NewTargets_Handler,
		},
		{
			MethodName: "FindSeries",
			Handler:    _PromConfigService_FindSeries_Handler,
		},
		{
			MethodName: "GetSeries",
			Handler:    _PromConfigService_GetSeries_Handler,
		},
		{
			MethodName: "UpdatePercentageAlert",
			Handler:    _PromConfigService_UpdatePercentageAlert_Handler,
		},
		{
			MethodName: "GetAllPercentageAlerts",
			Handler:    _PromConfigService_GetAllPercentageAlerts_Handler,
		},
		{
			MethodName: "HTMLRenderer",
			Handler:    _PromConfigService_HTMLRenderer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/promconfig/promconfig.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/promconfig/promconfig.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x4f, 0xdb, 0x4c,
	0x10, 0xfd, 0x9c, 0xf0, 0x19, 0x32, 0x89, 0x68, 0xba, 0x6a, 0xa9, 0xeb, 0x72, 0x88, 0x2c, 0x51,
	0x45, 0xa8, 0x32, 0x34, 0x48, 0x55, 0x2f, 0xb4, 0x02, 0x02, 0x14, 0x15, 0x42, 0xba, 0x86, 0x4a,
	0xf4, 0xe6, 0xda, 0x93, 0x60, 0xc9, 0xd9, 0x75, 0xd7, 0x1b, 0x68, 0xaf, 0xfd, 0x19, 0x3d, 0xf6,
	0x47, 0xf5, 0xf7, 0x54, 0x59, 0xaf, 0x9b, 0x0d, 0x10, 0xc4, 0xc9, 0xeb, 0x99, 0xf7, 0x66, 0x9e,
	0xdf, 0xac, 0x07, 0xb6, 0x33, 0xc1, 0x25, 0xcf, 0x37, 0x86, 0x3c, 0x0d, 0xd9, 0xd0, 0x8f, 0x38,
	0x13, 0x61, 0x7c, 0xcd, 0x79, 0xec, 0x33, 0x94, 0x1b, 0x61, 0x96, 0xe4, 0x1b, 0x99, 0xe0, 0xa3,
	0x88, 0xb3, 0x41, 0x32, 0x34, 0x8e, 0xbe, 0xe2, 0x11, 0x98, 0x46, 0x5c, 0xff, 0x9e, 0x1a, 0x11,
	0x1f, 0x8d, 0x38, 0xd3, 0x8f, 0x82, 0xeb, 0xbe, 0xb9, 0x07, 0x7f, 0x29, 0x47, 0x69, 0x8e, 0xe2,
	0x0a, 0x85, 0x71, 0x2c, 0x78, 0x5e, 0x0a, 0xf6, 0x59, 0x28, 0x86, 0x28, 0x09, 0x81, 0x85, 0x5e,
	0x38, 0x42, 0xc7, 0x6a, 0x59, 0xed, 0x1a, 0x55, 0x67, 0xb2, 0x0a, 0xb5, 0x9d, 0x38, 0x16, 0x98,
	0xe7, 0x98, 0x3b, 0x95, 0x56, 0xb5, 0x5d, 0xa3, 0xd3, 0x00, 0xd9, 0x84, 0x25, 0x8a, 0x19, 0x17,
	0x12, 0x85, 0x53, 0x6d, 0x59, 0xed, 0x7a, 0xe7, 0x89, 0x6f, 0x7c, 0x54, 0x99, 0xa3, 0xff, 0x50,
	0xde, 0xcb, 0x29, 0x83, 0xb8, 0x06, 0xbb, 0xe8, 0x39, 0xc5, 0xa5, 0x00, 0x85, 0xaa, 0xe3, 0x24,
	0x97, 0x33, 0x7d, 0xac, 0x87, 0xf4, 0x21, 0xaf, 0x60, 0xb1, 0xe0, 0x17, 0xaa, 0xeb, 0x1d, 0x62,
	0x12, 0x8a, 0x14, 0x2d, 0x21, 0xde, 0x6f, 0x0b, 0x1a, 0x7d, 0x14, 0x11, 0x32, 0xb9, 0x93, 0xa2,
	0x90, 0x64, 0x19, 0x2a, 0x47, 0x5d, 0xd5, 0x6a, 0x81, 0x56, 0x8e, 0xba, 0xa4, 0x05, 0xf5, 0x33,
	0x2e, 0xc3, 0xf4, 0x04, 0xa5, 0x48, 0x22, 0xa7, 0xa2, 0xd4, 0x9a, 0xa1, 0x09, 0x62, 0x8f, 0x8f,
	0x99, 0xd4, 0x88, 0x6a, 0x81, 0x30, 0x42, 0xe4, 0x1d, 0x2c, 0xee, 0x0f, 0x06, 0x18, 0xc9, 0xdc,
	0x59, 0x68, 0x59, 0xed, 0xe5, 0x8e, 0x63, 0x4a, 0x52, 0x7d, 0x75, 0x7e, 0x17, 0x7e, 0xfd, 0x7c,
	0x6e, 0x8f, 0x13, 0x26, 0xb7, 0x3a, 0xb4, 0x24, 0x79, 0x47, 0x50, 0x0f, 0x50, 0x24, 0x98, 0x9f,
	0x84, 0x32, 0xba, 0x24, 0x2b, 0x60, 0xf7, 0x05, 0x0e, 0x92, 0xef, 0x8e, 0xa5, 0xc6, 0xa2, 0xdf,
	0x88, 0x07, 0x8d, 0x7e, 0x28, 0x64, 0x12, 0xa6, 0x0a, 0xa7, 0xb4, 0x2e, 0xd1, 0x99, 0x98, 0xb7,
	0x0a, 0x76, 0x51, 0xea, 0xae, 0x99, 0x7b, 0x6f, 0x01, 0x8a, 0xac, 0xf2, 0x7e, 0xbd, 0xc4, 0xaa,
	0x3e, 0x37, 0x8c, 0x2c, 0x32, 0x54, 0x23, 0xbc, 0x2e, 0x34, 0x4d, 0x1b, 0xf5, 0xec, 0x6c, 0xf5,
	0x52, 0xf2, 0x67, 0xbe, 0xda, 0x44, 0x53, 0x8d, 0x5b, 0xdf, 0x84, 0x86, 0xe9, 0x06, 0x01, 0xb0,
	0x7b, 0xa7, 0xbb, 0xa7, 0xdd, 0x8b, 0xe6, 0x7f, 0xa4, 0x06, 0xff, 0x9f, 0x07, 0xfb, 0x34, 0x68,
	0x5a, 0x93, 0x63, 0x70, 0x11, 0x9c, 0xf6, 0x9b, 0x95, 0xce, 0x9f, 0x2a, 0x3c, 0xee, 0x0b, 0x3e,
	0xda, 0x53, 0x55, 0x03, 0x14, 0x57, 0x49, 0x84, 0xe4, 0x3d, 0x3c, 0xfa, 0x34, 0x46, 0xf1, 0xe3,
	0x80, 0x0b, 0x3d, 0x68, 0x72, 0xe7, 0xb5, 0x71, 0x57, 0x6e, 0xdf, 0x0d, 0x25, 0xbd, 0x03, 0xd0,
	0xc3, 0xeb, 0x92, 0x3b, 0x07, 0xe5, 0x36, 0x7c, 0xfd, 0x1f, 0x7e, 0xe6, 0x49, 0x4c, 0xb6, 0x01,
	0x0e, 0x12, 0x16, 0x6b, 0x7b, 0x9f, 0xdd, 0x36, 0x4b, 0x4d, 0x60, 0xb6, 0xa5, 0xe1, 0xf6, 0x6b,
	0xa8, 0x1d, 0xa2, 0xd4, 0xec, 0x99, 0xca, 0x73, 0x29, 0x1f, 0xe1, 0xe9, 0x79, 0x16, 0x87, 0x12,
	0xb5, 0x99, 0xe1, 0x10, 0x8b, 0x4b, 0x3c, 0xd7, 0x69, 0x77, 0x6e, 0x86, 0x74, 0x61, 0xe5, 0x10,
	0xe5, 0x4e, 0x9a, 0xde, 0x28, 0x76, 0x53, 0xcc, 0xea, 0xbc, 0x0a, 0x4a, 0xd2, 0x21, 0x34, 0x3e,
	0x9c, 0x9d, 0x1c, 0x53, 0x64, 0x31, 0x8a, 0xc9, 0x9f, 0xee, 0x1b, 0x6b, 0x27, 0x60, 0x49, 0x96,
	0xa1, 0xa4, 0xf8, 0x6d, 0x8c, 0xb9, 0x74, 0x5f, 0xdc, 0x99, 0xcb, 0x33, 0xce, 0x72, 0xdc, 0xed,
	0xc1, 0x1a, 0x43, 0x69, 0xee, 0x34, 0xbd, 0xe5, 0x26, 0x6b, 0xcd, 0x90, 0xf0, 0x65, 0xed, 0x41,
	0x1b, 0xf7, 0xab, 0xad, 0x76, 0xde, 0xd6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x5e, 0x76,
	0x7f, 0xa8, 0x05, 0x00, 0x00,
}
