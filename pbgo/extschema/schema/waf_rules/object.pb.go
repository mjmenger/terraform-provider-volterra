// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/waf_rules/object.proto

/*
	Package waf_rules is a generated protocol buffer package.

	It is generated from these files:
		ves.io/schema/waf_rules/object.proto
		ves.io/schema/waf_rules/public_crudapi.proto
		ves.io/schema/waf_rules/public_customapi.proto
		ves.io/schema/waf_rules/types.proto

	It has these top-level messages:
		Object
		SpecType
		StatusObject
		CreateRequest
		CreateResponse
		ReplaceRequest
		ReplaceResponse
		GetRequest
		GetResponse
		ListRequest
		ListResponseItem
		ListResponse
		DeleteRequest
		RulesReq
		RulesRsp
		WafRulesStatusReq
		WafRulesStatusRsp
		WafRulesStatus
		VirtualHostWafRulesStatusReq
		VirtualHostWafRulesStatusRsp
		Rules
		GlobalSpecType
		CreateSpecType
		ReplaceSpecType
		GetSpecType
*/
package waf_rules

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// WAF Rules Object
//
// x-displayName: "WAF Rules Object"
// WAF Rules Object to define low level detailed configuration for a WAF instance which can be associated with vhosts or routes.
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *ves_io_schema4.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *ves_io_schema4.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the WAF instance
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{0} }

func (m *Object) GetMetadata() *ves_io_schema4.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *ves_io_schema4.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

// WAF Rules Specification
//
// x-displayName: "Specification"
// Shape of the WAF specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	GcSpec *GlobalSpecType `protobuf:"bytes,1,opt,name=gc_spec,json=gcSpec" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()                    { *m = SpecType{} }
func (*SpecType) ProtoMessage()               {}
func (*SpecType) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{1} }

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

