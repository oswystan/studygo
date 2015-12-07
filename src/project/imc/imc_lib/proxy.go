//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: proxy.go
//     description:
//         created: 2015-12-06 19:28:22
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

type ImcProxy struct {
	writer *bufio.Writer
	reader *bufio.Reader
}

func (p *ImcProxy) DoLogin(name, pwd string) error {
	cmd := CMD_TYPE_LOGIN
	c := &ImcCmd{
		CmdType: &cmd,
		Login: &CmdLogin{
			UserName: proto.String(name),
			Passwd:   proto.String(pwd),
		},
	}

	return p.sendAndWait(c)
}

func (p *ImcProxy) DoModifyInfo(name, pwd, nick string) error {
	cmd := CMD_TYPE_MODIFYINFO
	c := &ImcCmd{
		CmdType: &cmd,
		ModifyInfo: &CmdModifyInfo{
			UserName:  proto.String(name),
			NewPasswd: proto.String(pwd),
			NickName:  proto.String(nick),
		},
	}

	return p.sendAndWait(c)
}
func (p *ImcProxy) DoLogout(name string) error {
	cmd := CMD_TYPE_LOGOUT
	c := &ImcCmd{
		CmdType: &cmd,
		Logout: &CmdLogout{
			UserName: proto.String(name),
		},
	}

	return p.sendAndWait(c)
}
func (p *ImcProxy) DoSendMsg(peer, msg string) error {
	cmd := CMD_TYPE_SENDMSG
	c := &ImcCmd{
		CmdType: &cmd,
		SendMsg: &CmdSendMsg{
			PeerName: proto.String(peer),
			MsgBody:  proto.String(msg),
		},
	}

	return p.sendAndWait(c)
}
func (p *ImcProxy) waitForResponse(reader *bufio.Reader) *ImcCmd {
	cmd := &ImcCmd{}

	data, err := readMsgData(reader)
	if err != nil {
		log.Println(err)
		return nil
	}

	proto.Unmarshal(data, cmd)
	return cmd
}

func (p *ImcProxy) sendAndWait(c *ImcCmd) error {
	data, err := proto.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}

	writeMsgData(p.writer, data)

	ack := p.waitForResponse(p.reader)
	if *ack.AckCommon.Status != RET_CODE_SUCCESS {
		return fmt.Errorf("%s", *ack.AckCommon.ErrorDesc)
	}
	return nil
}

func NewImcProxy(r *bufio.Reader, w *bufio.Writer) *ImcProxy {
	p := &ImcProxy{
		reader: r,
		writer: w,
	}

	return p
}

//==================================== END ======================================
