// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ResponseMatch struct {
	Matchers               []*matchers.HeaderMatcher      `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	ResponseCodeDetails    string                         `protobuf:"bytes,2,opt,name=response_code_details,json=responseCodeDetails,proto3" json:"response_code_details,omitempty"`
	ResponseTransformation *transformation.Transformation `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *ResponseMatch) Reset()         { *m = ResponseMatch{} }
func (m *ResponseMatch) String() string { return proto.CompactTextString(m) }
func (*ResponseMatch) ProtoMessage()    {}
func (*ResponseMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{0}
}
func (m *ResponseMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMatch.Unmarshal(m, b)
}
func (m *ResponseMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMatch.Marshal(b, m, deterministic)
}
func (m *ResponseMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMatch.Merge(m, src)
}
func (m *ResponseMatch) XXX_Size() int {
	return xxx_messageInfo_ResponseMatch.Size(m)
}
func (m *ResponseMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMatch proto.InternalMessageInfo

func (m *ResponseMatch) GetMatchers() []*matchers.HeaderMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

func (m *ResponseMatch) GetResponseCodeDetails() string {
	if m != nil {
		return m.ResponseCodeDetails
	}
	return ""
}

func (m *ResponseMatch) GetResponseTransformation() *transformation.Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type RequestMatch struct {
	Matcher                *matchers.Matcher              `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	ClearRouteCache        bool                           `protobuf:"varint,2,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	RequestTransformation  *transformation.Transformation `protobuf:"bytes,3,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	ResponseTransformation *transformation.Transformation `protobuf:"bytes,4,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *RequestMatch) Reset()         { *m = RequestMatch{} }
func (m *RequestMatch) String() string { return proto.CompactTextString(m) }
func (*RequestMatch) ProtoMessage()    {}
func (*RequestMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{1}
}
func (m *RequestMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMatch.Unmarshal(m, b)
}
func (m *RequestMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMatch.Marshal(b, m, deterministic)
}
func (m *RequestMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMatch.Merge(m, src)
}
func (m *RequestMatch) XXX_Size() int {
	return xxx_messageInfo_RequestMatch.Size(m)
}
func (m *RequestMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMatch proto.InternalMessageInfo

func (m *RequestMatch) GetMatcher() *matchers.Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (m *RequestMatch) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *RequestMatch) GetRequestTransformation() *transformation.Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RequestMatch) GetResponseTransformation() *transformation.Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type Transformations struct {
	// Apply a transformation to requests.
	RequestTransformation *transformation.Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// Clear the route cache if the request transformation was applied.
	ClearRouteCache bool `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	// Apply a transformation to responses.
	ResponseTransformation *transformation.Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *Transformations) Reset()         { *m = Transformations{} }
func (m *Transformations) String() string { return proto.CompactTextString(m) }
func (*Transformations) ProtoMessage()    {}
func (*Transformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{2}
}
func (m *Transformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformations.Unmarshal(m, b)
}
func (m *Transformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformations.Marshal(b, m, deterministic)
}
func (m *Transformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformations.Merge(m, src)
}
func (m *Transformations) XXX_Size() int {
	return xxx_messageInfo_Transformations.Size(m)
}
func (m *Transformations) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformations.DiscardUnknown(m)
}

var xxx_messageInfo_Transformations proto.InternalMessageInfo

