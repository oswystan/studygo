//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: imc.go
//     description:
//         created: 2015-12-06 09:19:42
//          author: wystan
//
//===============================================================================

package main

import (
	"bufio"
	"log"
	"net"
	imc "project/imc/imc_lib"

	"github.com/golang/protobuf/proto"
)

func doLogin(writer *bufio.Writer, name, pwd string) error {
	cmd := imc.CMD_TYPE_LOGIN
	login := &imc.ImcCmd{
		CmdType: &cmd,
		Login: &imc.CmdLogin{
			UserName: proto.String(name),
			Passwd:   proto.String(pwd),
		},
	}

	data, err := proto.Marshal(login)
	if err != nil {
		log.Println(err)
		return err
	}
	writer.Write(data)
	writer.Flush()
	return nil
}

func waitForResponse(reader *bufio.Reader) *imc.ImcCmd {
	data := make([]byte, 1024)
	reader.Read(data)
	cmd := &imc.ImcCmd{}
	proto.Unmarshal(data, cmd)
	return cmd
}

func doModifyInfo(writer *bufio.Writer, name, pwd, nick string) error {
	cmd := imc.CMD_TYPE_MODIFYINFO
	modify := &imc.ImcCmd{
		CmdType: &cmd,
		ModifyInfo: &imc.CmdModifyInfo{
			UserName:  proto.String(name),
			NewPasswd: proto.String(pwd),
			NickName:  proto.String(nick),
		},
	}

	data, err := proto.Marshal(modify)
	if err != nil {
		log.Println(err)
		return err
	}
	writer.Write(data)
	writer.Flush()

	return nil
}

func main() {
	log.Println("client started ...")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	doLogin(writer, "wangyu", "12")
	ack := waitForResponse(reader)
	if *ack.AckCommon.Status != imc.RET_CODE_SUCCESS {
		log.Println("fail to do login", *ack.AckCommon.ErrorDesc)
	}

	doModifyInfo(writer, "wangyu", "111", "wystan")
	ack = waitForResponse(reader)
	if *ack.AckCommon.Status != imc.RET_CODE_SUCCESS {
		log.Println("fail to do modifyinfo", *ack.AckCommon.ErrorDesc)
	}
}

//==================================== END ======================================
