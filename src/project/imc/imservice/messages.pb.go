// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package imservice is a generated protocol buffer package.

=========================================================

all messages used by the current chat server and client

=========================================================

It is generated from these files:
	messages.proto

It has these top-level messages:
	CmdLogin
	CmdAckCommon
	CmdModifyInfo
	CmdLogout
	CmdSendMsg
	CmdMessage
	ImcCmd
*/
package imservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMD_TYPE int32

const (
	CMD_TYPE_LOGIN          CMD_TYPE = 1
	CMD_TYPE_MODIFYINFO     CMD_TYPE = 2
	CMD_TYPE_LOGOUT         CMD_TYPE = 3
	CMD_TYPE_SENDMSG        CMD_TYPE = 4
	CMD_TYPE_MESSAGE        CMD_TYPE = 5
	CMD_TYPE_LOGIN_ACK      CMD_TYPE = 101
	CMD_TYPE_MODIFYINFO_ACK CMD_TYPE = 102
	CMD_TYPE_LOGOUT_ACK     CMD_TYPE = 103
	CMD_TYPE_SENDMSG_ACK    CMD_TYPE = 104
)

var CMD_TYPE_name = map[int32]string{
	1:   "LOGIN",
	2:   "MODIFYINFO",
	3:   "LOGOUT",
	4:   "SENDMSG",
	5:   "MESSAGE",
	101: "LOGIN_ACK",
	102: "MODIFYINFO_ACK",
	103: "LOGOUT_ACK",
	104: "SENDMSG_ACK",
}
var CMD_TYPE_value = map[string]int32{
	"LOGIN":          1,
	"MODIFYINFO":     2,
	"LOGOUT":         3,
	"SENDMSG":        4,
	"MESSAGE":        5,
	"LOGIN_ACK":      101,
	"MODIFYINFO_ACK": 102,
	"LOGOUT_ACK":     103,
	"SENDMSG_ACK":    104,
}