func (m *Transformations) GetRequestTransformation() *transformation.Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *Transformations) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *Transformations) GetResponseTransformation() *transformation.Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type RequestResponseTransformations struct {
	RequestTransforms    []*RequestMatch  `protobuf:"bytes,1,rep,name=request_transforms,json=requestTransforms,proto3" json:"request_transforms,omitempty"`
	ResponseTransforms   []*ResponseMatch `protobuf:"bytes,2,rep,name=response_transforms,json=responseTransforms,proto3" json:"response_transforms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RequestResponseTransformations) Reset()         { *m = RequestResponseTransformations{} }
func (m *RequestResponseTransformations) String() string { return proto.CompactTextString(m) }
func (*RequestResponseTransformations) ProtoMessage()    {}
func (*RequestResponseTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{3}
}
func (m *RequestResponseTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponseTransformations.Unmarshal(m, b)
}
func (m *RequestResponseTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponseTransformations.Marshal(b, m, deterministic)
}
func (m *RequestResponseTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponseTransformations.Merge(m, src)
}
func (m *RequestResponseTransformations) XXX_Size() int {
	return xxx_messageInfo_RequestResponseTransformations.Size(m)
}
func (m *RequestResponseTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponseTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponseTransformations proto.InternalMessageInfo

func (m *RequestResponseTransformations) GetRequestTransforms() []*RequestMatch {
	if m != nil {
		return m.RequestTransforms
	}
	return nil
}

func (m *RequestResponseTransformations) GetResponseTransforms() []*ResponseMatch {
	if m != nil {
		return m.ResponseTransforms
	}
	return nil
}

type TransformationStages struct {
	Early                *RequestResponseTransformations `protobuf:"bytes,1,opt,name=early,proto3" json:"early,omitempty"`
	Regular              *RequestResponseTransformations `protobuf:"bytes,2,opt,name=regular,proto3" json:"regular,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TransformationStages) Reset()         { *m = TransformationStages{} }
func (m *TransformationStages) String() string { return proto.CompactTextString(m) }
func (*TransformationStages) ProtoMessage()    {}
func (*TransformationStages) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{4}
}
func (m *TransformationStages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationStages.Unmarshal(m, b)
}
func (m *TransformationStages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationStages.Marshal(b, m, deterministic)
}
func (m *TransformationStages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationStages.Merge(m, src)
}
func (m *TransformationStages) XXX_Size() int {
	return xxx_messageInfo_TransformationStages.Size(m)
}
func (m *TransformationStages) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationStages.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationStages proto.InternalMessageInfo

func (m *TransformationStages) GetEarly() *RequestResponseTransformations {
	if m != nil {
		return m.Early
	}
	return nil
}

func (m *TransformationStages) GetRegular() *RequestResponseTransformations {
	if m != nil {
		return m.Regular
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseMatch)(nil), "transformation.options.gloo.solo.io.ResponseMatch")
	proto.RegisterType((*RequestMatch)(nil), "transformation.options.gloo.solo.io.RequestMatch")
	proto.RegisterType((*Transformations)(nil), "transformation.options.gloo.solo.io.Transformations")
	proto.RegisterType((*RequestResponseTransformations)(nil), "transformation.options.gloo.solo.io.RequestResponseTransformations")
	proto.RegisterType((*TransformationStages)(nil), "transformation.options.gloo.solo.io.TransformationStages")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto", fileDescriptor_d8def6c8e0694580)
}

