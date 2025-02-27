// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_upgradeapi.proto

// Site upgrade custom API
//
// x-displayName: "Site"
// These API provide a way to upgrade site's volterra software and OS

package site

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Upgrade SW request
//
// x-displayName: "Upgrade SW Request"
// Request for Upgrade SW of Site
type UpgradeSWRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Site namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-required
	// x-example: "ce398"
	// Site name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// SW/OS Version
	//
	// x-displayName: "Version"
	// x-required
	// x-example: "crt-20201010-600"
	// Version to upgraded to
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *UpgradeSWRequest) Reset()      { *m = UpgradeSWRequest{} }
func (*UpgradeSWRequest) ProtoMessage() {}
func (*UpgradeSWRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b988dc1927df83d3, []int{0}
}
func (m *UpgradeSWRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeSWRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeSWRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeSWRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeSWRequest.Merge(m, src)
}
func (m *UpgradeSWRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeSWRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeSWRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeSWRequest proto.InternalMessageInfo

func (m *UpgradeSWRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpgradeSWRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpgradeSWRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Upgrade SW Response
//
// x-displayName: "Upgrade SW Response"
// Response for SW Upgrade Request, empty because the only returned information
// is currently error message
type UpgradeSWResponse struct {
}

func (m *UpgradeSWResponse) Reset()      { *m = UpgradeSWResponse{} }
func (*UpgradeSWResponse) ProtoMessage() {}
func (*UpgradeSWResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b988dc1927df83d3, []int{1}
}
func (m *UpgradeSWResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeSWResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeSWResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeSWResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeSWResponse.Merge(m, src)
}
func (m *UpgradeSWResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeSWResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeSWResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeSWResponse proto.InternalMessageInfo

// Upgrade OS request
//
// x-displayName: "Upgrade OS Request"
// Request for Upgrade OS of Site
type UpgradeOSRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Site namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-required
	// x-example: "ce398"
	// Site name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// SW/OS Version
	//
	// x-displayName: "Version"
	// x-required
	// x-example: "7.2003.20"
	// Version to upgraded to
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *UpgradeOSRequest) Reset()      { *m = UpgradeOSRequest{} }
func (*UpgradeOSRequest) ProtoMessage() {}
func (*UpgradeOSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b988dc1927df83d3, []int{2}
}
func (m *UpgradeOSRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeOSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeOSRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeOSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeOSRequest.Merge(m, src)
}
func (m *UpgradeOSRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeOSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeOSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeOSRequest proto.InternalMessageInfo

func (m *UpgradeOSRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpgradeOSRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpgradeOSRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Upgrade OS Response
//
// x-displayName: "Upgrade OS Response"
// Response for OS Upgrade Request, empty because the only returned information
// is currently error message
type UpgradeOSResponse struct {
}

func (m *UpgradeOSResponse) Reset()      { *m = UpgradeOSResponse{} }
func (*UpgradeOSResponse) ProtoMessage() {}
func (*UpgradeOSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b988dc1927df83d3, []int{3}
}
func (m *UpgradeOSResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeOSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeOSResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeOSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeOSResponse.Merge(m, src)
}
func (m *UpgradeOSResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeOSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeOSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeOSResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpgradeSWRequest)(nil), "ves.io.schema.site.UpgradeSWRequest")
	golang_proto.RegisterType((*UpgradeSWRequest)(nil), "ves.io.schema.site.UpgradeSWRequest")
	proto.RegisterType((*UpgradeSWResponse)(nil), "ves.io.schema.site.UpgradeSWResponse")
	golang_proto.RegisterType((*UpgradeSWResponse)(nil), "ves.io.schema.site.UpgradeSWResponse")
	proto.RegisterType((*UpgradeOSRequest)(nil), "ves.io.schema.site.UpgradeOSRequest")
	golang_proto.RegisterType((*UpgradeOSRequest)(nil), "ves.io.schema.site.UpgradeOSRequest")
	proto.RegisterType((*UpgradeOSResponse)(nil), "ves.io.schema.site.UpgradeOSResponse")
	golang_proto.RegisterType((*UpgradeOSResponse)(nil), "ves.io.schema.site.UpgradeOSResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/site/public_upgradeapi.proto", fileDescriptor_b988dc1927df83d3)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_upgradeapi.proto", fileDescriptor_b988dc1927df83d3)
}

var fileDescriptor_b988dc1927df83d3 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xb1, 0x6e, 0x13, 0x41,
	0x10, 0xbd, 0x8d, 0x11, 0xc8, 0x57, 0x85, 0x83, 0xe2, 0x62, 0xa2, 0x15, 0xb2, 0x40, 0x20, 0x4b,
	0x77, 0x2b, 0x81, 0x44, 0x11, 0x2a, 0xa8, 0xa0, 0x32, 0x8a, 0x85, 0x10, 0x14, 0x44, 0x7b, 0xce,
	0x64, 0xb3, 0x60, 0xdf, 0x2c, 0xbb, 0x7b, 0x17, 0x10, 0x8a, 0x84, 0xfc, 0x05, 0x48, 0x7c, 0x01,
	0x1d, 0x3d, 0x65, 0x9a, 0x74, 0x50, 0x21, 0x0b, 0x9a, 0x94, 0xf8, 0x4c, 0x41, 0x99, 0x4f, 0x40,
	0x5e, 0x9f, 0xed, 0x33, 0x44, 0xa1, 0x4a, 0x37, 0x33, 0xef, 0xe9, 0xcd, 0x9b, 0xd9, 0x1d, 0xbf,
	0x95, 0x83, 0x89, 0x25, 0x32, 0xd3, 0xdd, 0x85, 0x3e, 0x67, 0x46, 0x5a, 0x60, 0x2a, 0x4b, 0x7a,
	0xb2, 0xbb, 0x95, 0x29, 0xa1, 0xf9, 0x36, 0x70, 0x25, 0x63, 0xa5, 0xd1, 0x62, 0x10, 0x4c, 0xb9,
	0xf1, 0x94, 0x1b, 0x4f, 0xb8, 0x8d, 0x48, 0x48, 0xbb, 0x9b, 0x25, 0x71, 0x17, 0xfb, 0x4c, 0xa0,
	0x40, 0xe6, 0xa8, 0x49, 0xb6, 0xe3, 0x32, 0x97, 0xb8, 0x68, 0x2a, 0xd1, 0x58, 0x17, 0x88, 0xa2,
	0x07, 0x8c, 0x2b, 0xc9, 0x78, 0x9a, 0xa2, 0xe5, 0x56, 0x62, 0x6a, 0x4a, 0xf4, 0xca, 0xb2, 0x19,
	0x54, 0x55, 0x90, 0x9e, 0xe0, 0xd4, 0xbe, 0x51, 0x30, 0xc3, 0xd7, 0x96, 0xf1, 0x2a, 0xd4, 0x5c,
	0x86, 0x72, 0x30, 0x90, 0xe6, 0xcb, 0xf2, 0xcd, 0xe7, 0xfe, 0xea, 0xe3, 0xe9, 0xc0, 0x9d, 0x27,
	0x9b, 0xf0, 0x2a, 0x03, 0x63, 0x83, 0x75, 0xbf, 0x9e, 0xf2, 0x3e, 0x18, 0xc5, 0xbb, 0x10, 0x92,
	0xab, 0xe4, 0x66, 0x7d, 0x73, 0x51, 0x08, 0x02, 0xff, 0xdc, 0x24, 0x09, 0x57, 0x1c, 0xe0, 0xe2,
	0x20, 0xf4, 0x2f, 0xe4, 0xa0, 0x8d, 0xc4, 0x34, 0xac, 0xb9, 0xf2, 0x2c, 0x6d, 0x5e, 0xf2, 0x2f,
	0x56, 0xf4, 0x8d, 0xc2, 0xd4, 0x40, 0xa5, 0x69, 0xbb, 0x73, 0xb6, 0x4d, 0x27, 0xfa, 0xd3, 0xa6,
	0xb7, 0x3e, 0xd6, 0x7c, 0xbf, 0xac, 0xde, 0x7b, 0xf4, 0x30, 0xf8, 0x4c, 0xfc, 0xfa, 0xdc, 0x59,
	0x70, 0x2d, 0xfe, 0xf7, 0x91, 0xe3, 0xbf, 0x17, 0xd3, 0xb8, 0xfe, 0x1f, 0x56, 0x39, 0xde, 0xd3,
	0xe2, 0x4b, 0x78, 0x39, 0x07, 0x13, 0x49, 0x8c, 0x04, 0xa4, 0xa0, 0x79, 0x2f, 0xda, 0xd3, 0xd2,
	0xc2, 0xe0, 0xc7, 0xaf, 0x0f, 0x2b, 0x77, 0x9b, 0x77, 0xca, 0xaf, 0xc6, 0xe6, 0xf3, 0x19, 0xf6,
	0x76, 0x1e, 0xef, 0xbb, 0x37, 0x2e, 0x2b, 0xfb, 0xac, 0xfc, 0x8e, 0x5b, 0x66, 0x6f, 0x83, 0xb4,
	0xaa, 0xae, 0xdb, 0x9d, 0x53, 0x5d, 0xcf, 0x37, 0x7b, 0xaa, 0xeb, 0xc5, 0x7e, 0xce, 0xc6, 0x35,
	0x9a, 0x0d, 0xd2, 0x6a, 0xdc, 0x38, 0x3c, 0x20, 0xb5, 0xef, 0x07, 0x64, 0xed, 0x04, 0x23, 0xed,
	0xe4, 0x05, 0x74, 0xed, 0xe0, 0x5b, 0xb8, 0xb2, 0x4a, 0xee, 0x0f, 0xc8, 0x70, 0x44, 0xbd, 0xa3,
	0x11, 0xf5, 0x8e, 0x47, 0x94, 0xbc, 0x2b, 0x28, 0xf9, 0x54, 0x50, 0xf2, 0xb5, 0xa0, 0x64, 0x58,
	0x50, 0xf2, 0xb3, 0xa0, 0xe4, 0x77, 0x41, 0xbd, 0xe3, 0x82, 0x92, 0xf7, 0x63, 0xea, 0x1d, 0x8e,
	0x29, 0x19, 0x8e, 0xa9, 0x77, 0x34, 0xa6, 0xde, 0xb3, 0x07, 0x02, 0xd5, 0x4b, 0x11, 0xe7, 0xd8,
	0xb3, 0xa0, 0x35, 0x8f, 0x33, 0xc3, 0x5c, 0xb0, 0x83, 0xba, 0x1f, 0x29, 0x8d, 0xb9, 0xdc, 0x06,
	0x1d, 0xcd, 0x60, 0xa6, 0x12, 0x81, 0x0c, 0x5e, 0xdb, 0xd9, 0x59, 0x2d, 0xae, 0x2b, 0x39, 0xef,
	0x2e, 0xe3, 0xf6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x94, 0x51, 0x0b, 0x24, 0x04, 0x00,
	0x00,
}

func (this *UpgradeSWRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeSWRequest)
	if !ok {
		that2, ok := that.(UpgradeSWRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *UpgradeSWResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeSWResponse)
	if !ok {
		that2, ok := that.(UpgradeSWResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UpgradeOSRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeOSRequest)
	if !ok {
		that2, ok := that.(UpgradeOSRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *UpgradeOSResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeOSResponse)
	if !ok {
		that2, ok := that.(UpgradeOSResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UpgradeSWRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&site.UpgradeSWRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpgradeSWResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&site.UpgradeSWResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpgradeOSRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&site.UpgradeOSRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpgradeOSResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&site.UpgradeOSResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicUpgradeapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpgradeAPIClient is the client API for UpgradeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpgradeAPIClient interface {
	// Upgrade SW
	//
	// x-displayName: "Upgrade SW"
	// Upgrade Site SW version
	UpgradeSW(ctx context.Context, in *UpgradeSWRequest, opts ...grpc.CallOption) (*UpgradeSWResponse, error)
	// Upgrade OS
	//
	// x-displayName: "Upgrade OS"
	// Upgrade Site OS version
	UpgradeOS(ctx context.Context, in *UpgradeOSRequest, opts ...grpc.CallOption) (*UpgradeOSResponse, error)
}

type upgradeAPIClient struct {
	cc *grpc.ClientConn
}

func NewUpgradeAPIClient(cc *grpc.ClientConn) UpgradeAPIClient {
	return &upgradeAPIClient{cc}
}

func (c *upgradeAPIClient) UpgradeSW(ctx context.Context, in *UpgradeSWRequest, opts ...grpc.CallOption) (*UpgradeSWResponse, error) {
	out := new(UpgradeSWResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.site.UpgradeAPI/UpgradeSW", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upgradeAPIClient) UpgradeOS(ctx context.Context, in *UpgradeOSRequest, opts ...grpc.CallOption) (*UpgradeOSResponse, error) {
	out := new(UpgradeOSResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.site.UpgradeAPI/UpgradeOS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpgradeAPIServer is the server API for UpgradeAPI service.
type UpgradeAPIServer interface {
	// Upgrade SW
	//
	// x-displayName: "Upgrade SW"
	// Upgrade Site SW version
	UpgradeSW(context.Context, *UpgradeSWRequest) (*UpgradeSWResponse, error)
	// Upgrade OS
	//
	// x-displayName: "Upgrade OS"
	// Upgrade Site OS version
	UpgradeOS(context.Context, *UpgradeOSRequest) (*UpgradeOSResponse, error)
}

// UnimplementedUpgradeAPIServer can be embedded to have forward compatible implementations.
type UnimplementedUpgradeAPIServer struct {
}

func (*UnimplementedUpgradeAPIServer) UpgradeSW(ctx context.Context, req *UpgradeSWRequest) (*UpgradeSWResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeSW not implemented")
}
func (*UnimplementedUpgradeAPIServer) UpgradeOS(ctx context.Context, req *UpgradeOSRequest) (*UpgradeOSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeOS not implemented")
}

func RegisterUpgradeAPIServer(s *grpc.Server, srv UpgradeAPIServer) {
	s.RegisterService(&_UpgradeAPI_serviceDesc, srv)
}

func _UpgradeAPI_UpgradeSW_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeSWRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpgradeAPIServer).UpgradeSW(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UpgradeAPI/UpgradeSW",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpgradeAPIServer).UpgradeSW(ctx, req.(*UpgradeSWRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpgradeAPI_UpgradeOS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeOSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpgradeAPIServer).UpgradeOS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UpgradeAPI/UpgradeOS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpgradeAPIServer).UpgradeOS(ctx, req.(*UpgradeOSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpgradeAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.UpgradeAPI",
	HandlerType: (*UpgradeAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeSW",
			Handler:    _UpgradeAPI_UpgradeSW_Handler,
		},
		{
			MethodName: "UpgradeOS",
			Handler:    _UpgradeAPI_UpgradeOS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_upgradeapi.proto",
}

func (m *UpgradeSWRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeSWRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeSWRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpgradeSWResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeSWResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeSWResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UpgradeOSRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeOSRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeOSRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicUpgradeapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpgradeOSResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeOSResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeOSResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPublicUpgradeapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicUpgradeapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpgradeSWRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	return n
}

func (m *UpgradeSWResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UpgradeOSRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPublicUpgradeapi(uint64(l))
	}
	return n
}

func (m *UpgradeOSResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicUpgradeapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicUpgradeapi(x uint64) (n int) {
	return sovPublicUpgradeapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UpgradeSWRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpgradeSWRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpgradeSWResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpgradeSWResponse{`,
		`}`,
	}, "")
	return s
}
func (this *UpgradeOSRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpgradeOSRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpgradeOSResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpgradeOSResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicUpgradeapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UpgradeSWRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicUpgradeapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpgradeSWRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeSWRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicUpgradeapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpgradeSWResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicUpgradeapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpgradeSWResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeSWResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicUpgradeapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpgradeOSRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicUpgradeapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpgradeOSRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeOSRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicUpgradeapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpgradeOSResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicUpgradeapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpgradeOSResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeOSResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicUpgradeapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicUpgradeapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPublicUpgradeapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicUpgradeapi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicUpgradeapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPublicUpgradeapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicUpgradeapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicUpgradeapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicUpgradeapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicUpgradeapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicUpgradeapi = fmt.Errorf("proto: unexpected end of group")
)
