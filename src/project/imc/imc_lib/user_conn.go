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
	"strings"
)

type userConn struct {
}

func (u *userConn) HandleMessage(data []byte, w *bufio.Writer) error {
	str := strings.TrimSuffix(string(data), "\n")
	log.Printf("%s\n", str)
	w.WriteString("<<<" + str + "\n")
	w.Flush()
	return nil
}

func newUserConn() *userConn {
	return &userConn{}
}

//==================================== END ======================================
