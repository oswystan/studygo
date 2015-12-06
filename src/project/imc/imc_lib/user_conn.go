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
	"log"

	"github.com/golang/protobuf/proto"
)

type userConn struct {
}

func (u *userConn) HandleMessage(data []byte, w *bufio.Writer) error {
	msg := &ImcCmd{}
	proto.Unmarshal(data, msg)
	log.Println(msg)
	return nil
}

func newUserConn() *userConn {
	return &userConn{}
}

//==================================== END ======================================
