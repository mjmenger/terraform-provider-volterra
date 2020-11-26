// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/virtual_network/types.proto

package virtual_network

import (
	rand "math/rand"
	testing "testing"

	time "time"

	proto "github.com/gogo/protobuf/proto"

	jsonpb "github.com/gogo/protobuf/jsonpb"

	fmt "fmt"

	parser "go/parser"

	golang_proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestNextHopInterfaceListProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &NextHopInterfaceList{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestNextHopInterfaceListMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &NextHopInterfaceList{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkNextHopInterfaceListProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*NextHopInterfaceList, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedNextHopInterfaceList(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkNextHopInterfaceListProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedNextHopInterfaceList(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &NextHopInterfaceList{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestStaticRouteViewTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StaticRouteViewType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestStaticRouteViewTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StaticRouteViewType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkStaticRouteViewTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*StaticRouteViewType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedStaticRouteViewType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkStaticRouteViewTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedStaticRouteViewType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &StaticRouteViewType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestActivePBRPoliciesTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ActivePBRPoliciesType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestActivePBRPoliciesTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ActivePBRPoliciesType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkActivePBRPoliciesTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*ActivePBRPoliciesType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedActivePBRPoliciesType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkActivePBRPoliciesTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedActivePBRPoliciesType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &ActivePBRPoliciesType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestGlobalSpecTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GlobalSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestGlobalSpecTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GlobalSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkGlobalSpecTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*GlobalSpecType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedGlobalSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGlobalSpecTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedGlobalSpecType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &GlobalSpecType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestCreateSpecTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CreateSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestCreateSpecTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CreateSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkCreateSpecTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CreateSpecType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedCreateSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkCreateSpecTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedCreateSpecType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &CreateSpecType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestReplaceSpecTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ReplaceSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestReplaceSpecTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ReplaceSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkReplaceSpecTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*ReplaceSpecType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedReplaceSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkReplaceSpecTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedReplaceSpecType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &ReplaceSpecType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestGetSpecTypeProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GetSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestGetSpecTypeMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GetSpecType{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkGetSpecTypeProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*GetSpecType, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedGetSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkGetSpecTypeProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedGetSpecType(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &GetSpecType{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestNextHopInterfaceListJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &NextHopInterfaceList{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestStaticRouteViewTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &StaticRouteViewType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestActivePBRPoliciesTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ActivePBRPoliciesType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestGlobalSpecTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GlobalSpecType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestCreateSpecTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CreateSpecType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestReplaceSpecTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &ReplaceSpecType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestGetSpecTypeJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &GetSpecType{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestNextHopInterfaceListProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &NextHopInterfaceList{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestNextHopInterfaceListProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &NextHopInterfaceList{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestStaticRouteViewTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &StaticRouteViewType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestStaticRouteViewTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &StaticRouteViewType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestActivePBRPoliciesTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &ActivePBRPoliciesType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestActivePBRPoliciesTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &ActivePBRPoliciesType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestGlobalSpecTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &GlobalSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestGlobalSpecTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &GlobalSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCreateSpecTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &CreateSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCreateSpecTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &CreateSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestReplaceSpecTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &ReplaceSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestReplaceSpecTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &ReplaceSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestGetSpecTypeProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &GetSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestGetSpecTypeProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &GetSpecType{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestNextHopInterfaceListGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNextHopInterfaceList(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestStaticRouteViewTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedStaticRouteViewType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestActivePBRPoliciesTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedActivePBRPoliciesType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestGlobalSpecTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedGlobalSpecType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestCreateSpecTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCreateSpecType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestReplaceSpecTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplaceSpecType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetSpecTypeGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedGetSpecType(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestNextHopInterfaceListSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedNextHopInterfaceList(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkNextHopInterfaceListSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*NextHopInterfaceList, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedNextHopInterfaceList(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestStaticRouteViewTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedStaticRouteViewType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkStaticRouteViewTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*StaticRouteViewType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedStaticRouteViewType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestActivePBRPoliciesTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedActivePBRPoliciesType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkActivePBRPoliciesTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*ActivePBRPoliciesType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedActivePBRPoliciesType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestGlobalSpecTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGlobalSpecType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkGlobalSpecTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*GlobalSpecType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedGlobalSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestCreateSpecTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCreateSpecType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkCreateSpecTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CreateSpecType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedCreateSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestReplaceSpecTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedReplaceSpecType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkReplaceSpecTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*ReplaceSpecType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedReplaceSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestGetSpecTypeSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedGetSpecType(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkGetSpecTypeSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*GetSpecType, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedGetSpecType(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestNextHopInterfaceListStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNextHopInterfaceList(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestStaticRouteViewTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedStaticRouteViewType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestActivePBRPoliciesTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedActivePBRPoliciesType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestGlobalSpecTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedGlobalSpecType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestCreateSpecTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCreateSpecType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestReplaceSpecTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplaceSpecType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestGetSpecTypeStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedGetSpecType(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
