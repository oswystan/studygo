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

type userConn struct {
	dispatcher map[CMD_TYPE]func(msg *ImcCmd) (data []byte)
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

func (u *userConn) dispatchMsg(msg *ImcCmd) (data []byte) {

	if fn, ok := u.dispatcher[*msg.CmdType]; ok {
		return fn(msg)
	} else {
		cmd := CMD_TYPE_MODIFYINFO_ACK
		status := RET_CODE_FAILED
		ack := &ImcCmd{
			CmdType: &cmd,
			AckCommon: &CmdAckCommon{
				Status:    &status,
				ErrorDesc: proto.String(fmt.Sprintf("unknow command type <%s>", msg.CmdType.String())),
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
	u.dispatcher = make(map[CMD_TYPE]func(msg *ImcCmd) []byte)
	u.dispatcher[CMD_TYPE_LOGIN] = u.doLogin
	u.dispatcher[CMD_TYPE_MODIFYINFO] = u.doModifyInfo
}

func newUserConn() *userConn {
	uc := &userConn{}
	uc.Init()
	return uc
}

//==================================== END ======================================
