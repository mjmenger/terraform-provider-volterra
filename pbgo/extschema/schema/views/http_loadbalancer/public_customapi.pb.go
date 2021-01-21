// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/http_loadbalancer/public_customapi.proto

package http_loadbalancer

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
import ves_io_schema_virtual_host_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Get DNS info Request
//
// x-displayName: "Get DNS Info Request"
// Request message for get-dns-info API
type GetDnsInfoRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace for the HTTP loadbalancer
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "value"
	// Name of the HTTP loadbalancer
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetDnsInfoRequest) Reset()      { *m = GetDnsInfoRequest{} }
func (*GetDnsInfoRequest) ProtoMessage() {}
func (*GetDnsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{0}
}

func (m *GetDnsInfoRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetDnsInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetDnsInfoResponse
//
// x-displayName: "Get DNS Info Response"
// Response for get-dns-info API
type GetDnsInfoResponse struct {
	// DNS information
	//
	// x-displayName: "DNS information"
	// DNS information object for this HTTP loadbalancer
	DnsInfo *ves_io_schema_virtual_host_dns_info.GlobalSpecType `protobuf:"bytes,1,opt,name=dns_info,json=dnsInfo" json:"dns_info,omitempty"`
}

func (m *GetDnsInfoResponse) Reset()      { *m = GetDnsInfoResponse{} }
func (*GetDnsInfoResponse) ProtoMessage() {}
func (*GetDnsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{1}
}

func (m *GetDnsInfoResponse) GetDnsInfo() *ves_io_schema_virtual_host_dns_info.GlobalSpecType {
	if m != nil {
		return m.DnsInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoRequest")
	golang_proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoRequest")
	proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoResponse")
	golang_proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoResponse")
}
func (this *GetDnsInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoRequest)
	if !ok {
		that2, ok := that.(GetDnsInfoRequest)
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
	return true
}
func (this *GetDnsInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoResponse)
	if !ok {
		that2, ok := that.(GetDnsInfoResponse)
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
	if !this.DnsInfo.Equal(that1.DnsInfo) {
		return false
	}
	return true
}
func (this *GetDnsInfoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&http_loadbalancer.GetDnsInfoRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetDnsInfoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&http_loadbalancer.GetDnsInfoResponse{")
	if this.DnsInfo != nil {
		s = append(s, "DnsInfo: "+fmt.Sprintf("%#v", this.DnsInfo)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
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

// Client API for CustomAPI service

type CustomAPIClient interface {
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given HTTP loadbalancer
	GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error) {
	out := new(GetDnsInfoResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.views.http_loadbalancer.CustomAPI/GetDnsInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomAPI service

type CustomAPIServer interface {
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given HTTP loadbalancer
	GetDnsInfo(context.Context, *GetDnsInfoRequest) (*GetDnsInfoResponse, error)
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetDnsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDnsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.http_loadbalancer.CustomAPI/GetDnsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, req.(*GetDnsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.http_loadbalancer.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDnsInfo",
			Handler:    _CustomAPI_GetDnsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/http_loadbalancer/public_customapi.proto",
}

func (m *GetDnsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *GetDnsInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DnsInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.DnsInfo.Size()))
		n1, err := m.DnsInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetDnsInfoRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *GetDnsInfoResponse) Size() (n int) {
	var l int
	_ = l
	if m.DnsInfo != nil {
		l = m.DnsInfo.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetDnsInfoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetDnsInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoResponse{`,
		`DnsInfo:` + strings.Replace(fmt.Sprintf("%v", this.DnsInfo), "GlobalSpecType", "ves_io_schema_virtual_host_dns_info.GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetDnsInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDnsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
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
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *GetDnsInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDnsInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsInfo == nil {
				m.DnsInfo = &ves_io_schema_virtual_host_dns_info.GlobalSpecType{}
			}
			if err := m.DnsInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicCustomapi
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPublicCustomapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPublicCustomapi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPublicCustomapi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPublicCustomapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/views/http_loadbalancer/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/http_loadbalancer/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x8b, 0x14, 0x41,
	0x10, 0xdd, 0x5e, 0x45, 0xbd, 0x31, 0xd1, 0x31, 0x59, 0xd6, 0xa5, 0x91, 0x05, 0xc1, 0xc0, 0xe9,
	0x96, 0xbb, 0x44, 0xc5, 0xc4, 0x8f, 0x63, 0xb9, 0x44, 0x65, 0x35, 0xd2, 0x60, 0xe9, 0x99, 0xa9,
	0x9d, 0x6d, 0x9d, 0xed, 0x6a, 0xbb, 0x7b, 0x46, 0x45, 0x0e, 0xe4, 0x72, 0x41, 0xf0, 0x4f, 0xf8,
	0x07, 0x04, 0xe1, 0x92, 0xcb, 0xbc, 0x48, 0x0e, 0x4d, 0x0c, 0xdd, 0x39, 0x03, 0xc1, 0xe4, 0x7e,
	0x82, 0x5c, 0xcf, 0xde, 0x87, 0xbb, 0x22, 0x87, 0x59, 0xd5, 0x3c, 0xde, 0xeb, 0xaa, 0x57, 0x6f,
	0x82, 0x1b, 0x25, 0x58, 0x26, 0x91, 0xdb, 0x64, 0x04, 0x63, 0xc1, 0x4b, 0x09, 0xcf, 0x2d, 0x1f,
	0x39, 0xa7, 0x07, 0x39, 0x8a, 0x34, 0x16, 0xb9, 0x50, 0x09, 0x18, 0xae, 0x8b, 0x38, 0x97, 0xc9,
	0x20, 0x29, 0xac, 0xc3, 0xb1, 0xd0, 0x92, 0x69, 0x83, 0x0e, 0xc3, 0x8b, 0x35, 0x9b, 0xd5, 0x6c,
	0xe6, 0xd9, 0x6c, 0x8e, 0xdd, 0x8e, 0x32, 0xe9, 0x46, 0x45, 0xcc, 0x12, 0x1c, 0xf3, 0x0c, 0x33,
	0xe4, 0x9e, 0x1d, 0x17, 0x43, 0xdf, 0xf9, 0xc6, 0x57, 0xb5, 0x6a, 0xbb, 0x93, 0x21, 0x66, 0x39,
	0x70, 0xa1, 0x25, 0x17, 0x4a, 0xa1, 0x13, 0x4e, 0xa2, 0xb2, 0x53, 0xf4, 0xfc, 0x9f, 0x13, 0xa3,
	0x3e, 0x0c, 0x76, 0x66, 0xd6, 0x11, 0xb9, 0x4c, 0x85, 0x83, 0x29, 0xda, 0x9d, 0x41, 0xc1, 0x82,
	0x2a, 0x67, 0x14, 0xae, 0xcc, 0x1a, 0x62, 0x5c, 0x21, 0xf2, 0xc1, 0x08, 0xad, 0x1b, 0xa4, 0xca,
	0x0e, 0xa4, 0x1a, 0x22, 0xc7, 0xf8, 0x09, 0x24, 0xae, 0x66, 0x74, 0x97, 0x83, 0xb3, 0x3d, 0x70,
	0x77, 0x94, 0x5d, 0x51, 0x43, 0xec, 0xc3, 0xb3, 0x02, 0xac, 0x0b, 0x3b, 0xc1, 0x82, 0x12, 0x63,
	0xb0, 0x5a, 0x24, 0xd0, 0x22, 0x17, 0xc8, 0xa5, 0x85, 0xfe, 0xc1, 0x87, 0x30, 0x0c, 0x8e, 0xef,
	0x36, 0xad, 0xa6, 0x07, 0x7c, 0xdd, 0x4d, 0x83, 0xf0, 0xb0, 0x8c, 0xd5, 0xa8, 0x2c, 0x84, 0x77,
	0x83, 0x53, 0x7b, 0xaf, 0x7a, 0x99, 0xd3, 0x8b, 0x4b, 0x6c, 0xd6, 0xf4, 0xbf, 0x4c, 0xc8, 0x7a,
	0x39, 0xc6, 0x22, 0x7f, 0xa0, 0x21, 0x79, 0xf8, 0x52, 0x43, 0xff, 0x64, 0x5a, 0xeb, 0x2e, 0x7e,
	0x68, 0x06, 0x0b, 0xb7, 0xfd, 0x15, 0x6f, 0xde, 0x5f, 0x09, 0x7f, 0x91, 0x20, 0x38, 0x78, 0x34,
	0xbc, 0xca, 0x8e, 0x74, 0x4f, 0x36, 0xb7, 0x6e, 0xfb, 0xda, 0x7f, 0x30, 0xeb, 0x0d, 0xbb, 0xa3,
	0xcd, 0x8f, 0x4d, 0x52, 0x7d, 0x6a, 0x9d, 0x2b, 0xc1, 0x46, 0x12, 0xa3, 0x0c, 0x14, 0x18, 0x91,
	0x47, 0x06, 0x44, 0xba, 0xf6, 0xf5, 0xc7, 0xbb, 0x66, 0x2f, 0x5c, 0x9e, 0xe6, 0x8f, 0xef, 0x3b,
	0x68, 0xf9, 0xab, 0xfd, 0x7a, 0x75, 0x3e, 0xaf, 0x53, 0x78, 0x95, 0x67, 0xe0, 0xa2, 0x54, 0xd9,
	0x68, 0xd7, 0x93, 0xf6, 0xf5, 0x8d, 0x75, 0x72, 0xec, 0xcb, 0x3a, 0xb9, 0x7c, 0xb4, 0x59, 0xef,
	0xf9, 0x23, 0xaf, 0x7d, 0x6e, 0x35, 0xcf, 0x90, 0x5b, 0x6f, 0xc8, 0xd6, 0x84, 0x36, 0xbe, 0x4d,
	0x68, 0x63, 0x67, 0x42, 0xc9, 0xeb, 0x8a, 0x92, 0xf7, 0x15, 0x25, 0x9b, 0x15, 0x25, 0x5b, 0x15,
	0x25, 0xdf, 0x2b, 0x4a, 0x7e, 0x56, 0xb4, 0xb1, 0x53, 0x51, 0xf2, 0x76, 0x9b, 0x36, 0x36, 0xb6,
	0x29, 0x79, 0xf4, 0x38, 0x43, 0xfd, 0x34, 0x63, 0x25, 0xe6, 0x0e, 0x8c, 0x11, 0xac, 0xb0, 0xdc,
	0x17, 0x43, 0x34, 0xe3, 0x48, 0x1b, 0x2c, 0x65, 0x0a, 0x26, 0xda, 0x83, 0xb9, 0x8e, 0x33, 0xe4,
	0xf0, 0xc2, 0x4d, 0x83, 0xf7, 0xef, 0x1f, 0x32, 0x3e, 0xe1, 0xb3, 0xb7, 0xf4, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0x80, 0x24, 0xb0, 0xc0, 0x03, 0x00, 0x00,
}
