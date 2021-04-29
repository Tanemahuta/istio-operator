// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1alpha1/istiocontrolplane.proto

// Copyright 2021 Cisco Systems, Inc. and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1alpha1 "istio.io/api/mesh/v1alpha1"
	math "math"
	math_bits "math/bits"
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

// IstioControlPlane defines an Istio control plane
//
// <!-- crd generation tags
// +cue-gen:IstioControlPlane:groupName:servicemesh.cisco.com
// +cue-gen:IstioControlPlane:version:v1alpha1
// +cue-gen:IstioControlPlane:storageVersion
// +cue-gen:IstioControlPlane:annotations:helm.sh/resource-policy=keep
// +cue-gen:IstioControlPlane:subresource:status
// +cue-gen:IstioControlPlane:scope:Namespaced
// +cue-gen:IstioControlPlane:resource:shortNames=icp,istiocp
// +cue-gen:IstioControlPlane:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:IstioControlPlane:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneSpec struct {
	// Defines mesh-wide settings for the Istio control plane.
	MeshConfig           *v1alpha1.MeshConfig `protobuf:"bytes,1,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IstioControlPlaneSpec) Reset()         { *m = IstioControlPlaneSpec{} }
func (m *IstioControlPlaneSpec) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneSpec) ProtoMessage()    {}
func (*IstioControlPlaneSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{0}
}
func (m *IstioControlPlaneSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneSpec.Merge(m, src)
}
func (m *IstioControlPlaneSpec) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneSpec proto.InternalMessageInfo

func (m *IstioControlPlaneSpec) GetMeshConfig() *v1alpha1.MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneStatus struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioControlPlaneStatus) Reset()         { *m = IstioControlPlaneStatus{} }
func (m *IstioControlPlaneStatus) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneStatus) ProtoMessage()    {}
func (*IstioControlPlaneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{1}
}
func (m *IstioControlPlaneStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneStatus.Merge(m, src)
}
func (m *IstioControlPlaneStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneStatus proto.InternalMessageInfo

func (m *IstioControlPlaneStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*IstioControlPlaneSpec)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneSpec")
	proto.RegisterType((*IstioControlPlaneStatus)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneStatus")
}

func init() {
	proto.RegisterFile("api/v1alpha1/istiocontrolplane.proto", fileDescriptor_6817de833805cb8b)
}

var fileDescriptor_6817de833805cb8b = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0xcf, 0x2c, 0x2e, 0xc9, 0xcc, 0x4f, 0xce,
	0xcf, 0x2b, 0x29, 0xca, 0xcf, 0x29, 0xc8, 0x49, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x03, 0x4b, 0xc4, 0xe7, 0x17, 0xa4, 0x16, 0x25, 0x96, 0xe4, 0x17, 0xe9, 0x95, 0x19,
	0xe9, 0x25, 0x16, 0x64, 0xea, 0xc1, 0xf4, 0x49, 0x49, 0xe5, 0xa6, 0x16, 0x67, 0x20, 0x8c, 0x49,
	0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x87, 0xe8, 0x55, 0x8a, 0xe4, 0x12, 0xf5, 0x04, 0xe9, 0x76, 0x86,
	0x18, 0x1b, 0x00, 0x32, 0x36, 0xb8, 0x20, 0x35, 0x59, 0xc8, 0x81, 0x8b, 0x1b, 0xa4, 0x2d, 0x1e,
	0xa2, 0x5a, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5e, 0x0f, 0x6c, 0x95, 0x1e, 0x48, 0x06,
	0x6e, 0xbe, 0x9e, 0x6f, 0x6a, 0x71, 0x86, 0x33, 0x58, 0x59, 0x10, 0x57, 0x2e, 0x9c, 0xad, 0xa4,
	0xcb, 0x25, 0x8e, 0x69, 0x74, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0x2a, 0xd8, 0x54, 0xce, 0x20, 0x30, 0xdb, 0xc9, 0xf9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x8c, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0x4a, 0xcc, 0xab, 0x4a, 0xcc, 0x4c, 0xce, 0xc9, 0x2f, 0x4d, 0x81, 0x84,
	0x81, 0x2e, 0xcc, 0xab, 0xfa, 0x65, 0x46, 0xfa, 0xc8, 0x41, 0x94, 0xc4, 0x06, 0xf6, 0x95, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x23, 0x8a, 0xb8, 0x39, 0x01, 0x00, 0x00,
}

func (m *IstioControlPlaneSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MeshConfig != nil {
		{
			size, err := m.MeshConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIstiocontrolplane(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IstioControlPlaneStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIstiocontrolplane(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIstiocontrolplane(dAtA []byte, offset int, v uint64) int {
	offset -= sovIstiocontrolplane(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioControlPlaneSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MeshConfig != nil {
		l = m.MeshConfig.Size()
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioControlPlaneStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIstiocontrolplane(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIstiocontrolplane(x uint64) (n int) {
	return sovIstiocontrolplane(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioControlPlaneSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeshConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MeshConfig == nil {
				m.MeshConfig = &v1alpha1.MeshConfig{}
			}
			if err := m.MeshConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IstioControlPlaneStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIstiocontrolplane(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIstiocontrolplane
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
					return 0, ErrIntOverflowIstiocontrolplane
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
					return 0, ErrIntOverflowIstiocontrolplane
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
				return 0, ErrInvalidLengthIstiocontrolplane
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIstiocontrolplane
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIstiocontrolplane
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIstiocontrolplane        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIstiocontrolplane          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIstiocontrolplane = fmt.Errorf("proto: unexpected end of group")
)
