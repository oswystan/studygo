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
	"time"

	"github.com/golang/protobuf/proto"
)

func doLogin(writer *bufio.Writer, name, pwd string) error {
	login := &imc.ImcCmd{
		CmdType: proto.Int32(1),
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

func doModifyInfo(writer *bufio.Writer, name, pwd, nick string) error {
	modify := &imc.ImcCmd{
		CmdType: proto.Int32(2),
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
	doLogin(writer, "wangyu", "12")
	time.Sleep(10 * time.Millisecond)
	doModifyInfo(writer, "wangyu", "111", "wystan")
	time.Sleep(10 * time.Millisecond)
}

//==================================== END ======================================
