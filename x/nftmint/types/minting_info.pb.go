// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: titan/nftmint/minting_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// MintingInfo defines the minting info for a class.
type MintingInfo struct {
	// class_id is a unique identifier of the class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// owner is the owner address of the class.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// next_token_id is the unique identifier of the next token that will be
	// minted under this class.
	NextTokenId uint64 `protobuf:"varint,3,opt,name=next_token_id,json=nextTokenId,proto3" json:"next_token_id,omitempty"`
}

func (m *MintingInfo) Reset()         { *m = MintingInfo{} }
func (m *MintingInfo) String() string { return proto.CompactTextString(m) }
func (*MintingInfo) ProtoMessage()    {}
func (*MintingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_81d80c867f9b0232, []int{0}
}
func (m *MintingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintingInfo.Merge(m, src)
}
func (m *MintingInfo) XXX_Size() int {
	return m.Size()
}
func (m *MintingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MintingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MintingInfo proto.InternalMessageInfo

func (m *MintingInfo) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MintingInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MintingInfo) GetNextTokenId() uint64 {
	if m != nil {
		return m.NextTokenId
	}
	return 0
}

func init() {
	proto.RegisterType((*MintingInfo)(nil), "titan.nftmint.MintingInfo")
}

func init() { proto.RegisterFile("titan/nftmint/minting_info.proto", fileDescriptor_81d80c867f9b0232) }

var fileDescriptor_81d80c867f9b0232 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc9, 0x2c, 0x49,
	0xcc, 0xd3, 0xcf, 0x4b, 0x2b, 0xc9, 0xcd, 0xcc, 0x2b, 0xd1, 0x07, 0x11, 0x99, 0x79, 0xe9, 0xf1,
	0x99, 0x79, 0x69, 0xf9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x60, 0x15, 0x7a, 0x50,
	0x15, 0x52, 0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x60, 0x49, 0x7d, 0x08, 0x07, 0xa2,
	0x52, 0xa9, 0x86, 0x8b, 0xdb, 0x17, 0xa2, 0xdf, 0x33, 0x2f, 0x2d, 0x5f, 0x48, 0x92, 0x8b, 0x23,
	0x39, 0x27, 0xb1, 0xb8, 0x38, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x1d,
	0xcc, 0xf7, 0x4c, 0x11, 0xd2, 0xe3, 0x62, 0xcd, 0x2f, 0xcf, 0x4b, 0x2d, 0x92, 0x60, 0x02, 0x89,
	0x3b, 0x49, 0x5c, 0xda, 0xa2, 0x2b, 0x02, 0x35, 0xca, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38,
	0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa2, 0x4c, 0x48, 0x89, 0x8b, 0x37, 0x2f, 0xb5, 0xa2,
	0x24, 0xbe, 0x24, 0x3f, 0x3b, 0x35, 0x0f, 0x64, 0x1e, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x37,
	0x48, 0x30, 0x04, 0x24, 0xe6, 0x99, 0xe2, 0xe4, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x60,
	0xcf, 0x94, 0x64, 0x57, 0x40, 0x18, 0xfa, 0x15, 0x70, 0x9f, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x7d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x6f, 0xac, 0x3d, 0x17, 0x01,
	0x00, 0x00,
}

func (m *MintingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextTokenId != 0 {
		i = encodeVarintMintingInfo(dAtA, i, uint64(m.NextTokenId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMintingInfo(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintMintingInfo(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMintingInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovMintingInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovMintingInfo(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMintingInfo(uint64(l))
	}
	if m.NextTokenId != 0 {
		n += 1 + sovMintingInfo(uint64(m.NextTokenId))
	}
	return n
}

func sovMintingInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMintingInfo(x uint64) (n int) {
	return sovMintingInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMintingInfo
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
			return fmt.Errorf("proto: MintingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintingInfo
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
				return ErrInvalidLengthMintingInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintingInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintingInfo
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
				return ErrInvalidLengthMintingInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintingInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextTokenId", wireType)
			}
			m.NextTokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintingInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextTokenId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMintingInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMintingInfo
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
func skipMintingInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMintingInfo
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
					return 0, ErrIntOverflowMintingInfo
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
					return 0, ErrIntOverflowMintingInfo
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
				return 0, ErrInvalidLengthMintingInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMintingInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMintingInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMintingInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMintingInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMintingInfo = fmt.Errorf("proto: unexpected end of group")
)
