//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: looper.go
//     description:
//         created: 2015-12-05 22:25:20
//          author: wystan
//
//===============================================================================

package imservice

import (
	"bufio"
	"log"
	"net"
	"sync"
)

type msgHandler interface {
	/* methods */
	HandleMessage(data []byte) error
	Exit()
}

type msgLooper struct {
	running int
	handler msgHandler
	conn    net.Conn
	wg      sync.WaitGroup
}

func (l *msgLooper) Loop() {
	l.wg.Add(1)
	l.running = 1
	defer l.conn.Close()

	reader := bufio.NewReader(l.conn)
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
	l.running = 0
	l.wg.Done()
}

func (l *msgLooper) Exit() {
	l.conn.Close()
	l.handler.Exit()
	if l.running == 1 {
		l.wg.Wait()
	}
}

func newMsgLooper(h msgHandler, c net.Conn) *msgLooper {
	return &msgLooper{handler: h, conn: c}
}

//==================================== END ======================================
