//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-05 20:55:10
//          author: wystan
//
//===============================================================================

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"study/protobuf/example"
)

func main() {

	modify := &example.Test{
		CmdType: proto.Int32(1),
		ModifyInfo: &example.CmdModifyInfo{
			UserName: proto.String("wangyu"),
			NewPwd:   proto.String("1234"),
			NickName: proto.String("wystan"),
		},
	}

	login := &example.Test{
		CmdType: proto.Int32(1),
		Login: &example.CmdLogin{
			UserName: proto.String("wangyu"),
			Passwd:   proto.String("1232323"),
		},
	}
	dataModify, err := proto.Marshal(modify)
	if err != nil {
		fmt.Printf("err\n")
		return
	}
	dataLogin, err := proto.Marshal(login)
	if err != nil {
		fmt.Printf("err\n")
		return
	}

	newTest := &example.Test{}
	err = proto.Unmarshal(dataModify, newTest)
	if err != nil {
		fmt.Printf("err modify\n")
		return
	}
	fmt.Printf("msg=%v\n", newTest)

	err = proto.Unmarshal(dataLogin, newTest)
	if err != nil {
		fmt.Printf("err login\n")
		return
	}
	fmt.Printf("msg=%v\n", newTest)
}

//==================================== END ======================================
