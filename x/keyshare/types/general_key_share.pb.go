// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/general_key_share.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
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

type GeneralKeyShare struct {
	Validator           string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	IdType              string `protobuf:"bytes,2,opt,name=idType,proto3" json:"idType,omitempty"`
	IdValue             string `protobuf:"bytes,3,opt,name=idValue,proto3" json:"idValue,omitempty"`
	KeyShare            string `protobuf:"bytes,4,opt,name=keyShare,proto3" json:"keyShare,omitempty"`
	KeyShareIndex       uint64 `protobuf:"varint,5,opt,name=keyShareIndex,proto3" json:"keyShareIndex,omitempty"`
	ReceivedTimestamp   uint64 `protobuf:"varint,6,opt,name=receivedTimestamp,proto3" json:"receivedTimestamp,omitempty"`
	ReceivedBlockHeight uint64 `protobuf:"varint,7,opt,name=receivedBlockHeight,proto3" json:"receivedBlockHeight,omitempty"`
}

func (m *GeneralKeyShare) Reset()         { *m = GeneralKeyShare{} }
func (m *GeneralKeyShare) String() string { return proto.CompactTextString(m) }
func (*GeneralKeyShare) ProtoMessage()    {}
func (*GeneralKeyShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_05ce460a69fa2745, []int{0}
}
func (m *GeneralKeyShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GeneralKeyShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GeneralKeyShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GeneralKeyShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralKeyShare.Merge(m, src)
}
func (m *GeneralKeyShare) XXX_Size() int {
	return m.Size()
}
func (m *GeneralKeyShare) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralKeyShare.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralKeyShare proto.InternalMessageInfo

func (m *GeneralKeyShare) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *GeneralKeyShare) GetIdType() string {
	if m != nil {
		return m.IdType
	}
	return ""
}

func (m *GeneralKeyShare) GetIdValue() string {
	if m != nil {
		return m.IdValue
	}
	return ""
}

func (m *GeneralKeyShare) GetKeyShare() string {
	if m != nil {
		return m.KeyShare
	}
	return ""
}

func (m *GeneralKeyShare) GetKeyShareIndex() uint64 {
	if m != nil {
		return m.KeyShareIndex
	}
	return 0
}

func (m *GeneralKeyShare) GetReceivedTimestamp() uint64 {
	if m != nil {
		return m.ReceivedTimestamp
	}
	return 0
}

func (m *GeneralKeyShare) GetReceivedBlockHeight() uint64 {
	if m != nil {
		return m.ReceivedBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*GeneralKeyShare)(nil), "fairyring.keyshare.GeneralKeyShare")
}

func init() {
	proto.RegisterFile("fairyring/keyshare/general_key_share.proto", fileDescriptor_05ce460a69fa2745)
}

var fileDescriptor_05ce460a69fa2745 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x39, 0x3b, 0x17, 0x10, 0x31, 0x82, 0x04, 0x91, 0x30, 0xc4, 0xc3, 0x10, 0x69,
	0x05, 0xbf, 0xc1, 0x0e, 0xea, 0xf0, 0x36, 0x87, 0x07, 0x2f, 0x23, 0x6d, 0x5f, 0xdb, 0xd0, 0x3f,
	0x29, 0x69, 0x36, 0x96, 0xb3, 0x5f, 0xc0, 0x8f, 0xe5, 0x71, 0x47, 0x8f, 0xd2, 0x7e, 0x11, 0x59,
	0x5c, 0x57, 0x44, 0x6f, 0x79, 0x9e, 0xdf, 0x2f, 0xef, 0xe1, 0xc1, 0x57, 0xaf, 0x5c, 0x28, 0xa3,
	0x44, 0x1e, 0x79, 0x09, 0x98, 0x32, 0xe6, 0x0a, 0xbc, 0x08, 0x72, 0x50, 0x3c, 0x9d, 0x27, 0x60,
	0xe6, 0xb6, 0x71, 0x0b, 0x25, 0xb5, 0x24, 0x64, 0xe7, 0xba, 0x8d, 0x7b, 0xf1, 0xd6, 0xc5, 0x47,
	0xf7, 0x3f, 0xfe, 0x23, 0x98, 0xa7, 0x4d, 0x47, 0xce, 0xf1, 0x60, 0xc9, 0x53, 0x11, 0x72, 0x2d,
	0x15, 0x45, 0x43, 0x34, 0x1a, 0x4c, 0xdb, 0x82, 0x9c, 0x62, 0x47, 0x84, 0x33, 0x53, 0x00, 0xed,
	0x5a, 0xb4, 0x4d, 0x84, 0xe2, 0xbe, 0x08, 0x9f, 0x79, 0xba, 0x00, 0xba, 0x67, 0x41, 0x13, 0xc9,
	0x19, 0x3e, 0x48, 0xb6, 0xb7, 0x69, 0xcf, 0xa2, 0x5d, 0x26, 0x97, 0xf8, 0xb0, 0x79, 0x4f, 0xf2,
	0x10, 0x56, 0x74, 0x7f, 0x88, 0x46, 0xbd, 0xe9, 0xef, 0x92, 0x5c, 0xe3, 0x63, 0x05, 0x01, 0x88,
	0x25, 0x84, 0x33, 0x91, 0x41, 0xa9, 0x79, 0x56, 0x50, 0xc7, 0x9a, 0x7f, 0x01, 0xb9, 0xc1, 0x27,
	0x4d, 0x39, 0x4e, 0x65, 0x90, 0x3c, 0x80, 0x88, 0x62, 0x4d, 0xfb, 0xd6, 0xff, 0x0f, 0x8d, 0x27,
	0x1f, 0x15, 0x43, 0xeb, 0x8a, 0xa1, 0xaf, 0x8a, 0xa1, 0xf7, 0x9a, 0x75, 0xd6, 0x35, 0xeb, 0x7c,
	0xd6, 0xac, 0xf3, 0xe2, 0x45, 0x42, 0xc7, 0x0b, 0xdf, 0x0d, 0x64, 0xe6, 0xdd, 0x71, 0xa1, 0xfc,
	0xcd, 0x2f, 0xaf, 0x1d, 0x7d, 0xd5, 0xce, 0xae, 0x4d, 0x01, 0xa5, 0xef, 0xd8, 0xad, 0x6f, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x07, 0x95, 0xb1, 0x99, 0x01, 0x00, 0x00,
}

