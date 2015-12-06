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
	"io"
	"log"
	"net"
)

type msgHandler interface {
	/* methods */
	HandleMessage(data []byte, w *bufio.Writer) error
}

type msgLooper struct {
	handler msgHandler
}

func (l *msgLooper) Loop(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)
	data := make([]byte, 1024)
	for {
		count, err := reader.Read(data)
		if err == io.EOF {
			log.Printf("client exit!")
			break
		}
		if err = l.handler.HandleMessage(data[0:count], writer); err != nil {
			break
		}
	}
}

func newMsgLooper(h msgHandler) *msgLooper {
	return &msgLooper{handler: h}
}

//==================================== END ======================================
