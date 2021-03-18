// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/foldertype.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type FolderType int32

const (
	FolderTypeSendReceive      FolderType = 0
	FolderTypeSendOnly         FolderType = 1
	FolderTypeReceiveOnly      FolderType = 2
	FolderTypeReceiveEncrypted FolderType = 3
)

var FolderType_name = map[int32]string{
	0: "FOLDER_TYPE_SEND_RECEIVE",
	1: "FOLDER_TYPE_SEND_ONLY",
	2: "FOLDER_TYPE_RECEIVE_ONLY",
	3: "FOLDER_TYPE_RECEIVE_ENCRYPTED",
}

var FolderType_value = map[string]int32{
	"FOLDER_TYPE_SEND_RECEIVE":      0,
	"FOLDER_TYPE_SEND_ONLY":         1,
	"FOLDER_TYPE_RECEIVE_ONLY":      2,
	"FOLDER_TYPE_RECEIVE_ENCRYPTED": 3,
}

func (FolderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea6ddb20c0633575, []int{0}
}

func init() {
	proto.RegisterEnum("config.FolderType", FolderType_name, FolderType_value)
}

func init() { proto.RegisterFile("lib/config/foldertype.proto", fileDescriptor_ea6ddb20c0633575) }

var fileDescriptor_ea6ddb20c0633575 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xc9, 0x4c, 0xd2,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcf, 0x49, 0x49, 0x2d, 0x2a, 0xa9, 0x2c,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x48, 0x29, 0x17, 0xa5, 0x16,
	0xe4, 0x17, 0xeb, 0x83, 0x05, 0x93, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x58, 0xeb, 0x17, 0x23, 0x17, 0x97, 0x1b, 0xd8, 0x84, 0x90, 0xca, 0x82, 0x54, 0x21,
	0x73, 0x2e, 0x09, 0x37, 0x7f, 0x1f, 0x17, 0xd7, 0xa0, 0xf8, 0x90, 0xc8, 0x00, 0xd7, 0xf8, 0x60,
	0x57, 0x3f, 0x97, 0xf8, 0x20, 0x57, 0x67, 0x57, 0xcf, 0x30, 0x57, 0x01, 0x06, 0x29, 0xc9, 0xae,
	0xb9, 0x0a, 0xa2, 0x08, 0xd5, 0xc1, 0xa9, 0x79, 0x29, 0x41, 0xa9, 0xc9, 0xa9, 0x99, 0x65, 0xa9,
	0x42, 0x86, 0x5c, 0xa2, 0x18, 0x1a, 0xfd, 0xfd, 0x7c, 0x22, 0x05, 0x18, 0xa5, 0xc4, 0xba, 0xe6,
	0x2a, 0x08, 0xa1, 0xea, 0xf2, 0xcf, 0xcb, 0xa9, 0x44, 0xb7, 0x0b, 0x6a, 0x0d, 0x44, 0x17, 0x13,
	0xba, 0x5d, 0x50, 0x7b, 0xc0, 0x1a, 0x1d, 0xb9, 0x64, 0xb1, 0x69, 0x74, 0xf5, 0x73, 0x0e, 0x8a,
	0x0c, 0x08, 0x71, 0x75, 0x11, 0x60, 0x96, 0x92, 0xeb, 0x9a, 0xab, 0x20, 0x85, 0xa1, 0xdb, 0x35,
	0x2f, 0xb9, 0xa8, 0xb2, 0xa0, 0x24, 0x35, 0x45, 0x8a, 0x65, 0xc5, 0x12, 0x39, 0x06, 0x27, 0xef,
	0x13, 0x0f, 0xe5, 0x18, 0x2e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x16, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2,
	0xca, 0xbc, 0xe4, 0x92, 0x8c, 0xcc, 0xbc, 0x74, 0x24, 0x16, 0x22, 0x1e, 0x92, 0xd8, 0xc0, 0x01,
	0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x87, 0xbe, 0x2d, 0x9c, 0x01, 0x00, 0x00,
}
