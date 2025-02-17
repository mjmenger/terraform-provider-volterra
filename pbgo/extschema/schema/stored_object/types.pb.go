// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/stored_object/types.proto

package stored_object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

// ContentTransport
//
// The content transport represents how the actual content was transported from the user. Bytes or string.
type ContentTransport int32

const (
	CONTENT_TRANSPORT_NONE   ContentTransport = 0
	CONTENT_TRANSPORT_STRING ContentTransport = 1
	CONTENT_TRANSPORT_BYTES  ContentTransport = 2
)

var ContentTransport_name = map[int32]string{
	0: "CONTENT_TRANSPORT_NONE",
	1: "CONTENT_TRANSPORT_STRING",
	2: "CONTENT_TRANSPORT_BYTES",
}

var ContentTransport_value = map[string]int32{
	"CONTENT_TRANSPORT_NONE":   0,
	"CONTENT_TRANSPORT_STRING": 1,
	"CONTENT_TRANSPORT_BYTES":  2,
}

func (ContentTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_157271185aa104cd, []int{0}
}

// GlobalSpecType
//
// x-displayName: "Specification"
// Shape of stored_object in the storage backend
type GlobalSpecType struct {
	// object_type
	//
	// x-displayName: "Object Type"
	// x-example: swagger
	// x-required
	// The type of the stored object to be uploaded.
	// The supported object_types are: "swagger", "certificate", "html", "javascript" & "generic"
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// name
	//
	// x-displayName: "Name"
	// x-example: volt-api-specs
	// x-required
	// Name of the stored_object. This is the user provided "name" for this object.
	// Internally we name this object as type-name-version and is the index key for this object.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// version
	//
	// x-displayName: "Version"
	// x-example: "V0-01-01-2021"
	// x-required
	// Version of the stored_object in "v{n}-{YY}-{MM}-{DD}" formatted string , where n is version number and YY/MM/DD is year, month and date.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// hash
	//
	// x-displayName: "Hash"
	// x-example: "value"
	// x-required
	// Hash value of the stored_object
	ContentHash string `protobuf:"bytes,4,opt,name=content_hash,json=contentHash,proto3" json:"content_hash,omitempty"`
	// content_format
	//
	// x-displayName: "Content Format"
	// x-example: "json, yaml, js, html"
	// The optional content format associated with object
	ContentFormat string `protobuf:"bytes,5,opt,name=content_format,json=contentFormat,proto3" json:"content_format,omitempty"`
	// content_transport
	//
	// This field is set automatically based on Contents choice String or Bytes.
	// It is used when the Content need to be returned to the user in the same format as it was received.
	ContentTransport ContentTransport `protobuf:"varint,6,opt,name=content_transport,json=contentTransport,proto3,enum=ves.io.schema.stored_object.ContentTransport" json:"content_transport,omitempty"`
	// content_url_in_storage
	//
	// URL to the actual object in the backend storage.
	ContentUrlInStorage string `protobuf:"bytes,7,opt,name=content_url_in_storage,json=contentUrlInStorage,proto3" json:"content_url_in_storage,omitempty"`
	// object_attributes
	//
	// x-displayName: "Attributes of a object specific to its type"
	// Select the object attributes specific to its type
	//
	// Types that are valid to be assigned to ObjectAttributes:
	//	*GlobalSpecType_NoAttributes
	//	*GlobalSpecType_MobileSdk
	ObjectAttributes isGlobalSpecType_ObjectAttributes `protobuf_oneof:"object_attributes"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_157271185aa104cd, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

type isGlobalSpecType_ObjectAttributes interface {
	isGlobalSpecType_ObjectAttributes()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GlobalSpecType_NoAttributes struct {
	NoAttributes *schema.Empty `protobuf:"bytes,9,opt,name=no_attributes,json=noAttributes,proto3,oneof" json:"no_attributes,omitempty"`
}
type GlobalSpecType_MobileSdk struct {
	MobileSdk *MobileSDKAttributes `protobuf:"bytes,10,opt,name=mobile_sdk,json=mobileSdk,proto3,oneof" json:"mobile_sdk,omitempty"`
}

func (*GlobalSpecType_NoAttributes) isGlobalSpecType_ObjectAttributes() {}
func (*GlobalSpecType_MobileSdk) isGlobalSpecType_ObjectAttributes()    {}

func (m *GlobalSpecType) GetObjectAttributes() isGlobalSpecType_ObjectAttributes {
	if m != nil {
		return m.ObjectAttributes
	}
	return nil
}

func (m *GlobalSpecType) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *GlobalSpecType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GlobalSpecType) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GlobalSpecType) GetContentHash() string {
	if m != nil {
		return m.ContentHash
	}
	return ""
}

func (m *GlobalSpecType) GetContentFormat() string {
	if m != nil {
		return m.ContentFormat
	}
	return ""
}

func (m *GlobalSpecType) GetContentTransport() ContentTransport {
	if m != nil {
		return m.ContentTransport
	}
	return CONTENT_TRANSPORT_NONE
}

func (m *GlobalSpecType) GetContentUrlInStorage() string {
	if m != nil {
		return m.ContentUrlInStorage
	}
	return ""
}

func (m *GlobalSpecType) GetNoAttributes() *schema.Empty {
	if x, ok := m.GetObjectAttributes().(*GlobalSpecType_NoAttributes); ok {
		return x.NoAttributes
	}
	return nil
}

func (m *GlobalSpecType) GetMobileSdk() *MobileSDKAttributes {
	if x, ok := m.GetObjectAttributes().(*GlobalSpecType_MobileSdk); ok {
		return x.MobileSdk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GlobalSpecType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GlobalSpecType_NoAttributes)(nil),
		(*GlobalSpecType_MobileSdk)(nil),
	}
}

func init() {
	proto.RegisterEnum("ves.io.schema.stored_object.ContentTransport", ContentTransport_name, ContentTransport_value)
	golang_proto.RegisterEnum("ves.io.schema.stored_object.ContentTransport", ContentTransport_name, ContentTransport_value)
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.stored_object.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.stored_object.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/stored_object/types.proto", fileDescriptor_157271185aa104cd)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/stored_object/types.proto", fileDescriptor_157271185aa104cd)
}

var fileDescriptor_157271185aa104cd = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xbd, 0x7d, 0xf3, 0xb6, 0xc9, 0xf6, 0x8f, 0xd2, 0x05, 0x15, 0x93, 0x96, 0x6d, 0x84,
	0x84, 0xa8, 0x2a, 0xc5, 0x46, 0xe5, 0x08, 0x97, 0xa6, 0x84, 0xb6, 0x20, 0x52, 0x6a, 0x9b, 0x03,
	0x5c, 0x2c, 0xdb, 0xd9, 0x3a, 0xa6, 0xb6, 0xc7, 0xb2, 0x37, 0x81, 0x1e, 0x90, 0x38, 0x73, 0x81,
	0x8f, 0xc1, 0x77, 0xe8, 0xa5, 0x47, 0x8e, 0x39, 0xe6, 0x48, 0x9d, 0x4b, 0x8f, 0xfd, 0x08, 0xa8,
	0x6b, 0x9b, 0x34, 0x29, 0xea, 0x6d, 0x66, 0x9e, 0xdf, 0x8c, 0x9f, 0xf5, 0xee, 0xe0, 0xc7, 0x7d,
	0x96, 0x28, 0x1e, 0xa8, 0x89, 0xd3, 0x65, 0x81, 0xa5, 0x26, 0x1c, 0x62, 0xd6, 0x31, 0xc1, 0xfe,
	0xc8, 0x1c, 0xae, 0xf2, 0x93, 0x88, 0x25, 0x4a, 0x14, 0x03, 0x07, 0xb2, 0x9a, 0x81, 0x4a, 0x06,
	0x2a, 0x13, 0x60, 0xad, 0xe1, 0x7a, 0xbc, 0xdb, 0xb3, 0x15, 0x07, 0x02, 0xd5, 0x05, 0x17, 0x54,
	0xd1, 0x63, 0xf7, 0x8e, 0x44, 0x26, 0x12, 0x11, 0x65, 0xb3, 0x6a, 0xeb, 0x2e, 0x80, 0xeb, 0xb3,
	0x31, 0xc5, 0xbd, 0x80, 0x25, 0xdc, 0x0a, 0xa2, 0x1c, 0x58, 0x9d, 0x74, 0x05, 0x11, 0xf7, 0x20,
	0xcc, 0x9d, 0xd4, 0x94, 0xdb, 0x2c, 0x3b, 0x10, 0x04, 0x10, 0x9a, 0xd7, 0x9c, 0xd7, 0xee, 0x4f,
	0xf2, 0xd7, 0xa5, 0xb5, 0x49, 0xa9, 0x6f, 0xf9, 0x5e, 0xc7, 0xe2, 0x2c, 0x57, 0xeb, 0x53, 0xaa,
	0xc7, 0x3e, 0x99, 0x13, 0x56, 0x1e, 0x7e, 0x2b, 0xe1, 0xa5, 0x5d, 0x1f, 0x6c, 0xcb, 0xd7, 0x23,
	0xe6, 0x18, 0x27, 0x11, 0x23, 0xeb, 0x78, 0x3e, 0xb3, 0x22, 0x3c, 0xc8, 0xa8, 0x8e, 0x36, 0x2a,
	0x1a, 0xce, 0x4a, 0x02, 0x20, 0xb8, 0x14, 0x5a, 0x01, 0x93, 0x67, 0x84, 0x22, 0x62, 0x22, 0xe3,
	0xb9, 0x3e, 0x8b, 0x13, 0x0f, 0x42, 0xf9, 0x3f, 0x51, 0x2e, 0x52, 0xb2, 0x89, 0x17, 0x1c, 0x08,
	0x39, 0x0b, 0xb9, 0xd9, 0xb5, 0x92, 0xae, 0x5c, 0xba, 0x92, 0x9b, 0x73, 0xc3, 0x2f, 0xe8, 0xe2,
	0x14, 0x21, 0x6d, 0x3e, 0x17, 0xf7, 0xac, 0xa4, 0x4b, 0x1e, 0xe1, 0xa5, 0x82, 0x3d, 0x82, 0x38,
	0xb0, 0xb8, 0xfc, 0xbf, 0x18, 0xb6, 0x98, 0x57, 0x5f, 0x8a, 0x22, 0x71, 0xf0, 0x72, 0x81, 0xf1,
	0xd8, 0x0a, 0x93, 0x08, 0x62, 0x2e, 0xcf, 0xd6, 0xd1, 0xc6, 0xd2, 0x56, 0x43, 0xb9, 0xe5, 0x96,
	0x95, 0x9d, 0xac, 0xcb, 0x28, 0x9a, 0xc6, 0x36, 0xaa, 0xce, 0x94, 0x44, 0x9e, 0xe3, 0x95, 0xe2,
	0x23, 0xbd, 0xd8, 0x37, 0xbd, 0xd0, 0xbc, 0x9a, 0x65, 0xb9, 0x4c, 0x9e, 0x9b, 0x3c, 0xc1, 0x9d,
	0x1c, 0x7b, 0x17, 0xfb, 0xfb, 0xa1, 0x9e, 0x31, 0xe4, 0x19, 0x5e, 0x0c, 0xc1, 0xb4, 0x38, 0x8f,
	0x3d, 0xbb, 0xc7, 0x59, 0x22, 0x57, 0xea, 0x68, 0x63, 0x7e, 0xeb, 0xee, 0x94, 0xbd, 0x56, 0x10,
	0xf1, 0x93, 0x3d, 0x49, 0x5b, 0x08, 0x61, 0xfb, 0x2f, 0x4b, 0x0e, 0x31, 0x0e, 0xc0, 0xf6, 0x7c,
	0x66, 0x26, 0x9d, 0x63, 0x19, 0x8b, 0xce, 0x27, 0xb7, 0x1e, 0xec, 0x8d, 0xc0, 0xf5, 0x17, 0xaf,
	0xc7, 0x53, 0xf6, 0x24, 0xad, 0x92, 0x4d, 0xd1, 0x3b, 0xc7, 0xcd, 0x07, 0x78, 0x39, 0xbf, 0xd4,
	0xb1, 0x27, 0x52, 0x3e, 0x3b, 0x45, 0x95, 0xc1, 0x29, 0x2a, 0xbf, 0x2a, 0x95, 0xcb, 0xd5, 0xca,
	0xa6, 0x87, 0xab, 0xd3, 0x7f, 0x88, 0xd4, 0xf0, 0xca, 0xce, 0x41, 0xdb, 0x68, 0xb5, 0x0d, 0xd3,
	0xd0, 0xb6, 0xdb, 0xfa, 0xdb, 0x03, 0xcd, 0x30, 0xdb, 0x07, 0xed, 0x56, 0x55, 0x22, 0x6b, 0x58,
	0xbe, 0xa9, 0xe9, 0x86, 0xb6, 0xdf, 0xde, 0xad, 0x22, 0xb2, 0x8a, 0xef, 0xdd, 0x54, 0x9b, 0xef,
	0x8d, 0x96, 0x5e, 0x9d, 0x69, 0x7e, 0x47, 0x83, 0x73, 0x2a, 0x0d, 0xcf, 0xa9, 0x74, 0x79, 0x4e,
	0xd1, 0xd7, 0x94, 0xa2, 0x9f, 0x29, 0x45, 0xbf, 0x52, 0x8a, 0x06, 0x29, 0x45, 0xc3, 0x94, 0xa2,
	0xdf, 0x29, 0x45, 0x17, 0x29, 0x95, 0x2e, 0x53, 0x8a, 0x7e, 0x8c, 0xa8, 0x74, 0x36, 0xa2, 0x68,
	0x30, 0xa2, 0xd2, 0x70, 0x44, 0xa5, 0x0f, 0x87, 0x2e, 0x44, 0xc7, 0xae, 0xd2, 0x07, 0x9f, 0xb3,
	0x38, 0xb6, 0x94, 0x5e, 0xa2, 0x8a, 0xe0, 0xea, 0xd5, 0x34, 0xa2, 0x18, 0xfa, 0x5e, 0x87, 0xc5,
	0x8d, 0x42, 0x56, 0x23, 0xdb, 0x05, 0x95, 0x7d, 0xe6, 0xc5, 0xc2, 0xfd, 0x63, 0xef, 0xec, 0x59,
	0xb1, 0x10, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xf9, 0xa7, 0x4a, 0x50, 0x04, 0x00,
	0x00,
}

func (x ContentTransport) String() string {
	s, ok := ContentTransport_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	if this.ObjectType != that1.ObjectType {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.ContentHash != that1.ContentHash {
		return false
	}
	if this.ContentFormat != that1.ContentFormat {
		return false
	}
	if this.ContentTransport != that1.ContentTransport {
		return false
	}
	if this.ContentUrlInStorage != that1.ContentUrlInStorage {
		return false
	}
	if that1.ObjectAttributes == nil {
		if this.ObjectAttributes != nil {
			return false
		}
	} else if this.ObjectAttributes == nil {
		return false
	} else if !this.ObjectAttributes.Equal(that1.ObjectAttributes) {
		return false
	}
	return true
}
func (this *GlobalSpecType_NoAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_NoAttributes)
	if !ok {
		that2, ok := that.(GlobalSpecType_NoAttributes)
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
	if !this.NoAttributes.Equal(that1.NoAttributes) {
		return false
	}
	return true
}
func (this *GlobalSpecType_MobileSdk) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_MobileSdk)
	if !ok {
		that2, ok := that.(GlobalSpecType_MobileSdk)
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
	if !this.MobileSdk.Equal(that1.MobileSdk) {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&stored_object.GlobalSpecType{")
	s = append(s, "ObjectType: "+fmt.Sprintf("%#v", this.ObjectType)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "ContentHash: "+fmt.Sprintf("%#v", this.ContentHash)+",\n")
	s = append(s, "ContentFormat: "+fmt.Sprintf("%#v", this.ContentFormat)+",\n")
	s = append(s, "ContentTransport: "+fmt.Sprintf("%#v", this.ContentTransport)+",\n")
	s = append(s, "ContentUrlInStorage: "+fmt.Sprintf("%#v", this.ContentUrlInStorage)+",\n")
	if this.ObjectAttributes != nil {
		s = append(s, "ObjectAttributes: "+fmt.Sprintf("%#v", this.ObjectAttributes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType_NoAttributes) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&stored_object.GlobalSpecType_NoAttributes{` +
		`NoAttributes:` + fmt.Sprintf("%#v", this.NoAttributes) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_MobileSdk) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&stored_object.GlobalSpecType_MobileSdk{` +
		`MobileSdk:` + fmt.Sprintf("%#v", this.MobileSdk) + `}`}, ", ")
	return s
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ObjectAttributes != nil {
		{
			size := m.ObjectAttributes.Size()
			i -= size
			if _, err := m.ObjectAttributes.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.ContentUrlInStorage) > 0 {
		i -= len(m.ContentUrlInStorage)
		copy(dAtA[i:], m.ContentUrlInStorage)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ContentUrlInStorage)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ContentTransport != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ContentTransport))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ContentFormat) > 0 {
		i -= len(m.ContentFormat)
		copy(dAtA[i:], m.ContentFormat)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ContentFormat)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ContentHash) > 0 {
		i -= len(m.ContentHash)
		copy(dAtA[i:], m.ContentHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ContentHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ObjectType) > 0 {
		i -= len(m.ObjectType)
		copy(dAtA[i:], m.ObjectType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ObjectType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GlobalSpecType_NoAttributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_NoAttributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoAttributes != nil {
		{
			size, err := m.NoAttributes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_MobileSdk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_MobileSdk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MobileSdk != nil {
		{
			size, err := m.MobileSdk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ObjectType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ContentHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ContentFormat)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ContentTransport != 0 {
		n += 1 + sovTypes(uint64(m.ContentTransport))
	}
	l = len(m.ContentUrlInStorage)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ObjectAttributes != nil {
		n += m.ObjectAttributes.Size()
	}
	return n
}

func (m *GlobalSpecType_NoAttributes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoAttributes != nil {
		l = m.NoAttributes.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_MobileSdk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MobileSdk != nil {
		l = m.MobileSdk.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`ObjectType:` + fmt.Sprintf("%v", this.ObjectType) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`ContentHash:` + fmt.Sprintf("%v", this.ContentHash) + `,`,
		`ContentFormat:` + fmt.Sprintf("%v", this.ContentFormat) + `,`,
		`ContentTransport:` + fmt.Sprintf("%v", this.ContentTransport) + `,`,
		`ContentUrlInStorage:` + fmt.Sprintf("%v", this.ContentUrlInStorage) + `,`,
		`ObjectAttributes:` + fmt.Sprintf("%v", this.ObjectAttributes) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_NoAttributes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_NoAttributes{`,
		`NoAttributes:` + strings.Replace(fmt.Sprintf("%v", this.NoAttributes), "Empty", "schema.Empty", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_MobileSdk) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_MobileSdk{`,
		`MobileSdk:` + strings.Replace(fmt.Sprintf("%v", this.MobileSdk), "MobileSDKAttributes", "MobileSDKAttributes", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentTransport", wireType)
			}
			m.ContentTransport = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContentTransport |= ContentTransport(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentUrlInStorage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentUrlInStorage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &schema.Empty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ObjectAttributes = &GlobalSpecType_NoAttributes{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MobileSdk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MobileSDKAttributes{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ObjectAttributes = &GlobalSpecType_MobileSdk{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