// WAF Rules Status
//
// x-displayName: "WAF Rules Status Object"
// Most recently observed status of object
type StatusObject struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard status's metadata
	Metadata *ves_io_schema4.StatusMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// object_refs
	//
	// x-displayName: "Config Object"
	// Object reference
	ObjectRefs []*ves_io_schema4.ObjectRefType `protobuf:"bytes,2,rep,name=object_refs,json=objectRefs" json:"object_refs,omitempty"`
	// conditions
	//
	// x-displayName: "Conditions"
	Conditions []*ves_io_schema4.ConditionType `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
	// mode
	//
	// x-displayName: "Mode"
	// WAF Mode is blocking or Alert
	Mode ves_io_schema4.WafModeType `protobuf:"varint,4,opt,name=mode,proto3,enum=ves.io.schema.WafModeType" json:"mode,omitempty"`
	// anomaly score threshold
	//
	// x-displayName: "Anomaly Score Threshold"
	// x-example: "5"
	// if anomaly score is same or above threshold , the req/resp will be blocked/alerted (depending
	// on non_blocking_mode configuration)
	AnomalyScoreThreshold uint32 `protobuf:"varint,5,opt,name=anomaly_score_threshold,json=anomalyScoreThreshold,proto3" json:"anomaly_score_threshold,omitempty"`
	// paranoia_level
	//
	// x-displayName: "Paranoia Level"
	// x-example: "3"
	// Paranoia level
	ParanoiaLevel uint32 `protobuf:"varint,6,opt,name=paranoia_level,json=paranoiaLevel,proto3" json:"paranoia_level,omitempty"`
	// rules
	//
	// x-displayName: "Rules"
	// WAF rules associated with this instance
	Rules []*Rules `protobuf:"bytes,7,rep,name=rules" json:"rules,omitempty"`
}

func (m *StatusObject) Reset()                    { *m = StatusObject{} }
func (*StatusObject) ProtoMessage()               {}
func (*StatusObject) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{2} }

func (m *StatusObject) GetMetadata() *ves_io_schema4.StatusMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatusObject) GetObjectRefs() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.ObjectRefs
	}
	return nil
}

func (m *StatusObject) GetConditions() []*ves_io_schema4.ConditionType {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *StatusObject) GetMode() ves_io_schema4.WafModeType {
	if m != nil {
		return m.Mode
	}
	return ves_io_schema4.BLOCK
}

func (m *StatusObject) GetAnomalyScoreThreshold() uint32 {
	if m != nil {
		return m.AnomalyScoreThreshold
	}
	return 0
}

func (m *StatusObject) GetParanoiaLevel() uint32 {
	if m != nil {
		return m.ParanoiaLevel
	}
	return 0
}

func (m *StatusObject) GetRules() []*Rules {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.waf_rules.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.waf_rules.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.waf_rules.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.waf_rules.SpecType")
	proto.RegisterType((*StatusObject)(nil), "ves.io.schema.waf_rules.StatusObject")
	golang_proto.RegisterType((*StatusObject)(nil), "ves.io.schema.waf_rules.StatusObject")
}
func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
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
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *StatusObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusObject)
	if !ok {
		that2, ok := that.(StatusObject)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.ObjectRefs) != len(that1.ObjectRefs) {
		return false
	}
	for i := range this.ObjectRefs {
		if !this.ObjectRefs[i].Equal(that1.ObjectRefs[i]) {
			return false
		}
	}
	if len(this.Conditions) != len(that1.Conditions) {
		return false
	}
	for i := range this.Conditions {
		if !this.Conditions[i].Equal(that1.Conditions[i]) {
			return false
		}
	}
	if this.Mode != that1.Mode {
		return false
	}
	if this.AnomalyScoreThreshold != that1.AnomalyScoreThreshold {
		return false
	}
	if this.ParanoiaLevel != that1.ParanoiaLevel {
		return false
	}
	if len(this.Rules) != len(that1.Rules) {
		return false
	}
	for i := range this.Rules {
		if !this.Rules[i].Equal(that1.Rules[i]) {
			return false
		}
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&waf_rules.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&waf_rules.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&waf_rules.StatusObject{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.ObjectRefs != nil {
		s = append(s, "ObjectRefs: "+fmt.Sprintf("%#v", this.ObjectRefs)+",\n")
	}
	if this.Conditions != nil {
		s = append(s, "Conditions: "+fmt.Sprintf("%#v", this.Conditions)+",\n")
	}
	s = append(s, "Mode: "+fmt.Sprintf("%#v", this.Mode)+",\n")
	s = append(s, "AnomalyScoreThreshold: "+fmt.Sprintf("%#v", this.AnomalyScoreThreshold)+",\n")
	s = append(s, "ParanoiaLevel: "+fmt.Sprintf("%#v", this.ParanoiaLevel)+",\n")
	if this.Rules != nil {
		s = append(s, "Rules: "+fmt.Sprintf("%#v", this.Rules)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SystemMetadata != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.SystemMetadata.Size()))
		n2, err := m.SystemMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Spec != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Spec.Size()))
		n3, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.GcSpec.Size()))
		n4, err := m.GcSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *StatusObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n5, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.ObjectRefs) > 0 {
		for _, msg := range m.ObjectRefs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Conditions) > 0 {
		for _, msg := range m.Conditions {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Mode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Mode))
	}
	if m.AnomalyScoreThreshold != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.AnomalyScoreThreshold))
	}
	if m.ParanoiaLevel != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.ParanoiaLevel))
	}
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedObject(r randyObject, easy bool) *Object {
	this := &Object{}
	if r.Intn(10) != 0 {
		this.Metadata = ves_io_schema4.NewPopulatedObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.SystemMetadata = ves_io_schema4.NewPopulatedSystemObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Spec = NewPopulatedSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSpecType(r randyObject, easy bool) *SpecType {
	this := &SpecType{}
	if r.Intn(10) != 0 {
		this.GcSpec = NewPopulatedGlobalSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStatusObject(r randyObject, easy bool) *StatusObject {
	this := &StatusObject{}
	if r.Intn(10) != 0 {
		this.Metadata = ves_io_schema4.NewPopulatedStatusMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.ObjectRefs = make([]*ves_io_schema4.ObjectRefType, v1)
		for i := 0; i < v1; i++ {
			this.ObjectRefs[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Conditions = make([]*ves_io_schema4.ConditionType, v2)
		for i := 0; i < v2; i++ {
			this.Conditions[i] = ves_io_schema4.NewPopulatedConditionType(r, easy)
		}
	}
	this.Mode = ves_io_schema4.WafModeType([]int32{0, 1}[r.Intn(2)])
	this.AnomalyScoreThreshold = uint32(r.Uint32())
	this.ParanoiaLevel = uint32(r.Uint32())
	if r.Intn(10) != 0 {
		v3 := r.Intn(5)
		this.Rules = make([]*Rules, v3)
		for i := 0; i < v3; i++ {
			this.Rules[i] = NewPopulatedRules(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyObject interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneObject(r randyObject) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringObject(r randyObject) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneObject(r)
	}
	return string(tmps)
}
func randUnrecognizedObject(r randyObject, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldObject(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldObject(dAtA []byte, r randyObject, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateObject(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateObject(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateObject(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Object) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *StatusObject) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if len(m.ObjectRefs) > 0 {
		for _, e := range m.ObjectRefs {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if m.Mode != 0 {
		n += 1 + sovObject(uint64(m.Mode))
	}
	if m.AnomalyScoreThreshold != 0 {
		n += 1 + sovObject(uint64(m.AnomalyScoreThreshold))
	}
	if m.ParanoiaLevel != 0 {
		n += 1 + sovObject(uint64(m.ParanoiaLevel))
	}
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	return n
}

func sovObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "ves_io_schema4.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "ves_io_schema4.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(fmt.Sprintf("%v", this.Spec), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusObject{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "StatusMetaType", "ves_io_schema4.StatusMetaType", 1) + `,`,
		`ObjectRefs:` + strings.Replace(fmt.Sprintf("%v", this.ObjectRefs), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`Conditions:` + strings.Replace(fmt.Sprintf("%v", this.Conditions), "ConditionType", "ves_io_schema4.ConditionType", 1) + `,`,
		`Mode:` + fmt.Sprintf("%v", this.Mode) + `,`,
		`AnomalyScoreThreshold:` + fmt.Sprintf("%v", this.AnomalyScoreThreshold) + `,`,
		`ParanoiaLevel:` + fmt.Sprintf("%v", this.ParanoiaLevel) + `,`,
		`Rules:` + strings.Replace(fmt.Sprintf("%v", this.Rules), "Rules", "Rules", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &ves_io_schema4.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func (m *StatusObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: StatusObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.StatusMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectRefs = append(m.ObjectRefs, &ves_io_schema4.ObjectRefType{})
			if err := m.ObjectRefs[len(m.ObjectRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &ves_io_schema4.ConditionType{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= (ves_io_schema4.WafModeType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnomalyScoreThreshold", wireType)
			}
			m.AnomalyScoreThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnomalyScoreThreshold |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParanoiaLevel", wireType)
			}
			m.ParanoiaLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParanoiaLevel |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &Rules{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowObject
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
				next, err := skipObject(dAtA[start:])
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
	ErrInvalidLengthObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/waf_rules/object.proto", fileDescriptorObject) }
func init() { golang_proto.RegisterFile("ves.io/schema/waf_rules/object.proto", fileDescriptorObject) }

var fileDescriptorObject = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd3, 0x4c,
	0x1c, 0xc7, 0x3b, 0xdb, 0x6e, 0x77, 0x77, 0xb6, 0xdb, 0xe7, 0x21, 0x22, 0x1b, 0xeb, 0x3a, 0xd4,
	0xae, 0x60, 0x11, 0x3a, 0x91, 0xfa, 0x07, 0x14, 0x11, 0xa9, 0x07, 0x0f, 0xb6, 0x08, 0xe9, 0x82,
	0xe0, 0x25, 0x4c, 0x92, 0x49, 0x1a, 0x4d, 0x3a, 0x21, 0x33, 0xa9, 0xf6, 0x20, 0xf8, 0x0e, 0xf4,
	0x65, 0x88, 0xaf, 0x40, 0x3d, 0xf5, 0xb8, 0xc7, 0xc5, 0xd3, 0x1e, 0x6d, 0xf6, 0xa2, 0xb7, 0x3d,
	0x7a, 0x94, 0x4c, 0x9a, 0xb2, 0x2d, 0x2d, 0x78, 0x9b, 0xf4, 0xfb, 0x99, 0x4f, 0xf3, 0xfb, 0x66,
	0x06, 0xde, 0x18, 0x51, 0x8e, 0x3d, 0xa6, 0x71, 0x6b, 0x40, 0x03, 0xa2, 0xbd, 0x25, 0x8e, 0x11,
	0xc5, 0x3e, 0xe5, 0x1a, 0x33, 0x5f, 0x53, 0x4b, 0xe0, 0x30, 0x62, 0x82, 0x29, 0xfb, 0x19, 0x85,
	0x33, 0x0a, 0xcf, 0xa9, 0x5a, 0xcb, 0xf5, 0xc4, 0x20, 0x36, 0xb1, 0xc5, 0x02, 0xcd, 0x65, 0x2e,
	0xd3, 0x24, 0x6f, 0xc6, 0x8e, 0x7c, 0x92, 0x0f, 0x72, 0x95, 0x79, 0x6a, 0x57, 0x17, 0xff, 0x8d,
	0x85, 0xc2, 0x63, 0x43, 0x3e, 0x0b, 0xaf, 0x2c, 0x86, 0x62, 0x1c, 0xd2, 0x3c, 0x3a, 0x58, 0x8c,
	0x46, 0xc4, 0xf7, 0x6c, 0x22, 0xe8, 0x2c, 0x3d, 0x5c, 0x37, 0xc3, 0x05, 0x45, 0xe3, 0x18, 0xc0,
	0xf2, 0x0b, 0x39, 0x93, 0xf2, 0x00, 0x6e, 0x07, 0x54, 0x10, 0x9b, 0x08, 0xa2, 0x82, 0x3a, 0x68,
	0xee, 0xb6, 0xaf, 0xe1, 0xc5, 0x01, 0x33, 0xb0, 0x47, 0x05, 0x39, 0x1a, 0x87, 0x54, 0x9f, 0xe3,
	0x4a, 0x17, 0xfe, 0xc7, 0xc7, 0x5c, 0xd0, 0xc0, 0x98, 0x1b, 0x36, 0xa4, 0xe1, 0x70, 0xc9, 0xd0,
	0x97, 0xd4, 0x92, 0xa7, 0x9a, 0xed, 0xed, 0xe5, 0xb6, 0x7b, 0xb0, 0xc4, 0x43, 0x6a, 0xa9, 0x45,
	0xa9, 0xb8, 0x8e, 0xd7, 0xb4, 0x8c, 0xfb, 0x21, 0xb5, 0xa4, 0x40, 0xe2, 0x8d, 0x2e, 0xdc, 0xce,
	0x7f, 0x51, 0x9e, 0xc0, 0x2d, 0xd7, 0x32, 0xa4, 0x25, 0x1b, 0xe5, 0xe6, 0x5a, 0xcb, 0x33, 0x9f,
	0x99, 0xc4, 0x9f, 0xbb, 0xca, 0xae, 0x95, 0xae, 0x1b, 0xdf, 0x8a, 0xb0, 0xd2, 0x17, 0x44, 0xc4,
	0xfc, 0x9f, 0xeb, 0xc9, 0xf0, 0x15, 0xf5, 0x3c, 0x87, 0xbb, 0xd9, 0xb9, 0x31, 0x22, 0xea, 0x70,
	0x75, 0xa3, 0x5e, 0x6c, 0xee, 0xb6, 0x0f, 0x56, 0x96, 0xab, 0x53, 0x27, 0xdd, 0xdc, 0xa9, 0x7c,
	0x79, 0xbf, 0x33, 0x7f, 0x47, 0x1d, 0xb2, 0x3c, 0xe4, 0xca, 0x23, 0x08, 0x2d, 0x36, 0xb4, 0x3d,
	0x79, 0x46, 0xd4, 0xe2, 0x4a, 0xd7, 0xd3, 0x1c, 0x90, 0x2f, 0x72, 0x81, 0x57, 0x30, 0x2c, 0x05,
	0xcc, 0xa6, 0x6a, 0xa9, 0x0e, 0x9a, 0xd5, 0x76, 0x6d, 0x69, 0xdf, 0x4b, 0xe2, 0xf4, 0x98, 0x4d,
	0xb3, 0x52, 0x53, 0x4e, 0xb9, 0x0f, 0xf7, 0xc9, 0x90, 0x05, 0xc4, 0x1f, 0x1b, 0xdc, 0x62, 0x11,
	0x35, 0xc4, 0x20, 0xa2, 0x7c, 0xc0, 0x7c, 0x5b, 0xdd, 0xac, 0x83, 0xe6, 0x9e, 0x7e, 0x79, 0x16,
	0xf7, 0xd3, 0xf4, 0x28, 0x0f, 0x95, 0xdb, 0xb0, 0x1a, 0x92, 0x88, 0x0c, 0x99, 0x47, 0x0c, 0x9f,
	0x8e, 0xa8, 0xaf, 0x96, 0x53, 0xbc, 0xb3, 0xf3, 0xfd, 0xf7, 0xa4, 0x58, 0xba, 0xb5, 0xa1, 0x96,
	0xf4, 0xbd, 0x1c, 0xe8, 0xa6, 0xb9, 0x72, 0x17, 0x6e, 0xca, 0x61, 0xd5, 0x2d, 0x39, 0x12, 0x5a,
	0xfb, 0xc1, 0x74, 0x59, 0x49, 0x06, 0x3f, 0xbc, 0xf4, 0xe3, 0xf1, 0xff, 0xb0, 0x0a, 0x2b, 0x79,
	0xd7, 0x38, 0xf6, 0xec, 0xce, 0x47, 0x70, 0x32, 0x45, 0x85, 0xd3, 0x29, 0x2a, 0x9c, 0x4f, 0x11,
	0xf8, 0x33, 0x45, 0xe0, 0x43, 0x82, 0xc0, 0xe7, 0x04, 0x81, 0xaf, 0x09, 0x02, 0x93, 0x04, 0x81,
	0xe3, 0x04, 0x81, 0x93, 0x04, 0x81, 0xd3, 0x04, 0x81, 0x9f, 0x09, 0x02, 0xbf, 0x12, 0x54, 0x38,
	0x4f, 0x10, 0xf8, 0x74, 0x86, 0x0a, 0x93, 0x33, 0x04, 0x5e, 0xf5, 0x5c, 0x16, 0xbe, 0x71, 0xf1,
	0x88, 0xf9, 0x82, 0x46, 0x11, 0xc1, 0x31, 0xd7, 0xe4, 0xc2, 0x61, 0x51, 0xd0, 0x0a, 0x23, 0x36,
	0xf2, 0x6c, 0x1a, 0xb5, 0xf2, 0x58, 0x0b, 0x4d, 0x97, 0x69, 0xf4, 0x9d, 0x98, 0xdd, 0xb5, 0xe5,
	0x2b, 0x67, 0x96, 0xe5, 0x6d, 0xbb, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x63, 0x35, 0x78,
	0x58, 0x04, 0x00, 0x00,
}
