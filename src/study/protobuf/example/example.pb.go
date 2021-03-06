// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	CmdLogin
	CmdModifyInfo
	Test
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CmdLogin struct {
	UserName         *string `protobuf:"bytes,1,req,name=UserName" json:"UserName,omitempty"`
	Passwd           *string `protobuf:"bytes,2,req,name=Passwd" json:"Passwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdLogin) Reset()                    { *m = CmdLogin{} }
func (m *CmdLogin) String() string            { return proto.CompactTextString(m) }
func (*CmdLogin) ProtoMessage()               {}
func (*CmdLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CmdLogin) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *CmdLogin) GetPasswd() string {
	if m != nil && m.Passwd != nil {
		return *m.Passwd
	}
	return ""
}

type CmdModifyInfo struct {
	UserName         *string `protobuf:"bytes,1,req,name=UserName" json:"UserName,omitempty"`
	NewPwd           *string `protobuf:"bytes,2,req,name=NewPwd" json:"NewPwd,omitempty"`
	NickName         *string `protobuf:"bytes,3,req,name=NickName" json:"NickName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdModifyInfo) Reset()                    { *m = CmdModifyInfo{} }
func (m *CmdModifyInfo) String() string            { return proto.CompactTextString(m) }
func (*CmdModifyInfo) ProtoMessage()               {}
func (*CmdModifyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CmdModifyInfo) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *CmdModifyInfo) GetNewPwd() string {
	if m != nil && m.NewPwd != nil {
		return *m.NewPwd
	}
	return ""
}

func (m *CmdModifyInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

type Test struct {
	CmdType          *int32         `protobuf:"varint,1,req,name=CmdType" json:"CmdType,omitempty"`
	Login            *CmdLogin      `protobuf:"bytes,2,opt,name=Login" json:"Login,omitempty"`
	ModifyInfo       *CmdModifyInfo `protobuf:"bytes,3,opt,name=ModifyInfo" json:"ModifyInfo,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Test) GetCmdType() int32 {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return 0
}

func (m *Test) GetLogin() *CmdLogin {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *Test) GetModifyInfo() *CmdModifyInfo {
	if m != nil {
		return m.ModifyInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdLogin)(nil), "example.CmdLogin")
	proto.RegisterType((*CmdModifyInfo)(nil), "example.CmdModifyInfo")
	proto.RegisterType((*Test)(nil), "example.Test")
}

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x74, 0xb8,
	0x38, 0x9c, 0x73, 0x53, 0x7c, 0xf2, 0xd3, 0x33, 0xf3, 0x84, 0x04, 0xb8, 0x38, 0x42, 0x8b, 0x53,
	0x8b, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x85, 0xf8, 0xb8, 0xd8, 0x02,
	0x12, 0x8b, 0x8b, 0xcb, 0x53, 0x24, 0x98, 0x40, 0x7c, 0x25, 0x67, 0x2e, 0x5e, 0xa0, 0x6a, 0xdf,
	0xfc, 0x94, 0xcc, 0xb4, 0x4a, 0xcf, 0xbc, 0xb4, 0x7c, 0xec, 0x5a, 0xfc, 0x52, 0xcb, 0x03, 0x60,
	0x5a, 0x40, 0x2a, 0xfc, 0x32, 0x93, 0xb3, 0xc1, 0x2a, 0x98, 0xc1, 0x86, 0xa4, 0x72, 0xb1, 0x84,
	0xa4, 0x16, 0x97, 0x08, 0xf1, 0x73, 0xb1, 0x03, 0x0d, 0x0b, 0xa9, 0x2c, 0x80, 0x68, 0x65, 0x15,
	0x52, 0xe0, 0x62, 0x05, 0x3b, 0x04, 0xa8, 0x93, 0x51, 0x83, 0xdb, 0x48, 0x50, 0x0f, 0xe6, 0x66,
	0xb8, 0x0b, 0xb5, 0xb8, 0xb8, 0x10, 0x96, 0x03, 0x8d, 0x03, 0x29, 0x13, 0x43, 0x56, 0x86, 0x90,
	0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x7d, 0x54, 0x24, 0xf2, 0x00, 0x00, 0x00,
}
