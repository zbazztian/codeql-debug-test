// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/protocol/deviceid_test.proto

package protocol

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/syncthing/syncthing/proto/ext"
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

type TestOldDeviceID struct {
	Test []byte `protobuf:"bytes,1,opt,name=test,proto3" json:"test" xml:"test"`
}

func (m *TestOldDeviceID) Reset()         { *m = TestOldDeviceID{} }
func (m *TestOldDeviceID) String() string { return proto.CompactTextString(m) }
func (*TestOldDeviceID) ProtoMessage()    {}
func (*TestOldDeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a75253a19e48a2, []int{0}
}
func (m *TestOldDeviceID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestOldDeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestOldDeviceID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestOldDeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestOldDeviceID.Merge(m, src)
}
func (m *TestOldDeviceID) XXX_Size() int {
	return m.ProtoSize()
}
func (m *TestOldDeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_TestOldDeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_TestOldDeviceID proto.InternalMessageInfo

type TestNewDeviceID struct {
	Test DeviceID `protobuf:"bytes,1,opt,name=test,proto3,customtype=DeviceID" json:"test" xml:"test"`
}

func (m *TestNewDeviceID) Reset()         { *m = TestNewDeviceID{} }
func (m *TestNewDeviceID) String() string { return proto.CompactTextString(m) }
func (*TestNewDeviceID) ProtoMessage()    {}
func (*TestNewDeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a75253a19e48a2, []int{1}
}
func (m *TestNewDeviceID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestNewDeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestNewDeviceID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestNewDeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestNewDeviceID.Merge(m, src)
}
func (m *TestNewDeviceID) XXX_Size() int {
	return m.ProtoSize()
}
func (m *TestNewDeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_TestNewDeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_TestNewDeviceID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestOldDeviceID)(nil), "protocol.TestOldDeviceID")
	proto.RegisterType((*TestNewDeviceID)(nil), "protocol.TestNewDeviceID")
}

func init() { proto.RegisterFile("lib/protocol/deviceid_test.proto", fileDescriptor_f4a75253a19e48a2) }

var fileDescriptor_f4a75253a19e48a2 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xc9, 0x4c, 0xd2,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xcd,
	0x4c, 0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0x0b, 0x0b, 0x71, 0xc0, 0x64, 0xa5, 0x38, 0x53,
	0x2b, 0xa0, 0x82, 0x52, 0xca, 0x45, 0xa9, 0x05, 0xf9, 0xc5, 0x10, 0x8d, 0x49, 0xa5, 0x69, 0xfa,
	0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xa4, 0x64, 0xcb, 0xc5, 0x1f, 0x92, 0x5a,
	0x5c, 0xe2, 0x9f, 0x93, 0xe2, 0x02, 0x36, 0xd7, 0xd3, 0x45, 0x48, 0x8b, 0x8b, 0x05, 0x64, 0xb4,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xd8, 0xab, 0x7b, 0xf2, 0x60, 0xfe, 0xa7, 0x7b, 0xf2,
	0x5c, 0x15, 0xb9, 0x39, 0x56, 0x4a, 0x20, 0x8e, 0x52, 0x10, 0x58, 0x4c, 0x29, 0x10, 0xa2, 0xdd,
	0x2f, 0xb5, 0x1c, 0xae, 0xdd, 0x0e, 0x45, 0xbb, 0xd6, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9,
	0x73, 0xc0, 0xe4, 0xb1, 0x1b, 0xd7, 0x71, 0x41, 0x85, 0x11, 0x62, 0xa4, 0x93, 0xef, 0x89, 0x87,
	0x72, 0x0c, 0x17, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31, 0x5e, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x94, 0x76, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x65, 0x5e,
	0x72, 0x49, 0x46, 0x66, 0x5e, 0x3a, 0x12, 0x0b, 0x39, 0xb8, 0x92, 0xd8, 0xc0, 0x2c, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x0a, 0x77, 0x43, 0x45, 0x01, 0x00, 0x00,
}

func (m *TestOldDeviceID) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestOldDeviceID) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestOldDeviceID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Test) > 0 {
		i -= len(m.Test)
		copy(dAtA[i:], m.Test)
		i = encodeVarintDeviceidTest(dAtA, i, uint64(len(m.Test)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestNewDeviceID) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestNewDeviceID) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestNewDeviceID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Test.ProtoSize()
		i -= size
		if _, err := m.Test.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeviceidTest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDeviceidTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceidTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestOldDeviceID) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Test)
	if l > 0 {
		n += 1 + l + sovDeviceidTest(uint64(l))
	}
	return n
}

func (m *TestNewDeviceID) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Test.ProtoSize()
	n += 1 + l + sovDeviceidTest(uint64(l))
	return n
}

func sovDeviceidTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceidTest(x uint64) (n int) {
	return sovDeviceidTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestOldDeviceID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceidTest
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
			return fmt.Errorf("proto: TestOldDeviceID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestOldDeviceID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceidTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Test = append(m.Test[:0], dAtA[iNdEx:postIndex]...)
			if m.Test == nil {
				m.Test = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceidTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceidTest
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
func (m *TestNewDeviceID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceidTest
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
			return fmt.Errorf("proto: TestNewDeviceID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestNewDeviceID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceidTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Test.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceidTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceidTest
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
func skipDeviceidTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceidTest
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
					return 0, ErrIntOverflowDeviceidTest
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
					return 0, ErrIntOverflowDeviceidTest
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
				return 0, ErrInvalidLengthDeviceidTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceidTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceidTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceidTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceidTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceidTest = fmt.Errorf("proto: unexpected end of group")
)
