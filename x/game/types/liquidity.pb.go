// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmichorizon/qwoyn/game/liquidity.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Liquidity struct {
	Amounts []types.Coin `protobuf:"bytes,1,rep,name=amounts,proto3" json:"amounts"`
}

func (m *Liquidity) Reset()         { *m = Liquidity{} }
func (m *Liquidity) String() string { return proto.CompactTextString(m) }
func (*Liquidity) ProtoMessage()    {}
func (*Liquidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_91c84b5ca6a1b13f, []int{0}
}
func (m *Liquidity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Liquidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Liquidity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Liquidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Liquidity.Merge(m, src)
}
func (m *Liquidity) XXX_Size() int {
	return m.Size()
}
func (m *Liquidity) XXX_DiscardUnknown() {
	xxx_messageInfo_Liquidity.DiscardUnknown(m)
}

var xxx_messageInfo_Liquidity proto.InternalMessageInfo

func (m *Liquidity) GetAmounts() []types.Coin {
	if m != nil {
		return m.Amounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Liquidity)(nil), "cosmichorizon.qwoyn.game.Liquidity")
}

func init() {
	proto.RegisterFile("cosmichorizon/qwoyn/game/liquidity.proto", fileDescriptor_91c84b5ca6a1b13f)
}

var fileDescriptor_91c84b5ca6a1b13f = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x13, 0x81, 0x40, 0x84, 0xad, 0x62, 0x28, 0x1d, 0x5c, 0xc4, 0xd4, 0x01, 0xfc, 0x54,
	0x98, 0x58, 0x8b, 0xd4, 0x89, 0x89, 0x91, 0xcd, 0x4e, 0x8d, 0xfb, 0xa4, 0xda, 0x2f, 0x8d, 0x6d,
	0x20, 0x7c, 0x05, 0x9f, 0xd5, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x8f, 0xa0, 0xc4, 0xcd, 0x10, 0xb6,
	0x6b, 0xdf, 0x63, 0xe9, 0xf8, 0x66, 0xb3, 0x9c, 0x9c, 0xc1, 0x7c, 0x4d, 0x25, 0x7e, 0x92, 0x85,
	0xed, 0x3b, 0x55, 0x16, 0xb4, 0x30, 0x0a, 0x36, 0xb8, 0x0d, 0xb8, 0x42, 0x5f, 0xf1, 0xa2, 0x24,
	0x4f, 0xa3, 0xf1, 0x80, 0xe4, 0x1d, 0xc9, 0x5b, 0x72, 0xc2, 0xda, 0x86, 0x1c, 0x48, 0xe1, 0x14,
	0xbc, 0xcd, 0xa5, 0xf2, 0x62, 0x0e, 0x39, 0xa1, 0x8d, 0x2f, 0x27, 0x4c, 0x13, 0xe9, 0x8d, 0x82,
	0xee, 0x24, 0xc3, 0x2b, 0xac, 0x42, 0x29, 0x3c, 0x52, 0xdf, 0x4f, 0xff, 0xf7, 0x1e, 0x8d, 0x72,
	0x5e, 0x98, 0xe2, 0x00, 0x5c, 0x68, 0xd2, 0xd4, 0x45, 0x68, 0x53, 0xbc, 0xbd, 0x5e, 0x66, 0x67,
	0x4f, 0xbd, 0xe3, 0xe8, 0x21, 0x3b, 0x15, 0x86, 0x82, 0xf5, 0x6e, 0x9c, 0x5e, 0x1d, 0xcd, 0xce,
	0xef, 0x2e, 0x79, 0xb4, 0xe2, 0xad, 0x15, 0x3f, 0x58, 0xf1, 0x47, 0x42, 0xbb, 0x38, 0xde, 0xfd,
	0x4c, 0x93, 0xe7, 0x9e, 0x5f, 0x2c, 0x77, 0x35, 0x4b, 0xf7, 0x35, 0x4b, 0x7f, 0x6b, 0x96, 0x7e,
	0x35, 0x2c, 0xd9, 0x37, 0x2c, 0xf9, 0x6e, 0x58, 0xf2, 0x72, 0xa3, 0xd1, 0xaf, 0x83, 0xe4, 0x39,
	0x19, 0x88, 0xbf, 0xbf, 0x1d, 0x0e, 0xf5, 0x11, 0xa7, 0xf2, 0x55, 0xa1, 0x9c, 0x3c, 0xe9, 0xb4,
	0xee, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xad, 0x7d, 0x32, 0x53, 0x01, 0x00, 0x00,
}

func (m *Liquidity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Liquidity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Liquidity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Liquidity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	return n
}

func sovLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidity(x uint64) (n int) {
	return sovLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Liquidity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: Liquidity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Liquidity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amounts = append(m.Amounts, types.Coin{})
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func skipLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
				return 0, ErrInvalidLengthLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidity = fmt.Errorf("proto: unexpected end of group")
)
