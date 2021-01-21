// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/token/public_customapi.proto

package token

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

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

// Change token
//
// x-displayName: "Change Token"
// Generic response when object is changed
type ObjectChangeResp struct {
	// Token object
	//
	// x-displayName: "Token Object"
	Obj *Object `protobuf:"bytes,1,opt,name=obj" json:"obj,omitempty"`
}

func (m *ObjectChangeResp) Reset()                    { *m = ObjectChangeResp{} }
func (*ObjectChangeResp) ProtoMessage()               {}
func (*ObjectChangeResp) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{0} }

func (m *ObjectChangeResp) GetObj() *Object {
	if m != nil {
		return m.Obj
	}
	return nil
}

// Change token state
//
// x-displayName: "State Request"
// Changes token state, only state is changing, namespace and name is used to find token to change
type StateReq struct {
	// Token namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace where token is residing
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Token name
	//
	// x-displayName: "Name"
	// x-example: "value"
	// Token object name to change, it's usually same as Metadata.Uid
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Token state
	//
	// x-displayName: "Token State"
	// Change token state to this state
	State State `protobuf:"varint,3,opt,name=state,proto3,enum=ves.io.schema.token.State" json:"state,omitempty"`
}

func (m *StateReq) Reset()                    { *m = StateReq{} }
func (*StateReq) ProtoMessage()               {}
func (*StateReq) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{1} }

func (m *StateReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *StateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StateReq) GetState() State {
	if m != nil {
		return m.State
	}
	return UNKNOWN
}