var fileDescriptor_d8def6c8e0694580 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbb, 0x8e, 0x13, 0x31,
	0x14, 0xd5, 0x24, 0x0b, 0xbb, 0x78, 0x41, 0xab, 0xf5, 0x3e, 0x88, 0x52, 0x44, 0x51, 0xb6, 0x09,
	0x48, 0xd8, 0xda, 0x20, 0xd1, 0xd1, 0x10, 0x10, 0x34, 0x5b, 0x60, 0xb6, 0x00, 0x24, 0x14, 0x9c,
	0xc9, 0xcd, 0x64, 0x60, 0x32, 0xd7, 0xd8, 0x4e, 0xc8, 0x96, 0xfc, 0x0d, 0x9f, 0x40, 0xc3, 0x3f,
	0xf0, 0x0d, 0x34, 0x54, 0x94, 0xf4, 0xc8, 0xf3, 0x88, 0x76, 0xf2, 0x18, 0x05, 0x91, 0xce, 0xf7,
	0xd8, 0xf7, 0xdc, 0x7b, 0x8e, 0x5f, 0xe4, 0x75, 0x10, 0xda, 0xd1, 0xa4, 0xcf, 0x7c, 0x1c, 0x73,
	0x83, 0x11, 0x3e, 0x08, 0x91, 0x07, 0x11, 0x22, 0x57, 0x1a, 0x3f, 0x80, 0x6f, 0x4d, 0x1a, 0x49,
	0x15, 0xf2, 0xe9, 0x39, 0x47, 0x65, 0x43, 0x8c, 0x0d, 0xb7, 0x5a, 0xc6, 0x66, 0x88, 0x7a, 0x2c,
	0x5d, 0xbc, 0x10, 0x32, 0xa5, 0xd1, 0x22, 0x3d, 0x5b, 0x40, 0xb3, 0x5c, 0xe6, 0xf8, 0x98, 0x2b,
	0xc5, 0x42, 0xac, 0x37, 0x02, 0xc4, 0x20, 0x02, 0x9e, 0xa4, 0xf4, 0x27, 0x43, 0xfe, 0x59, 0x4b,
	0xa5, 0x40, 0x9b, 0x94, 0xa4, 0xfe, 0x68, 0x7d, 0x2f, 0x3e, 0x6a, 0xe0, 0x63, 0x69, 0xfd, 0x11,
	0x68, 0x33, 0x1f, 0x64, 0x79, 0x97, 0x6b, 0xf2, 0x60, 0x66, 0x41, 0xc7, 0x32, 0xe2, 0x10, 0x4f,
	0xf1, 0x2a, 0x09, 0x63, 0xb3, 0xa9, 0xa4, 0xfa, 0x71, 0x80, 0x01, 0x26, 0x43, 0xee, 0x46, 0x19,
	0x4a, 0x61, 0x66, 0x53, 0x10, 0x66, 0x36, 0xc5, 0x5a, 0xbf, 0x3d, 0x72, 0x47, 0x80, 0x51, 0x18,
	0x1b, 0xb8, 0x70, 0xad, 0xd1, 0x67, 0x64, 0x2f, 0xef, 0xb1, 0xe6, 0x35, 0xab, 0xed, 0xfd, 0xce,
	0x3d, 0x36, 0x6f, 0xda, 0x69, 0x29, 0x18, 0xc3, 0x5e, 0x80, 0x1c, 0x80, 0xbe, 0x48, 0x17, 0x88,
	0x79, 0x2a, 0xed, 0x90, 0x13, 0x9d, 0xf1, 0xf6, 0x7c, 0x1c, 0x40, 0x6f, 0x00, 0x56, 0x86, 0x91,
	0xa9, 0x55, 0x9a, 0x5e, 0xfb, 0x96, 0x38, 0xca, 0x27, 0xbb, 0x38, 0x80, 0xa7, 0xe9, 0x14, 0x95,
	0xe4, 0xee, 0x3c, 0xa7, 0xa8, 0xab, 0x56, 0x6d, 0x7a, 0xed, 0xfd, 0x4e, 0x9b, 0x25, 0x76, 0x30,
	0xa9, 0x42, 0x36, 0xed, 0xb0, 0x61, 0x18, 0x59, 0xd0, 0x6c, 0x64, 0xad, 0x62, 0x97, 0x85, 0xf5,
	0xe2, 0x34, 0x27, 0x2a, 0xe2, 0xad, 0xef, 0x15, 0x72, 0x5b, 0xc0, 0xa7, 0x09, 0x18, 0x9b, 0xca,
	0x7d, 0x4c, 0x76, 0xb3, 0x9e, 0x6b, 0x5e, 0x52, 0xe3, 0xac, 0x4c, 0x6d, 0xae, 0x33, 0xcf, 0xa1,
	0xf7, 0xc9, 0xa1, 0x1f, 0x81, 0xd4, 0x3d, 0x8d, 0x13, 0x0b, 0x3d, 0x5f, 0xfa, 0x23, 0x48, 0x24,
	0xee, 0x89, 0x83, 0x64, 0x42, 0x38, 0xbc, 0xeb, 0x60, 0xda, 0x23, 0xa7, 0x3a, 0x2d, 0xfd, 0xbf,
	0xea, 0x4e, 0x32, 0x9e, 0x22, 0x5c, 0xe6, 0xdf, 0xce, 0x96, 0xfc, 0xfb, 0x52, 0x21, 0x07, 0x45,
	0xc8, 0x94, 0xe8, 0xf2, 0xb6, 0xa3, 0x6b, 0xa5, 0xc9, 0xd5, 0xd5, 0x26, 0x97, 0x78, 0x50, 0xd9,
	0x92, 0x07, 0xbf, 0x3c, 0xd2, 0xc8, 0xce, 0x90, 0x58, 0xb9, 0xc2, 0xd0, 0xf7, 0x84, 0x2e, 0x59,
	0x92, 0x5f, 0xa7, 0x73, 0xb6, 0xc1, 0x83, 0xc3, 0xae, 0x1f, 0x52, 0x71, 0xb8, 0xe8, 0x8b, 0xa1,
	0x3e, 0x39, 0x5a, 0xd6, 0xe9, 0x6e, 0x97, 0x2b, 0xd1, 0xd9, 0xb0, 0xc4, 0xb5, 0x7b, 0x2f, 0xe8,
	0x92, 0x5a, 0xd3, 0xfa, 0xe1, 0x91, 0xe3, 0xa2, 0xb4, 0x57, 0x56, 0x06, 0x60, 0xe8, 0x1b, 0x72,
	0x03, 0xa4, 0x8e, 0xae, 0xb2, 0x1d, 0xee, 0xfe, 0x8b, 0xa4, 0x35, 0x9e, 0x89, 0x94, 0x91, 0xbe,
	0x23, 0xbb, 0x1a, 0x82, 0x49, 0x24, 0x75, 0xb6, 0x61, 0x5b, 0x21, 0xcf, 0x39, 0x9f, 0xbc, 0xfc,
	0xf6, 0x67, 0xc7, 0xfb, 0xfa, 0xb3, 0xe1, 0xbd, 0x7d, 0xbe, 0xd9, 0x8f, 0xa2, 0x3e, 0x06, 0xe5,
	0xbf, 0x4a, 0xff, 0x66, 0xf2, 0x94, 0x3e, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xd4, 0x2f,
	0x89, 0xa3, 0x06, 0x00, 0x00,
}