func (m *GeneralKeyShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GeneralKeyShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GeneralKeyShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceivedBlockHeight != 0 {
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(m.ReceivedBlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.ReceivedTimestamp != 0 {
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(m.ReceivedTimestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.KeyShareIndex != 0 {
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(m.KeyShareIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.KeyShare) > 0 {
		i -= len(m.KeyShare)
		copy(dAtA[i:], m.KeyShare)
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(len(m.KeyShare)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IdValue) > 0 {
		i -= len(m.IdValue)
		copy(dAtA[i:], m.IdValue)
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(len(m.IdValue)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IdType) > 0 {
		i -= len(m.IdType)
		copy(dAtA[i:], m.IdType)
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(len(m.IdType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintGeneralKeyShare(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGeneralKeyShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovGeneralKeyShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GeneralKeyShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovGeneralKeyShare(uint64(l))
	}
	l = len(m.IdType)
	if l > 0 {
		n += 1 + l + sovGeneralKeyShare(uint64(l))
	}
	l = len(m.IdValue)
	if l > 0 {
		n += 1 + l + sovGeneralKeyShare(uint64(l))
	}
	l = len(m.KeyShare)
	if l > 0 {
		n += 1 + l + sovGeneralKeyShare(uint64(l))
	}
	if m.KeyShareIndex != 0 {
		n += 1 + sovGeneralKeyShare(uint64(m.KeyShareIndex))
	}
	if m.ReceivedTimestamp != 0 {
		n += 1 + sovGeneralKeyShare(uint64(m.ReceivedTimestamp))
	}
	if m.ReceivedBlockHeight != 0 {
		n += 1 + sovGeneralKeyShare(uint64(m.ReceivedBlockHeight))
	}
	return n
}

func sovGeneralKeyShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGeneralKeyShare(x uint64) (n int) {
	return sovGeneralKeyShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GeneralKeyShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeneralKeyShare
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
			return fmt.Errorf("proto: GeneralKeyShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeneralKeyShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
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
				return ErrInvalidLengthGeneralKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGeneralKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
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
				return ErrInvalidLengthGeneralKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGeneralKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
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
				return ErrInvalidLengthGeneralKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGeneralKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
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
				return ErrInvalidLengthGeneralKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGeneralKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyShare = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShareIndex", wireType)
			}
			m.KeyShareIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyShareIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedTimestamp", wireType)
			}
			m.ReceivedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedBlockHeight", wireType)
			}
			m.ReceivedBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeneralKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGeneralKeyShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGeneralKeyShare
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
func skipGeneralKeyShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGeneralKeyShare
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
					return 0, ErrIntOverflowGeneralKeyShare
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
					return 0, ErrIntOverflowGeneralKeyShare
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
				return 0, ErrInvalidLengthGeneralKeyShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGeneralKeyShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGeneralKeyShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGeneralKeyShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGeneralKeyShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGeneralKeyShare = fmt.Errorf("proto: unexpected end of group")
)