func init() {
	proto.RegisterType((*ObjectChangeResp)(nil), "ves.io.schema.token.ObjectChangeResp")
	golang_proto.RegisterType((*ObjectChangeResp)(nil), "ves.io.schema.token.ObjectChangeResp")
	proto.RegisterType((*StateReq)(nil), "ves.io.schema.token.StateReq")
	golang_proto.RegisterType((*StateReq)(nil), "ves.io.schema.token.StateReq")
}
func (this *ObjectChangeResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ObjectChangeResp)
	if !ok {
		that2, ok := that.(ObjectChangeResp)
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
	if !this.Obj.Equal(that1.Obj) {
		return false
	}
	return true
}
func (this *StateReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StateReq)
	if !ok {
		that2, ok := that.(StateReq)
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
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *ObjectChangeResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&token.ObjectChangeResp{")
	if this.Obj != nil {
		s = append(s, "Obj: "+fmt.Sprintf("%#v", this.Obj)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StateReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&token.StateReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
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
	// Token states
	//
	// x-displayName: "Set Token State"
	// TokenState changes token status, it can be used to disable token.
	TokenState(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*ObjectChangeResp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) TokenState(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*ObjectChangeResp, error) {
	out := new(ObjectChangeResp)
	err := grpc.Invoke(ctx, "/ves.io.schema.token.CustomAPI/TokenState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomAPI service

type CustomAPIServer interface {
	// Token states
	//
	// x-displayName: "Set Token State"
	// TokenState changes token status, it can be used to disable token.
	TokenState(context.Context, *StateReq) (*ObjectChangeResp, error)
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_TokenState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).TokenState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.token.CustomAPI/TokenState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).TokenState(ctx, req.(*StateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.token.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TokenState",
			Handler:    _CustomAPI_TokenState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/token/public_customapi.proto",
}

func (m *ObjectChangeResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectChangeResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Obj != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.Obj.Size()))
		n1, err := m.Obj.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *StateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateReq) MarshalTo(dAtA []byte) (int, error) {
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
	if m.State != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.State))
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
func (m *ObjectChangeResp) Size() (n int) {
	var l int
	_ = l
	if m.Obj != nil {
		l = m.Obj.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *StateReq) Size() (n int) {
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
	if m.State != 0 {
		n += 1 + sovPublicCustomapi(uint64(m.State))
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
func (this *ObjectChangeResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ObjectChangeResp{`,
		`Obj:` + strings.Replace(fmt.Sprintf("%v", this.Obj), "Object", "Object", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StateReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StateReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
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
func (m *ObjectChangeResp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ObjectChangeResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectChangeResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Obj", wireType)
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
			if m.Obj == nil {
				m.Obj = &Object{}
			}
			if err := m.Obj.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *StateReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (State(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
	proto.RegisterFile("ves.io/schema/token/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/token/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0xc6, 0xb7, 0x77, 0x55, 0xdc, 0x16, 0x44, 0x46, 0x0f, 0xcb, 0x24, 0xb6, 0xcb, 0x80, 0xb0,
	0x04, 0x66, 0x5a, 0xc6, 0x83, 0xe0, 0x49, 0x13, 0x10, 0x72, 0x52, 0x56, 0xbd, 0x78, 0x91, 0x9e,
	0xb1, 0xd2, 0x3b, 0xc9, 0xce, 0x54, 0x67, 0xba, 0x67, 0x54, 0xc2, 0x82, 0xe4, 0x09, 0x04, 0x5f,
	0x42, 0x9f, 0x21, 0x97, 0xbd, 0xe9, 0x49, 0x82, 0x5e, 0x3c, 0xba, 0x13, 0x0f, 0x1e, 0xf3, 0x08,
	0xb2, 0x3d, 0x93, 0x3f, 0x0b, 0x6b, 0x6e, 0x55, 0x5d, 0xbf, 0xfa, 0xba, 0xba, 0xeb, 0xa3, 0x6b,
	0x25, 0xe8, 0x20, 0x41, 0xae, 0xe3, 0x11, 0xa4, 0x82, 0x1b, 0xdc, 0x81, 0x8c, 0xab, 0x22, 0x1a,
	0x27, 0xf1, 0xeb, 0xb8, 0xd0, 0x06, 0x53, 0xa1, 0x92, 0x40, 0xe5, 0x68, 0xd0, 0xb9, 0x59, 0xb3,
	0x41, 0xcd, 0x06, 0x96, 0x75, 0x7d, 0x99, 0x98, 0x51, 0x11, 0x05, 0x31, 0xa6, 0x5c, 0xa2, 0x44,
	0x6e, 0xd9, 0xa8, 0xd8, 0xb2, 0x99, 0x4d, 0x6c, 0x54, 0x6b, 0xb8, 0xab, 0x12, 0x51, 0x8e, 0x81,
	0x0b, 0x95, 0x70, 0x91, 0x65, 0x68, 0x84, 0x49, 0x30, 0xd3, 0x4d, 0x75, 0x65, 0x71, 0x1a, 0x54,
	0xe7, 0x8b, 0xfd, 0x65, 0xa3, 0x62, 0xb4, 0x0d, 0xb1, 0x69, 0x88, 0x3b, 0xcb, 0x08, 0xf3, 0x5e,
	0xc1, 0x89, 0x84, 0xb7, 0x08, 0x94, 0xa0, 0x21, 0x2b, 0x17, 0xaf, 0xf1, 0x9e, 0xd0, 0x1b, 0x4f,
	0xad, 0xe8, 0xc6, 0x48, 0x64, 0x12, 0x86, 0xa0, 0x95, 0x13, 0xd2, 0x0e, 0x46, 0xdb, 0x3d, 0xd2,
	0x27, 0x83, 0x6b, 0xe1, 0x4a, 0xb0, 0xe4, 0x1f, 0x82, 0xba, 0x67, 0xbd, 0x33, 0x9d, 0x90, 0xe1,
	0x1c, 0xf6, 0x32, 0x7a, 0xf5, 0xb9, 0x11, 0x06, 0x86, 0xb0, 0xeb, 0xac, 0xd2, 0x6e, 0x26, 0x52,
	0xd0, 0x4a, 0xc4, 0x60, 0x55, 0xba, 0xc3, 0xb3, 0x03, 0xc7, 0xa1, 0x97, 0xe6, 0x49, 0xaf, 0x6d,
	0x0b, 0x36, 0x76, 0xee, 0xd1, 0xcb, 0x7a, 0xde, 0xdd, 0xeb, 0xf4, 0xc9, 0xe0, 0x7a, 0xe8, 0x2e,
	0xbd, 0xb3, 0xd6, 0xaf, 0xc1, 0x70, 0x46, 0x68, 0x77, 0xc3, 0x6e, 0xec, 0xf1, 0xb3, 0x4d, 0xe7,
	0x0b, 0xa1, 0xf4, 0xc5, 0x1c, 0xb2, 0x8c, 0x73, 0xfb, 0x82, 0x7e, 0xd8, 0x75, 0xef, 0x5e, 0xf0,
	0xa4, 0xb3, 0x6f, 0xf0, 0x5e, 0x56, 0x5f, 0x7b, 0xb7, 0x4a, 0xd0, 0x7e, 0x82, 0xbe, 0x84, 0x0c,
	0x72, 0x31, 0xf6, 0xdf, 0xe6, 0x89, 0x81, 0xfd, 0x9f, 0x7f, 0x3e, 0xb5, 0x1f, 0x78, 0x61, 0xe3,
	0x1c, 0x7e, 0xfa, 0x36, 0xcd, 0xf7, 0x4e, 0xe3, 0x49, 0xbd, 0x90, 0xe6, 0x68, 0xc2, 0xed, 0xd8,
	0x0f, 0xc9, 0x9a, 0x3b, 0x98, 0x1e, 0x90, 0xce, 0x8f, 0x03, 0xe2, 0xfe, 0x7f, 0x88, 0xfd, 0xef,
	0xbd, 0xf6, 0x23, 0xb2, 0xbe, 0x77, 0x38, 0x63, 0xad, 0x5f, 0x33, 0xd6, 0x3a, 0x9e, 0x31, 0xf2,
	0xa1, 0x62, 0xe4, 0x73, 0xc5, 0xc8, 0xb7, 0x8a, 0x91, 0xc3, 0x8a, 0x91, 0xdf, 0x15, 0x23, 0x7f,
	0x2b, 0xd6, 0x3a, 0xae, 0x18, 0xf9, 0x78, 0xc4, 0x5a, 0xd3, 0x23, 0x46, 0x5e, 0x6d, 0x4a, 0x54,
	0x3b, 0x32, 0x28, 0x71, 0x6c, 0x20, 0xcf, 0x45, 0x50, 0x68, 0x6e, 0x83, 0x2d, 0xcc, 0x53, 0x5f,
	0xe5, 0x58, 0x26, 0x6f, 0x20, 0xf7, 0x4f, 0xca, 0x5c, 0x45, 0x12, 0x39, 0xbc, 0x33, 0x8d, 0x41,
	0xce, 0x1b, 0x29, 0xba, 0x62, 0xfd, 0x71, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x7d,
	0x3e, 0x92, 0x33, 0x03, 0x00, 0x00,
}
