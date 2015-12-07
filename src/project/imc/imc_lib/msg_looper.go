//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: msg_looper.go
//     description:
//         created: 2015-12-05 22:25:20
//          author: wystan
//
//===============================================================================

package imc_lib

import (
	"bufio"
	"log"
	"net"
)

type msgHandler interface {
	/* methods */
	HandleMessage(data []byte) error
	Exit()
}

type msgLooper struct {
	handler msgHandler
}

func (l *msgLooper) Loop(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		data, err := readMsgData(reader)
		if err != nil {
			log.Printf("client exit! %v", err)
			break
		}
		if err = l.handler.HandleMessage(data); err != nil {
			log.Println(err)
		}
	}
	l.handler.Exit()
}

func newMsgLooper(h msgHandler) *msgLooper {
	return &msgLooper{handler: h}
}

//==================================== END ======================================
