//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: user_conn.go
//     description:
//         created: 2015-12-05 22:53:42
//          author: wystan
//
//===============================================================================

package imc_lib

import (
	"bufio"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

const (
	STATE_UNAUTHORIZED = iota
	STATE_AUTHORIZED
	STATE_MAX
)

type dispatcherType map[CMD_TYPE]func(msg *ImcCmd) (data []byte)

type userConn struct {
	dispatcher [STATE_MAX]dispatcherType
	state      int
}

func (u *userConn) doLogin(msg *ImcCmd) (data []byte) {
	log.Printf("LOGIN (user:%s,password:%s)", *msg.Login.UserName, *msg.Login.Passwd)

	cmd := CMD_TYPE_LOGIN_ACK
	status := RET_CODE_SUCCESS
	ack := &ImcCmd{
		CmdType: &cmd,
		AckCommon: &CmdAckCommon{
			Status:    &status,
			ErrorDesc: proto.String(""),
		},
	}
	data, _ = proto.Marshal(ack)

	// chage to state so the client can do
	// anything they want
	u.state = STATE_AUTHORIZED
	return data
}

func (u *userConn) doModifyInfo(msg *ImcCmd) (data []byte) {

	log.Printf("MODIFYINFO (user:%s, passwd:%s, nickname:%s)",
		*msg.ModifyInfo.UserName,
		*msg.ModifyInfo.NewPasswd,
		*msg.ModifyInfo.NickName)

	cmd := CMD_TYPE_MODIFYINFO_ACK
	status := RET_CODE_SUCCESS
	ack := &ImcCmd{
		CmdType: &cmd,
		AckCommon: &CmdAckCommon{
			Status:    &status,
			ErrorDesc: proto.String(""),
		},
	}
	data, _ = proto.Marshal(ack)
	return data
}
func (u *userConn) doLogout(msg *ImcCmd) (data []byte) {

	log.Printf("LOGOUT (user:%s)", *msg.Logout.UserName)

	cmd := CMD_TYPE_LOGOUT_ACK
	status := RET_CODE_SUCCESS
	ack := &ImcCmd{
		CmdType: &cmd,
		AckCommon: &CmdAckCommon{
			Status:    &status,
			ErrorDesc: proto.String(""),
		},
	}
	data, _ = proto.Marshal(ack)
	u.state = STATE_UNAUTHORIZED
	return data
}
func (u *userConn) dispatchMsg(msg *ImcCmd) (data []byte) {

	dispatcher := u.dispatcher[u.state]
	if fn, ok := dispatcher[*msg.CmdType]; ok {
		return fn(msg)
	} else {
		cmd := CMD_TYPE_MODIFYINFO_ACK
		status := RET_CODE_FAILED
		ack := &ImcCmd{
			CmdType: &cmd,
			AckCommon: &CmdAckCommon{
				Status:    &status,
				ErrorDesc: proto.String(fmt.Sprintf("unauthorized command type <%s>", msg.CmdType.String())),
			},
		}
		data, _ = proto.Marshal(ack)
		return data
	}
}

// Unmarshal the bytes to a ImcCmd and dispatch
// the msg to the corresponding handler
func (u *userConn) HandleMessage(data []byte, w *bufio.Writer) error {
	msg := &ImcCmd{}
	proto.Unmarshal(data, msg)
	ack := u.dispatchMsg(msg)
	w.Write(ack)
	w.Flush()
	return nil
}

// init the dispatcher of current userConn
func (u *userConn) Init() {
	u.state = STATE_UNAUTHORIZED

	for i := 0; i < STATE_MAX; i++ {
		u.dispatcher[i] = make(dispatcherType)
	}

	unauthorized := u.dispatcher[STATE_UNAUTHORIZED]
	unauthorized[CMD_TYPE_LOGIN] = u.doLogin

	authorized := u.dispatcher[STATE_AUTHORIZED]
	authorized[CMD_TYPE_LOGIN] = u.doLogin
	authorized[CMD_TYPE_MODIFYINFO] = u.doModifyInfo
	authorized[CMD_TYPE_LOGOUT] = u.doLogout
}

func newUserConn() *userConn {
	uc := &userConn{}
	uc.Init()
	return uc
}

//==================================== END ======================================
