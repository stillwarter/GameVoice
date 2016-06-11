// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Handshake
	HandshakeAck
	EnterGroup
	EnterGroupAck
	NotifyEnterGroup
	NotifyLeaveGroup
	LeaveGroup
	AdjustPlayerAttr
	Request
	Response
	Noitfy
	ControlMessage
	PlayerRequestConnect
	PlayerRequestConnectAck
	VoiceData
	VoiceHeartbeat
	VoiceMessage
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MSG_CONST int32

const (
	MSG_CONST_CONST_VALUE           MSG_CONST = 0
	MSG_CONST_MAX_PACKET_SIZE       MSG_CONST = 1024
	MSG_CONST_MAX_VOICE_PACKET_SIZE MSG_CONST = 1024
)

var MSG_CONST_name = map[int32]string{
	0:    "CONST_VALUE",
	1024: "MAX_PACKET_SIZE",
	// Duplicate value: 1024: "MAX_VOICE_PACKET_SIZE",
}
var MSG_CONST_value = map[string]int32{
	"CONST_VALUE":           0,
	"MAX_PACKET_SIZE":       1024,
	"MAX_VOICE_PACKET_SIZE": 1024,
}

func (x MSG_CONST) String() string {
	return proto.EnumName(MSG_CONST_name, int32(x))
}
func (MSG_CONST) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MSG int32

const (
	MSG_NONE               MSG = 0
	MSG_HANDSHAKE          MSG = 1
	MSG_HANDSHAKE_ACK      MSG = 2
	MSG_ENTER_GROUP        MSG = 3
	MSG_ENTER_GROUP_ACK    MSG = 4
	MSG_NOTIFY_ENTER_GROUP MSG = 5
	MSG_NOTIFY_LEAVE_GROUP MSG = 6
	MSG_LEAVE_GROUP        MSG = 7
	MSG_ADJUST_PLAYER_ATTR MSG = 8
	MSG_ADJUST_GROUP_ATTR  MSG = 9
)

var MSG_name = map[int32]string{
	0: "NONE",
	1: "HANDSHAKE",
	2: "HANDSHAKE_ACK",
	3: "ENTER_GROUP",
	4: "ENTER_GROUP_ACK",
	5: "NOTIFY_ENTER_GROUP",
	6: "NOTIFY_LEAVE_GROUP",
	7: "LEAVE_GROUP",
	8: "ADJUST_PLAYER_ATTR",
	9: "ADJUST_GROUP_ATTR",
}
var MSG_value = map[string]int32{
	"NONE":               0,
	"HANDSHAKE":          1,
	"HANDSHAKE_ACK":      2,
	"ENTER_GROUP":        3,
	"ENTER_GROUP_ACK":    4,
	"NOTIFY_ENTER_GROUP": 5,
	"NOTIFY_LEAVE_GROUP": 6,
	"LEAVE_GROUP":        7,
	"ADJUST_PLAYER_ATTR": 8,
	"ADJUST_GROUP_ATTR":  9,
}