func (this *ResponseMatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseMatch)
	if !ok {
		that2, ok := that.(ResponseMatch)
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
	if len(this.Matchers) != len(that1.Matchers) {
		return false
	}
	for i := range this.Matchers {
		if !this.Matchers[i].Equal(that1.Matchers[i]) {
			return false
		}
	}
	if this.ResponseCodeDetails != that1.ResponseCodeDetails {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestMatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestMatch)
	if !ok {
		that2, ok := that.(RequestMatch)
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
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	if this.ClearRouteCache != that1.ClearRouteCache {
		return false
	}
	if !this.RequestTransformation.Equal(that1.RequestTransformation) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Transformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformations)
	if !ok {
		that2, ok := that.(Transformations)
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
	if !this.RequestTransformation.Equal(that1.RequestTransformation) {
		return false
	}
	if this.ClearRouteCache != that1.ClearRouteCache {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestResponseTransformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestResponseTransformations)
	if !ok {
		that2, ok := that.(RequestResponseTransformations)
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
	if len(this.RequestTransforms) != len(that1.RequestTransforms) {
		return false
	}
	for i := range this.RequestTransforms {
		if !this.RequestTransforms[i].Equal(that1.RequestTransforms[i]) {
			return false
		}
	}
	if len(this.ResponseTransforms) != len(that1.ResponseTransforms) {
		return false
	}
	for i := range this.ResponseTransforms {
		if !this.ResponseTransforms[i].Equal(that1.ResponseTransforms[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationStages) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationStages)
	if !ok {
		that2, ok := that.(TransformationStages)
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
	if !this.Early.Equal(that1.Early) {
		return false
	}
	if !this.Regular.Equal(that1.Regular) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}