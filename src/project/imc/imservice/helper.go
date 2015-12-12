//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: helper.go
//     description:
//         created: 2015-12-09 16:10:50
//          author: wystan
//
//===============================================================================

package imservice

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Helper struct {
	proxy     *ImcProxy
	receiver  *receiver
	messageCh chan *Message
}
type Message struct {
	From     string
	To       string
	Body     string
	Datetime time.Time
}

type receiver struct {
	/* data */
	ch chan<- *Message
}

func (r *receiver) Receive(msg *ImcCmd) int {
	if *msg.CmdType != CMD_TYPE_MESSAGE {
		return MSG_CONTINUE
	}
	m := &Message{}
	m.From = *msg.Message.From
	m.To = *msg.Message.To
	m.Body = *msg.Message.MsgBody
	m.Datetime = time.Unix(*msg.Message.Datetime, 0)

	r.ch <- m

	return MSG_STOP
}

func (h *Helper) Connect(ip, port string) error {
	if h.proxy != nil {
		return fmt.Errorf("Close it first")
	}

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return err
	}

	h.proxy = NewImcProxy(conn)
	h.proxy.RegisterReceiver(h.receiver)
	err = h.proxy.Start()
	return err
}

func (h *Helper) RunCommand(cmd string) error {
	para := strings.Split(cmd, " ")
	switch para[0] {
	case "login":
		if len(para) != 3 {
			return fmt.Errorf("please check the parameters %s", cmd)
		}
		return h.proxy.DoLogin(para[1], para[2])

	case "talk":
		if len(para) < 3 {
			return fmt.Errorf("please check the parameters %s", cmd)
		}

		msg := strings.Join(para[2:], " ")
		return h.proxy.DoSendMsg(para[1], msg)

	case "modify":
		if len(para) != 4 {
			return fmt.Errorf("please check the parameters %s", cmd)
		}
		return h.proxy.DoModifyInfo(para[1], para[2], para[3])

	case "logout":
		if len(para) != 2 {
			return fmt.Errorf("please check the parameters %s", cmd)
		}
		return h.proxy.DoLogout(para[1])

	default:
		return fmt.Errorf("invalid command %s", cmd)
	}
	return nil
}
func (h *Helper) PickMessage() *Message {
	m := <-h.messageCh
	return m
}

func (h *Helper) Close() error {
	if h.proxy == nil {
		return nil
	}
	h.proxy.Stop()
	h.receiver = nil
	close(h.messageCh)

	return nil
}

func NewHelper() *Helper {
	helper := &Helper{proxy: nil}
	helper.messageCh = make(chan *Message, 10)
	helper.receiver = &receiver{ch: helper.messageCh}
	return helper
}

//==================================== END ======================================
