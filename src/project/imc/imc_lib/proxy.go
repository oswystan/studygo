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

	data, err := proto.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}
	p.writer.Write(data)
	p.writer.Flush()

	ack := p.waitForResponse(p.reader)
	if *ack.AckCommon.Status != RET_CODE_SUCCESS {
		return fmt.Errorf("%s", *ack.AckCommon.ErrorDesc)
	}
	return nil
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

	data, err := proto.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}
	p.writer.Write(data)
	p.writer.Flush()

	ack := p.waitForResponse(p.reader)
	if *ack.AckCommon.Status != RET_CODE_SUCCESS {
		return fmt.Errorf("%s", *ack.AckCommon.ErrorDesc)
	}
	return nil
}
func (p *ImcProxy) DoLogout(name string) error {
	cmd := CMD_TYPE_LOGOUT
	c := &ImcCmd{
		CmdType: &cmd,
		Logout: &CmdLogout{
			UserName: proto.String(name),
		},
	}

	data, err := proto.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}
	p.writer.Write(data)
	p.writer.Flush()

	ack := p.waitForResponse(p.reader)
	if *ack.AckCommon.Status != RET_CODE_SUCCESS {
		return fmt.Errorf("%s", *ack.AckCommon.ErrorDesc)
	}
	return nil
}
func (p *ImcProxy) waitForResponse(reader *bufio.Reader) *ImcCmd {
	data := make([]byte, 1024)
	p.reader.Read(data)
	cmd := &ImcCmd{}
	proto.Unmarshal(data, cmd)
	return cmd
}

func NewImcProxy(r *bufio.Reader, w *bufio.Writer) *ImcProxy {
	p := &ImcProxy{
		reader: r,
		writer: w,
	}

	return p
}

//==================================== END ======================================