func (x CMD_TYPE) Enum() *CMD_TYPE {
	p := new(CMD_TYPE)
	*p = x
	return p
}
func (x CMD_TYPE) String() string {
	return proto.EnumName(CMD_TYPE_name, int32(x))
}
func (x *CMD_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMD_TYPE_value, data, "CMD_TYPE")
	if err != nil {
		return err
	}
	*x = CMD_TYPE(value)
	return nil
}
func (CMD_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RET_CODE int32

const (
	RET_CODE_SUCCESS RET_CODE = 1
	RET_CODE_FAILED  RET_CODE = 2
)

var RET_CODE_name = map[int32]string{
	1: "SUCCESS",
	2: "FAILED",
}
var RET_CODE_value = map[string]int32{
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x RET_CODE) Enum() *RET_CODE {
	p := new(RET_CODE)
	*p = x
	return p
}
func (x RET_CODE) String() string {
	return proto.EnumName(RET_CODE_name, int32(x))
}
func (x *RET_CODE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RET_CODE_value, data, "RET_CODE")
	if err != nil {
		return err
	}
	*x = RET_CODE(value)
	return nil
}
func (RET_CODE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type CmdAckCommon struct {
	Status           *RET_CODE `protobuf:"varint,1,req,name=Status,enum=imservice.RET_CODE" json:"Status,omitempty"`
	ErrorDesc        *string   `protobuf:"bytes,2,opt,name=ErrorDesc" json:"ErrorDesc,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *CmdAckCommon) Reset()                    { *m = CmdAckCommon{} }
func (m *CmdAckCommon) String() string            { return proto.CompactTextString(m) }
func (*CmdAckCommon) ProtoMessage()               {}
func (*CmdAckCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CmdAckCommon) GetStatus() RET_CODE {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return RET_CODE_SUCCESS
}

func (m *CmdAckCommon) GetErrorDesc() string {
	if m != nil && m.ErrorDesc != nil {
		return *m.ErrorDesc
	}
	return ""
}

type CmdModifyInfo struct {
	UserName         *string `protobuf:"bytes,1,req,name=UserName" json:"UserName,omitempty"`
	NewPasswd        *string `protobuf:"bytes,2,req,name=NewPasswd" json:"NewPasswd,omitempty"`
	NickName         *string `protobuf:"bytes,3,req,name=NickName" json:"NickName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdModifyInfo) Reset()                    { *m = CmdModifyInfo{} }
func (m *CmdModifyInfo) String() string            { return proto.CompactTextString(m) }
func (*CmdModifyInfo) ProtoMessage()               {}
func (*CmdModifyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CmdModifyInfo) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *CmdModifyInfo) GetNewPasswd() string {
	if m != nil && m.NewPasswd != nil {
		return *m.NewPasswd
	}
	return ""
}

func (m *CmdModifyInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

type CmdLogout struct {
	UserName         *string `protobuf:"bytes,1,req,name=UserName" json:"UserName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdLogout) Reset()                    { *m = CmdLogout{} }
func (m *CmdLogout) String() string            { return proto.CompactTextString(m) }
func (*CmdLogout) ProtoMessage()               {}
func (*CmdLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CmdLogout) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

type CmdSendMsg struct {
	PeerName         *string `protobuf:"bytes,1,req,name=PeerName" json:"PeerName,omitempty"`
	MsgBody          *string `protobuf:"bytes,2,req,name=MsgBody" json:"MsgBody,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdSendMsg) Reset()                    { *m = CmdSendMsg{} }
func (m *CmdSendMsg) String() string            { return proto.CompactTextString(m) }
func (*CmdSendMsg) ProtoMessage()               {}
func (*CmdSendMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CmdSendMsg) GetPeerName() string {
	if m != nil && m.PeerName != nil {
		return *m.PeerName
	}
	return ""
}

func (m *CmdSendMsg) GetMsgBody() string {
	if m != nil && m.MsgBody != nil {
		return *m.MsgBody
	}
	return ""
}

type CmdMessage struct {
	From             *string `protobuf:"bytes,1,req,name=From" json:"From,omitempty"`
	To               *string `protobuf:"bytes,2,req,name=To" json:"To,omitempty"`
	MsgBody          *string `protobuf:"bytes,3,req,name=MsgBody" json:"MsgBody,omitempty"`
	Datetime         *int64  `protobuf:"varint,4,req,name=Datetime" json:"Datetime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdMessage) Reset()                    { *m = CmdMessage{} }
func (m *CmdMessage) String() string            { return proto.CompactTextString(m) }
func (*CmdMessage) ProtoMessage()               {}
func (*CmdMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CmdMessage) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *CmdMessage) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

func (m *CmdMessage) GetMsgBody() string {
	if m != nil && m.MsgBody != nil {
		return *m.MsgBody
	}
	return ""
}

func (m *CmdMessage) GetDatetime() int64 {
	if m != nil && m.Datetime != nil {
		return *m.Datetime
	}
	return 0
}

type ImcCmd struct {
	CmdType          *CMD_TYPE      `protobuf:"varint,1,req,name=CmdType,enum=imservice.CMD_TYPE" json:"CmdType,omitempty"`
	Login            *CmdLogin      `protobuf:"bytes,2,opt,name=Login" json:"Login,omitempty"`
	ModifyInfo       *CmdModifyInfo `protobuf:"bytes,3,opt,name=ModifyInfo" json:"ModifyInfo,omitempty"`
	AckCommon        *CmdAckCommon  `protobuf:"bytes,4,opt,name=AckCommon" json:"AckCommon,omitempty"`
	Logout           *CmdLogout     `protobuf:"bytes,5,opt,name=Logout" json:"Logout,omitempty"`
	SendMsg          *CmdSendMsg    `protobuf:"bytes,6,opt,name=SendMsg" json:"SendMsg,omitempty"`
	Message          *CmdMessage    `protobuf:"bytes,7,opt,name=Message" json:"Message,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ImcCmd) Reset()                    { *m = ImcCmd{} }
func (m *ImcCmd) String() string            { return proto.CompactTextString(m) }
func (*ImcCmd) ProtoMessage()               {}
func (*ImcCmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ImcCmd) GetCmdType() CMD_TYPE {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CMD_TYPE_LOGIN
}

func (m *ImcCmd) GetLogin() *CmdLogin {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *ImcCmd) GetModifyInfo() *CmdModifyInfo {
	if m != nil {
		return m.ModifyInfo
	}
	return nil
}

func (m *ImcCmd) GetAckCommon() *CmdAckCommon {
	if m != nil {
		return m.AckCommon
	}
	return nil
}

func (m *ImcCmd) GetLogout() *CmdLogout {
	if m != nil {
		return m.Logout
	}
	return nil
}

func (m *ImcCmd) GetSendMsg() *CmdSendMsg {
	if m != nil {
		return m.SendMsg
	}
	return nil
}

func (m *ImcCmd) GetMessage() *CmdMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdLogin)(nil), "imservice.CmdLogin")
	proto.RegisterType((*CmdAckCommon)(nil), "imservice.CmdAckCommon")
	proto.RegisterType((*CmdModifyInfo)(nil), "imservice.CmdModifyInfo")
	proto.RegisterType((*CmdLogout)(nil), "imservice.CmdLogout")
	proto.RegisterType((*CmdSendMsg)(nil), "imservice.CmdSendMsg")
	proto.RegisterType((*CmdMessage)(nil), "imservice.CmdMessage")
	proto.RegisterType((*ImcCmd)(nil), "imservice.ImcCmd")
	proto.RegisterEnum("imservice.CMD_TYPE", CMD_TYPE_name, CMD_TYPE_value)
	proto.RegisterEnum("imservice.RET_CODE", RET_CODE_name, RET_CODE_value)
}

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6e, 0x9b, 0x40,
	0x10, 0x85, 0x65, 0x6c, 0x63, 0x33, 0x4e, 0x5c, 0xba, 0x6d, 0x55, 0x6e, 0x2a, 0x55, 0x24, 0xaa,
	0x2a, 0x2b, 0x72, 0xa5, 0xbe, 0x81, 0xcb, 0x8f, 0x85, 0x1a, 0xc0, 0x2a, 0xf8, 0x22, 0x57, 0x16,
	0x82, 0xb5, 0x8b, 0xa2, 0xf5, 0x46, 0x2c, 0x69, 0xe4, 0xa7, 0xe8, 0x3b, 0xf4, 0x49, 0x3b, 0xbb,
	0xc6, 0x3f, 0x28, 0xbe, 0x42, 0x3b, 0xf3, 0xed, 0x99, 0xc3, 0x99, 0x85, 0x31, 0xa3, 0x42, 0x64,
	0x1b, 0x2a, 0xa6, 0x4f, 0x15, 0xaf, 0x39, 0x31, 0x4a, 0x26, 0x68, 0xf5, 0xa7, 0xcc, 0xa9, 0x7d,
	0x07, 0x43, 0x87, 0x15, 0xf7, 0x7c, 0x53, 0x6e, 0x89, 0x09, 0xc3, 0x25, 0xd6, 0xa3, 0x8c, 0x51,
	0xab, 0xf3, 0x59, 0xfb, 0x6a, 0x90, 0x31, 0xe8, 0x8b, 0x4c, 0x88, 0x97, 0xc2, 0xd2, 0xe4, 0xd9,
	0xf6, 0xe1, 0x0a, 0xe9, 0x59, 0xfe, 0xe8, 0x70, 0xc6, 0xf8, 0x96, 0xdc, 0x80, 0x9e, 0xd4, 0x59,
	0xfd, 0x2c, 0x14, 0x3f, 0xfe, 0xfe, 0x6e, 0x7a, 0x54, 0x9e, 0xfe, 0xf2, 0xd2, 0x95, 0x13, 0xbb,
	0x1e, 0x79, 0x0b, 0x86, 0x57, 0x55, 0xbc, 0x72, 0xa9, 0xc8, 0x51, 0xa7, 0xa3, 0x74, 0xae, 0x51,
	0x27, 0xe4, 0x45, 0xb9, 0xde, 0x05, 0xdb, 0x35, 0xbf, 0x30, 0x1a, 0x6f, 0x45, 0xf4, 0xe5, 0x7c,
	0xba, 0x84, 0xa2, 0x32, 0x7f, 0x54, 0x50, 0x57, 0xf9, 0xf9, 0x04, 0xc6, 0xde, 0x3d, 0x7f, 0xae,
	0x5f, 0x6b, 0xd8, 0xdf, 0x00, 0xb0, 0x9d, 0xd0, 0x6d, 0x11, 0x8a, 0x8d, 0xec, 0x2f, 0x68, 0x6b,
	0xc6, 0x1b, 0x18, 0x60, 0xe3, 0x07, 0x2f, 0x76, 0xcd, 0xff, 0x05, 0xea, 0x42, 0xb8, 0x4f, 0x8b,
	0x5c, 0x41, 0xcf, 0xaf, 0x38, 0x6b, 0x60, 0x00, 0x2d, 0xe5, 0x8d, 0x93, 0xb3, 0x8b, 0xdd, 0x83,
	0x35, 0x37, 0xab, 0x69, 0x5d, 0xa2, 0x76, 0x0f, 0x2b, 0x5d, 0xfb, 0x9f, 0x06, 0x7a, 0xc0, 0x72,
	0x94, 0x23, 0xb7, 0x30, 0xc0, 0x4f, 0xba, 0x7b, 0xa2, 0x17, 0x62, 0x72, 0x42, 0x77, 0x95, 0x3e,
	0x2c, 0x3c, 0x62, 0x43, 0x5f, 0xad, 0x41, 0x45, 0x34, 0x6a, 0x33, 0x87, 0x0d, 0xdd, 0x01, 0x9c,
	0x42, 0xc3, 0xd1, 0x12, 0xb4, 0xda, 0xe0, 0x59, 0xa8, 0x13, 0x30, 0x8e, 0xab, 0x42, 0x57, 0x12,
	0xfe, 0xd8, 0x86, 0x4f, 0x9b, 0xbc, 0x05, 0x7d, 0x1f, 0xa3, 0xd5, 0x57, 0xe0, 0xfb, 0x57, 0xe3,
	0x65, 0xc4, 0x5f, 0x60, 0xd0, 0xa4, 0x69, 0xe9, 0x0a, 0xfb, 0xd0, 0xc6, 0x0e, 0x51, 0x23, 0xd7,
	0x84, 0x68, 0x0d, 0x2e, 0x71, 0x4d, 0x73, 0xf2, 0xb7, 0x83, 0xcf, 0xef, 0x10, 0x80, 0x81, 0x01,
	0xc4, 0xf3, 0x20, 0x32, 0x3b, 0xf8, 0xee, 0x20, 0x8c, 0xdd, 0xc0, 0x7f, 0x08, 0x22, 0x3f, 0x36,
	0x35, 0xcc, 0x5e, 0xc7, 0x56, 0xbc, 0x4c, 0xcd, 0x2e, 0x19, 0xa1, 0x07, 0x2f, 0x72, 0xc3, 0x64,
	0x6e, 0xf6, 0xe4, 0x21, 0xf4, 0x92, 0x64, 0x36, 0xf7, 0xcc, 0x3e, 0xb9, 0x06, 0x43, 0x09, 0xac,
	0x66, 0xce, 0x4f, 0x93, 0x12, 0x02, 0xe3, 0x93, 0x88, 0xaa, 0xad, 0xa5, 0xf0, 0x5e, 0x48, 0x9d,
	0x37, 0xb8, 0xc8, 0x51, 0x23, 0xa6, 0x0a, 0xbf, 0x27, 0x37, 0x30, 0x3c, 0x3e, 0x5c, 0x39, 0x69,
	0xe9, 0x38, 0xa8, 0x8f, 0x96, 0xd0, 0x82, 0x3f, 0x0b, 0xee, 0x3d, 0xd7, 0xd4, 0xfe, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x7b, 0xf7, 0x9c, 0x50, 0x03, 0x00, 0x00,
}
