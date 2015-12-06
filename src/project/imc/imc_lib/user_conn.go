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
	dispatcher map[CMD_TYPE]func(msg *ImcCmd) error
}

func (u *userConn) doLogin(msg *ImcCmd) error {
	log.Println("call doLogin")
	return nil
}

func (u *userConn) doModifyInfo(msg *ImcCmd) error {
	log.Println("call doModifyInfo")
	return nil
}

func (u *userConn) dispatchMsg(msg *ImcCmd) error {

	if fn, ok := u.dispatcher[*msg.CmdType]; ok {
		return fn(msg)
	} else {
		return fmt.Errorf("unknown command type <%s>", msg.CmdType.String())
	}
}

func (u *userConn) HandleMessage(data []byte, w *bufio.Writer) error {
	msg := &ImcCmd{}
	proto.Unmarshal(data, msg)
	return u.dispatchMsg(msg)
}

func (u *userConn) Init() {
	u.dispatcher = make(map[CMD_TYPE]func(msg *ImcCmd) error)
	u.dispatcher[CMD_TYPE_LOGIN] = u.doLogin
	u.dispatcher[CMD_TYPE_MODIFYINFO] = u.doModifyInfo
}

func newUserConn() *userConn {
	uc := &userConn{}
	uc.Init()
	return uc
}

//==================================== END ======================================
