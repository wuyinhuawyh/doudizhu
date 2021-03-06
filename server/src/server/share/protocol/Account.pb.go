// Code generated by protoc-gen-go.
// source: Account.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	Account.proto

It has these top-level messages:
	Account
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Account struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}

// 客户端消息
// pid = 1
type Account_RegisterPlayer struct {
	Playername       *string `protobuf:"bytes,1,opt,name=playername" json:"playername,omitempty"`
	Passworld        *string `protobuf:"bytes,2,opt,name=passworld" json:"passworld,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Account_RegisterPlayer) Reset()         { *m = Account_RegisterPlayer{} }
func (m *Account_RegisterPlayer) String() string { return proto.CompactTextString(m) }
func (*Account_RegisterPlayer) ProtoMessage()    {}

func (m *Account_RegisterPlayer) GetPlayername() string {
	if m != nil && m.Playername != nil {
		return *m.Playername
	}
	return ""
}

func (m *Account_RegisterPlayer) GetPassworld() string {
	if m != nil && m.Passworld != nil {
		return *m.Passworld
	}
	return ""
}

// pid = 1
type Account_RegisterResult struct {
	Result           *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Account_RegisterResult) Reset()         { *m = Account_RegisterResult{} }
func (m *Account_RegisterResult) String() string { return proto.CompactTextString(m) }
func (*Account_RegisterResult) ProtoMessage()    {}

func (m *Account_RegisterResult) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// pid = 2
type Account_LoginInfo struct {
	Playername       *string `protobuf:"bytes,1,opt,name=playername" json:"playername,omitempty"`
	Passworld        *string `protobuf:"bytes,2,opt,name=passworld" json:"passworld,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Account_LoginInfo) Reset()         { *m = Account_LoginInfo{} }
func (m *Account_LoginInfo) String() string { return proto.CompactTextString(m) }
func (*Account_LoginInfo) ProtoMessage()    {}

func (m *Account_LoginInfo) GetPlayername() string {
	if m != nil && m.Playername != nil {
		return *m.Playername
	}
	return ""
}

func (m *Account_LoginInfo) GetPassworld() string {
	if m != nil && m.Passworld != nil {
		return *m.Passworld
	}
	return ""
}

// pid = 2
type Account_LoginResult struct {
	Result           *int32  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Gameserver       *string `protobuf:"bytes,2,opt,name=gameserver" json:"gameserver,omitempty"`
	PlayerId         *int32  `protobuf:"varint,3,opt,name=player_id" json:"player_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Account_LoginResult) Reset()         { *m = Account_LoginResult{} }
func (m *Account_LoginResult) String() string { return proto.CompactTextString(m) }
func (*Account_LoginResult) ProtoMessage()    {}

func (m *Account_LoginResult) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *Account_LoginResult) GetGameserver() string {
	if m != nil && m.Gameserver != nil {
		return *m.Gameserver
	}
	return ""
}

func (m *Account_LoginResult) GetPlayerId() int32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

// pid = 101 返回游戏服在线人数
type Account_GameResult struct {
	Count            *int32  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	GameAddress      *string `protobuf:"bytes,2,opt,name=game_address" json:"game_address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Account_GameResult) Reset()         { *m = Account_GameResult{} }
func (m *Account_GameResult) String() string { return proto.CompactTextString(m) }
func (*Account_GameResult) ProtoMessage()    {}

func (m *Account_GameResult) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *Account_GameResult) GetGameAddress() string {
	if m != nil && m.GameAddress != nil {
		return *m.GameAddress
	}
	return ""
}

// pid = 102
type Account_NoteGame struct {
	PlayerId         *int32 `protobuf:"varint,1,opt,name=player_id" json:"player_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Account_NoteGame) Reset()         { *m = Account_NoteGame{} }
func (m *Account_NoteGame) String() string { return proto.CompactTextString(m) }
func (*Account_NoteGame) ProtoMessage()    {}

func (m *Account_NoteGame) GetPlayerId() int32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}