func (x MSG) String() string {
	return proto.EnumName(MSG_name, int32(x))
}
func (MSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ACK int32

const (
	ACK_OK                ACK = 0
	ACK_HANDSHAKE_FAILED  ACK = 1
	ACK_ADD_PLAYER_FAILED ACK = 2
)

var ACK_name = map[int32]string{
	0: "OK",
	1: "HANDSHAKE_FAILED",
	2: "ADD_PLAYER_FAILED",
}
var ACK_value = map[string]int32{
	"OK":                0,
	"HANDSHAKE_FAILED":  1,
	"ADD_PLAYER_FAILED": 2,
}

func (x ACK) String() string {
	return proto.EnumName(ACK_name, int32(x))
}
func (ACK) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type VOICE_MSG int32

const (
	VOICE_MSG_VOICE_MSG_NONE      VOICE_MSG = 0
	VOICE_MSG_REQEUST_CONNECT     VOICE_MSG = 1
	VOICE_MSG_REQEUST_CONNECT_ACK VOICE_MSG = 2
	VOICE_MSG_VOICE_DATA          VOICE_MSG = 3
	VOICE_MSG_VOICE_HEARTBEAT     VOICE_MSG = 4
)

var VOICE_MSG_name = map[int32]string{
	0: "VOICE_MSG_NONE",
	1: "REQEUST_CONNECT",
	2: "REQEUST_CONNECT_ACK",
	3: "VOICE_DATA",
	4: "VOICE_HEARTBEAT",
}
var VOICE_MSG_value = map[string]int32{
	"VOICE_MSG_NONE":      0,
	"REQEUST_CONNECT":     1,
	"REQEUST_CONNECT_ACK": 2,
	"VOICE_DATA":          3,
	"VOICE_HEARTBEAT":     4,
}

func (x VOICE_MSG) String() string {
	return proto.EnumName(VOICE_MSG_name, int32(x))
}
func (VOICE_MSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type VOICE_ACK int32

const (
	VOICE_ACK_VOICE_OK VOICE_ACK = 0
)

var VOICE_ACK_name = map[int32]string{
	0: "VOICE_OK",
}
var VOICE_ACK_value = map[string]int32{
	"VOICE_OK": 0,
}

func (x VOICE_ACK) String() string {
	return proto.EnumName(VOICE_ACK_name, int32(x))
}
func (VOICE_ACK) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Handshake struct {
	ProductKey string `protobuf:"bytes,2,opt,name=ProductKey,json=productKey" json:"ProductKey,omitempty"`
}

func (m *Handshake) Reset()                    { *m = Handshake{} }
func (m *Handshake) String() string            { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()               {}
func (*Handshake) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HandshakeAck struct {
}

func (m *HandshakeAck) Reset()                    { *m = HandshakeAck{} }
func (m *HandshakeAck) String() string            { return proto.CompactTextString(m) }
func (*HandshakeAck) ProtoMessage()               {}
func (*HandshakeAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EnterGroup struct {
	GroupId  uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
	CanSpeak uint32 `protobuf:"varint,3,opt,name=canSpeak" json:"canSpeak,omitempty"`
}

func (m *EnterGroup) Reset()                    { *m = EnterGroup{} }
func (m *EnterGroup) String() string            { return proto.CompactTextString(m) }
func (*EnterGroup) ProtoMessage()               {}
func (*EnterGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type EnterGroupAck struct {
	GroupId     uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId    uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
	PlayerIndex uint32 `protobuf:"varint,3,opt,name=playerIndex" json:"playerIndex,omitempty"`
}

func (m *EnterGroupAck) Reset()                    { *m = EnterGroupAck{} }
func (m *EnterGroupAck) String() string            { return proto.CompactTextString(m) }
func (*EnterGroupAck) ProtoMessage()               {}
func (*EnterGroupAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NotifyEnterGroup struct {
	GroupId  uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
	CanSpeak uint32 `protobuf:"varint,3,opt,name=canSpeak" json:"canSpeak,omitempty"`
}

func (m *NotifyEnterGroup) Reset()                    { *m = NotifyEnterGroup{} }
func (m *NotifyEnterGroup) String() string            { return proto.CompactTextString(m) }
func (*NotifyEnterGroup) ProtoMessage()               {}
func (*NotifyEnterGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type NotifyLeaveGroup struct {
	GroupId  uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
}

func (m *NotifyLeaveGroup) Reset()                    { *m = NotifyLeaveGroup{} }
func (m *NotifyLeaveGroup) String() string            { return proto.CompactTextString(m) }
func (*NotifyLeaveGroup) ProtoMessage()               {}
func (*NotifyLeaveGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type LeaveGroup struct {
	GroupId  uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
}

func (m *LeaveGroup) Reset()                    { *m = LeaveGroup{} }
func (m *LeaveGroup) String() string            { return proto.CompactTextString(m) }
func (*LeaveGroup) ProtoMessage()               {}
func (*LeaveGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type AdjustPlayerAttr struct {
	GroupId  uint64 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	PlayerId uint64 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
	CanSpeak uint32 `protobuf:"varint,3,opt,name=canSpeak" json:"canSpeak,omitempty"`
}

func (m *AdjustPlayerAttr) Reset()                    { *m = AdjustPlayerAttr{} }
func (m *AdjustPlayerAttr) String() string            { return proto.CompactTextString(m) }
func (*AdjustPlayerAttr) ProtoMessage()               {}
func (*AdjustPlayerAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Request struct {
	Handshake        *Handshake        `protobuf:"bytes,1,opt,name=handshake" json:"handshake,omitempty"`
	EnterGroup       *EnterGroup       `protobuf:"bytes,2,opt,name=enterGroup" json:"enterGroup,omitempty"`
	LeaveGroup       *LeaveGroup       `protobuf:"bytes,3,opt,name=leaveGroup" json:"leaveGroup,omitempty"`
	AdjustPlayerAttr *AdjustPlayerAttr `protobuf:"bytes,4,opt,name=adjustPlayerAttr" json:"adjustPlayerAttr,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Request) GetHandshake() *Handshake {
	if m != nil {
		return m.Handshake
	}
	return nil
}

func (m *Request) GetEnterGroup() *EnterGroup {
	if m != nil {
		return m.EnterGroup
	}
	return nil
}

func (m *Request) GetLeaveGroup() *LeaveGroup {
	if m != nil {
		return m.LeaveGroup
	}
	return nil
}

func (m *Request) GetAdjustPlayerAttr() *AdjustPlayerAttr {
	if m != nil {
		return m.AdjustPlayerAttr
	}
	return nil
}

type Response struct {
	Code          ACK            `protobuf:"varint,1,opt,name=Code,json=code,enum=message.ACK" json:"Code,omitempty"`
	HandshakeAck  *HandshakeAck  `protobuf:"bytes,3,opt,name=handshakeAck" json:"handshakeAck,omitempty"`
	EnterGroupAck *EnterGroupAck `protobuf:"bytes,4,opt,name=enterGroupAck" json:"enterGroupAck,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Response) GetHandshakeAck() *HandshakeAck {
	if m != nil {
		return m.HandshakeAck
	}
	return nil
}

func (m *Response) GetEnterGroupAck() *EnterGroupAck {
	if m != nil {
		return m.EnterGroupAck
	}
	return nil
}

type Noitfy struct {
	NotifyEnterGroup *NotifyEnterGroup `protobuf:"bytes,1,opt,name=notifyEnterGroup" json:"notifyEnterGroup,omitempty"`
	NotifyLeaveGroup *NotifyLeaveGroup `protobuf:"bytes,2,opt,name=notifyLeaveGroup" json:"notifyLeaveGroup,omitempty"`
}

func (m *Noitfy) Reset()                    { *m = Noitfy{} }
func (m *Noitfy) String() string            { return proto.CompactTextString(m) }
func (*Noitfy) ProtoMessage()               {}
func (*Noitfy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Noitfy) GetNotifyEnterGroup() *NotifyEnterGroup {
	if m != nil {
		return m.NotifyEnterGroup
	}
	return nil
}

func (m *Noitfy) GetNotifyLeaveGroup() *NotifyLeaveGroup {
	if m != nil {
		return m.NotifyLeaveGroup
	}
	return nil
}

type ControlMessage struct {
	Req    *Request  `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
	Res    *Response `protobuf:"bytes,2,opt,name=res" json:"res,omitempty"`
	Notify *Noitfy   `protobuf:"bytes,3,opt,name=notify" json:"notify,omitempty"`
}

func (m *ControlMessage) Reset()                    { *m = ControlMessage{} }
func (m *ControlMessage) String() string            { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()               {}
func (*ControlMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ControlMessage) GetReq() *Request {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *ControlMessage) GetRes() *Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func (m *ControlMessage) GetNotify() *Noitfy {
	if m != nil {
		return m.Notify
	}
	return nil
}

type PlayerRequestConnect struct {
	PlayerIndex uint32 `protobuf:"varint,1,opt,name=PlayerIndex,json=playerIndex" json:"PlayerIndex,omitempty"`
	PlayerID    uint64 `protobuf:"varint,2,opt,name=PlayerID,json=playerID" json:"PlayerID,omitempty"`
	GroupID     uint64 `protobuf:"varint,3,opt,name=GroupID,json=groupID" json:"GroupID,omitempty"`
}

func (m *PlayerRequestConnect) Reset()                    { *m = PlayerRequestConnect{} }
func (m *PlayerRequestConnect) String() string            { return proto.CompactTextString(m) }
func (*PlayerRequestConnect) ProtoMessage()               {}
func (*PlayerRequestConnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type PlayerRequestConnectAck struct {
	Code VOICE_ACK `protobuf:"varint,1,opt,name=Code,json=code,enum=message.VOICE_ACK" json:"Code,omitempty"`
}

func (m *PlayerRequestConnectAck) Reset()                    { *m = PlayerRequestConnectAck{} }
func (m *PlayerRequestConnectAck) String() string            { return proto.CompactTextString(m) }
func (*PlayerRequestConnectAck) ProtoMessage()               {}
func (*PlayerRequestConnectAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type VoiceData struct {
	PlayerIndex uint32 `protobuf:"varint,1,opt,name=PlayerIndex,json=playerIndex" json:"PlayerIndex,omitempty"`
	PlayerID    uint64 `protobuf:"varint,2,opt,name=PlayerID,json=playerID" json:"PlayerID,omitempty"`
	FrameIndex  uint32 `protobuf:"varint,3,opt,name=FrameIndex,json=frameIndex" json:"FrameIndex,omitempty"`
	Data        []byte `protobuf:"bytes,4,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *VoiceData) Reset()                    { *m = VoiceData{} }
func (m *VoiceData) String() string            { return proto.CompactTextString(m) }
func (*VoiceData) ProtoMessage()               {}
func (*VoiceData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type VoiceHeartbeat struct {
	PlayerIndex uint32 `protobuf:"varint,1,opt,name=PlayerIndex,json=playerIndex" json:"PlayerIndex,omitempty"`
	PlayerID    uint64 `protobuf:"varint,2,opt,name=PlayerID,json=playerID" json:"PlayerID,omitempty"`
}

func (m *VoiceHeartbeat) Reset()                    { *m = VoiceHeartbeat{} }
func (m *VoiceHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*VoiceHeartbeat) ProtoMessage()               {}
func (*VoiceHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type VoiceMessage struct {
	PlayerRequestConnect    *PlayerRequestConnect    `protobuf:"bytes,1,opt,name=playerRequestConnect" json:"playerRequestConnect,omitempty"`
	PlayerRequestConnectAck *PlayerRequestConnectAck `protobuf:"bytes,2,opt,name=playerRequestConnectAck" json:"playerRequestConnectAck,omitempty"`
	VoiceData               *VoiceData               `protobuf:"bytes,3,opt,name=voiceData" json:"voiceData,omitempty"`
	VoiceHeartbeat          *VoiceHeartbeat          `protobuf:"bytes,4,opt,name=voiceHeartbeat" json:"voiceHeartbeat,omitempty"`
}

func (m *VoiceMessage) Reset()                    { *m = VoiceMessage{} }
func (m *VoiceMessage) String() string            { return proto.CompactTextString(m) }
func (*VoiceMessage) ProtoMessage()               {}
func (*VoiceMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *VoiceMessage) GetPlayerRequestConnect() *PlayerRequestConnect {
	if m != nil {
		return m.PlayerRequestConnect
	}
	return nil
}

func (m *VoiceMessage) GetPlayerRequestConnectAck() *PlayerRequestConnectAck {
	if m != nil {
		return m.PlayerRequestConnectAck
	}
	return nil
}

func (m *VoiceMessage) GetVoiceData() *VoiceData {
	if m != nil {
		return m.VoiceData
	}
	return nil
}

func (m *VoiceMessage) GetVoiceHeartbeat() *VoiceHeartbeat {
	if m != nil {
		return m.VoiceHeartbeat
	}
	return nil
}

func init() {
	proto.RegisterType((*Handshake)(nil), "message.Handshake")
	proto.RegisterType((*HandshakeAck)(nil), "message.HandshakeAck")
	proto.RegisterType((*EnterGroup)(nil), "message.EnterGroup")
	proto.RegisterType((*EnterGroupAck)(nil), "message.EnterGroupAck")
	proto.RegisterType((*NotifyEnterGroup)(nil), "message.NotifyEnterGroup")
	proto.RegisterType((*NotifyLeaveGroup)(nil), "message.NotifyLeaveGroup")
	proto.RegisterType((*LeaveGroup)(nil), "message.LeaveGroup")
	proto.RegisterType((*AdjustPlayerAttr)(nil), "message.AdjustPlayerAttr")
	proto.RegisterType((*Request)(nil), "message.Request")
	proto.RegisterType((*Response)(nil), "message.Response")
	proto.RegisterType((*Noitfy)(nil), "message.Noitfy")
	proto.RegisterType((*ControlMessage)(nil), "message.ControlMessage")
	proto.RegisterType((*PlayerRequestConnect)(nil), "message.PlayerRequestConnect")
	proto.RegisterType((*PlayerRequestConnectAck)(nil), "message.PlayerRequestConnectAck")
	proto.RegisterType((*VoiceData)(nil), "message.VoiceData")
	proto.RegisterType((*VoiceHeartbeat)(nil), "message.VoiceHeartbeat")
	proto.RegisterType((*VoiceMessage)(nil), "message.VoiceMessage")
	proto.RegisterEnum("message.MSG_CONST", MSG_CONST_name, MSG_CONST_value)
	proto.RegisterEnum("message.MSG", MSG_name, MSG_value)
	proto.RegisterEnum("message.ACK", ACK_name, ACK_value)
	proto.RegisterEnum("message.VOICE_MSG", VOICE_MSG_name, VOICE_MSG_value)
	proto.RegisterEnum("message.VOICE_ACK", VOICE_ACK_name, VOICE_ACK_value)
}

var fileDescriptor0 = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x89, 0x49, 0x93, 0xd3, 0x24, 0x75, 0xa7, 0xed, 0x36, 0xbb, 0x12, 0xa8, 0x32, 0x12,
	0xa0, 0x20, 0xad, 0x50, 0xf7, 0x0a, 0x84, 0x84, 0x66, 0xe3, 0x69, 0x13, 0x9a, 0x3a, 0xd9, 0x89,
	0x5b, 0xb1, 0x7b, 0x41, 0xe4, 0x4d, 0xa6, 0x3f, 0xb4, 0x6b, 0x67, 0x6d, 0xb7, 0xa2, 0x17, 0x48,
	0x88, 0x87, 0xe0, 0x0d, 0xb8, 0xe7, 0x2d, 0x78, 0x19, 0x1e, 0x82, 0x99, 0xf1, 0xd8, 0x9e, 0xfc,
	0xc0, 0xc5, 0x46, 0xbd, 0xf3, 0xf9, 0x3f, 0xe7, 0x3b, 0xdf, 0x99, 0x04, 0x1a, 0xef, 0x58, 0x1c,
	0xfb, 0x97, 0xec, 0xf9, 0x2c, 0x0a, 0x93, 0x10, 0x6d, 0x28, 0xd1, 0xfe, 0x0a, 0x6a, 0x5d, 0x3f,
	0x98, 0xc6, 0x57, 0xfe, 0x0d, 0x43, 0x9f, 0x02, 0x0c, 0xa3, 0x70, 0x7a, 0x37, 0x49, 0x4e, 0xd8,
	0x43, 0xab, 0x74, 0x60, 0x7c, 0x59, 0xa3, 0x30, 0xcb, 0x35, 0x76, 0x13, 0xea, 0xb9, 0x33, 0x9e,
	0xdc, 0xd8, 0x3f, 0x01, 0x90, 0x20, 0x61, 0xd1, 0x71, 0x14, 0xde, 0xcd, 0x50, 0x0b, 0x36, 0x2e,
	0xc5, 0x47, 0x6f, 0xda, 0x32, 0x78, 0xa8, 0x49, 0x33, 0x11, 0x3d, 0x83, 0xea, 0xec, 0xd6, 0x7f,
	0x60, 0x11, 0x37, 0x95, 0xa4, 0x29, 0x97, 0x85, 0x6d, 0xe2, 0x07, 0xa3, 0x19, 0xf3, 0x6f, 0x5a,
	0x65, 0x6e, 0x6b, 0xd0, 0x5c, 0xb6, 0x2f, 0xa1, 0x51, 0xe4, 0xe7, 0x05, 0x3f, 0xb0, 0xc4, 0x01,
	0x6c, 0xaa, 0xef, 0x60, 0xca, 0x7e, 0x51, 0x55, 0x74, 0x95, 0x3d, 0x05, 0xcb, 0x0d, 0x93, 0xeb,
	0x8b, 0x87, 0x47, 0x1d, 0xa7, 0x9b, 0x55, 0xe9, 0x33, 0xff, 0x9e, 0xad, 0x51, 0xc5, 0x7e, 0x09,
	0xb0, 0x76, 0x0e, 0x3e, 0x33, 0x9e, 0xfe, 0x7c, 0x17, 0x27, 0x43, 0xa9, 0xc1, 0x49, 0x12, 0x3d,
	0xc2, 0xcc, 0xff, 0x18, 0xb0, 0x41, 0xd9, 0xfb, 0x3b, 0x16, 0x27, 0xe8, 0x6b, 0xa8, 0x5d, 0x65,
	0xf4, 0x91, 0xf9, 0x37, 0x0f, 0xd1, 0xf3, 0x8c, 0x97, 0x39, 0xb1, 0x68, 0xe1, 0x84, 0x5e, 0x00,
	0xb0, 0x7c, 0x23, 0xb2, 0xee, 0xe6, 0xe1, 0x4e, 0x1e, 0x52, 0x2c, 0x8b, 0x6a, 0x6e, 0x22, 0xe8,
	0x36, 0x07, 0x47, 0x36, 0xa4, 0x07, 0x15, 0xb8, 0x51, 0xcd, 0x0d, 0x11, 0xb0, 0xfc, 0x05, 0x34,
	0x5a, 0xa6, 0x0c, 0x7d, 0x9a, 0x87, 0x2e, 0xc2, 0x45, 0x97, 0x42, 0xec, 0x3f, 0x0d, 0xa8, 0x52,
	0x16, 0xcf, 0xc2, 0x20, 0x66, 0x9c, 0x77, 0x66, 0x27, 0x9c, 0xa6, 0xa3, 0x36, 0x0f, 0xeb, 0x45,
	0x9e, 0xce, 0x09, 0x35, 0x27, 0xdc, 0x82, 0xbe, 0x81, 0xfa, 0x95, 0x76, 0x50, 0xaa, 0xd9, 0xbd,
	0x65, 0x50, 0xb8, 0x91, 0xce, 0xb9, 0xa2, 0xef, 0xa0, 0xc1, 0xf4, 0xdb, 0x50, 0xdd, 0x3e, 0x59,
	0x81, 0x8e, 0x08, 0x9e, 0x77, 0xb6, 0xff, 0x30, 0xa0, 0xe2, 0x86, 0xd7, 0xc9, 0xc5, 0x83, 0x98,
	0x3c, 0x58, 0xe0, 0xbe, 0x5a, 0x4e, 0x31, 0xf9, 0xe2, 0x71, 0xd0, 0xa5, 0x90, 0x22, 0x4d, 0x01,
	0xb0, 0x5a, 0xd8, 0x62, 0x1a, 0x6d, 0x03, 0x4b, 0x21, 0xf6, 0xef, 0x06, 0x34, 0x3b, 0x61, 0x90,
	0x44, 0xe1, 0xed, 0x69, 0x1a, 0x85, 0x6c, 0x28, 0x47, 0xec, 0xbd, 0xea, 0xc9, 0xca, 0x93, 0x29,
	0x56, 0x51, 0x61, 0x44, 0x9f, 0x09, 0x9f, 0x58, 0x15, 0xdc, 0xd6, 0x7c, 0xd2, 0x55, 0x08, 0xa7,
	0x18, 0x7d, 0x01, 0x95, 0xb4, 0x9e, 0xc2, 0x79, 0x4b, 0x6b, 0x4c, 0x40, 0x41, 0x95, 0xd9, 0x0e,
	0x60, 0x37, 0xdd, 0xa9, 0xaa, 0xc1, 0x1b, 0x0a, 0xd8, 0x24, 0x11, 0x0f, 0xc9, 0x50, 0x7b, 0x48,
	0x8c, 0xa5, 0x87, 0x44, 0x9c, 0x82, 0xf2, 0x70, 0x16, 0xce, 0xc4, 0x11, 0xc7, 0x25, 0x67, 0xe4,
	0xa6, 0xb2, 0x7e, 0x5c, 0x8e, 0x8d, 0x61, 0x7f, 0x55, 0x3d, 0xb1, 0xe6, 0xcf, 0xe7, 0x38, 0x54,
	0x9c, 0xcb, 0xf9, 0xa0, 0xd7, 0x21, 0xe3, 0x9c, 0x49, 0xf6, 0xaf, 0x50, 0x3b, 0x0f, 0xaf, 0x27,
	0xcc, 0xf1, 0x13, 0x7f, 0xcd, 0x3e, 0xf9, 0xaf, 0xc0, 0x51, 0xe4, 0xbf, 0x63, 0xfa, 0x6b, 0x09,
	0x17, 0xb9, 0x06, 0x21, 0x30, 0x45, 0x15, 0x49, 0xb8, 0x3a, 0x35, 0xa7, 0xfc, 0xdb, 0x76, 0xa1,
	0x29, 0xcb, 0x77, 0x99, 0x1f, 0x25, 0x6f, 0x99, 0xbf, 0x26, 0x56, 0xf6, 0x5f, 0x25, 0xa8, 0xcb,
	0x84, 0x19, 0x09, 0x5e, 0xc1, 0xee, 0x6c, 0x05, 0x44, 0x8a, 0x15, 0x9f, 0xe4, 0xb8, 0xac, 0xc2,
	0x91, 0xae, 0x0c, 0x45, 0x6f, 0x60, 0x7f, 0xb6, 0x1a, 0x75, 0xc5, 0xa3, 0x83, 0xff, 0xcd, 0x2a,
	0xae, 0xea, 0xbf, 0x12, 0x88, 0xa7, 0xee, 0x3e, 0x5b, 0x87, 0x62, 0x9b, 0xb6, 0xbb, 0xcc, 0x42,
	0x0b, 0x27, 0xf4, 0x3d, 0x34, 0xef, 0xe7, 0x10, 0x54, 0x07, 0xbd, 0x3f, 0x1f, 0x96, 0x9b, 0xe9,
	0x82, 0x7b, 0x9b, 0x42, 0xed, 0x74, 0x74, 0x3c, 0xee, 0x0c, 0xdc, 0x91, 0x87, 0xb6, 0x60, 0x53,
	0x7e, 0x8c, 0xcf, 0x71, 0xff, 0x8c, 0x58, 0x1f, 0xa1, 0x5d, 0xd8, 0x3a, 0xc5, 0x3f, 0x8e, 0x87,
	0x9c, 0x31, 0xc4, 0x1b, 0x8f, 0x7a, 0x6f, 0x88, 0xf5, 0x5b, 0x95, 0xaf, 0x60, 0x4f, 0x68, 0x53,
	0x32, 0xcd, 0xdb, 0x9e, 0x95, 0x2c, 0xa3, 0xfd, 0xb7, 0x01, 0x65, 0x9e, 0x14, 0x55, 0xc1, 0x74,
	0x07, 0xae, 0xc8, 0xd3, 0xe0, 0xff, 0x17, 0xb0, 0xeb, 0x8c, 0xba, 0xf8, 0x84, 0x58, 0x06, 0xda,
	0x86, 0x46, 0x2e, 0x0a, 0x36, 0x5a, 0x25, 0x51, 0x9a, 0xb8, 0x1e, 0xa1, 0xe3, 0x63, 0x3a, 0x38,
	0x1b, 0x5a, 0x65, 0xb4, 0x03, 0x5b, 0x9a, 0x42, 0x7a, 0x99, 0xe8, 0x09, 0x20, 0x77, 0xe0, 0xf5,
	0x8e, 0x5e, 0x8f, 0x75, 0xe7, 0x8f, 0x35, 0x7d, 0x9f, 0xe0, 0x73, 0xa2, 0xf4, 0x15, 0x91, 0x55,
	0x57, 0x6c, 0x08, 0x47, 0xec, 0xfc, 0x70, 0xc6, 0x47, 0x1c, 0xf6, 0xf1, 0x6b, 0x9e, 0x01, 0x7b,
	0x1e, 0xb5, 0xaa, 0x68, 0x0f, 0xb6, 0x95, 0x5e, 0x95, 0x13, 0xea, 0x5a, 0xfb, 0x5b, 0x28, 0xf3,
	0xc2, 0xa8, 0x02, 0xa5, 0xc1, 0x89, 0x84, 0xc3, 0x2a, 0xfa, 0x3e, 0xc2, 0xbd, 0x3e, 0x71, 0xf8,
	0x34, 0x32, 0xd6, 0xc9, 0x12, 0x2a, 0x75, 0xa9, 0x1d, 0xf1, 0xdb, 0x92, 0x08, 0x09, 0x28, 0x10,
	0x67, 0x7a, 0x26, 0x8c, 0x15, 0x28, 0x7c, 0x42, 0x4a, 0x5e, 0x11, 0x51, 0x94, 0xa3, 0xee, 0x92,
	0x8e, 0xc7, 0x93, 0xed, 0xc3, 0xce, 0x82, 0x52, 0x01, 0xd4, 0x04, 0x48, 0x33, 0x38, 0xd8, 0xc3,
	0x29, 0x3e, 0xa9, 0xdc, 0x25, 0x98, 0x7a, 0x2f, 0x09, 0xf6, 0x2c, 0xb3, 0xfd, 0x34, 0xab, 0x29,
	0xba, 0xae, 0x43, 0x35, 0x15, 0x44, 0xef, 0x6f, 0x2b, 0xf2, 0x2f, 0xdc, 0x8b, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xdc, 0xd5, 0x34, 0x1e, 0xd3, 0x09, 0x00, 0x00,
}