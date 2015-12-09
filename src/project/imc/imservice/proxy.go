//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: proxy.go
//     description:
//         created: 2015-12-06 19:28:22
//          author: wystan
//
//===============================================================================

package imservice

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/proto"
)

const (
	MSG_CONTINUE = iota
	MSG_STOP
)

const REQUEST_ACK_MARGIN = 100

type ImcProxy struct {
	writer *bufio.Writer
	reader *bufio.Reader
	con    net.Conn

	looper *msgLooper
	router *messageRouter

	ack     *ackReceiver
	notify  *notifyReceiver
	unknown *unauthorizedReceiver

	remotech chan *ImcCmd
}

type MessageReceiver interface {
	Receive(msg *ImcCmd) int
}

type messageRouter struct {
	receiver *list.List
	sync.RWMutex
}

type ackReceiver struct {
	/* data */
	outch         chan<- (*ImcCmd)
	cmdTypeNeeded CMD_TYPE
}
type unauthorizedReceiver struct {
	/* data */
}
type notifyReceiver struct {
	/* data */
	client MessageReceiver
}

func (r *ackReceiver) Receive(msg *ImcCmd) int {
	if *msg.CmdType != r.cmdTypeNeeded {
		return MSG_CONTINUE
	}
	r.outch <- msg
	return MSG_STOP
}
func (r *notifyReceiver) Receive(msg *ImcCmd) int {
	if r.client != nil {
		return r.client.Receive(msg)
	}
	return MSG_CONTINUE
}
func (r *unauthorizedReceiver) Receive(msg *ImcCmd) int {
	log.Printf("unauthorized msg : %#v", msg)
	return MSG_STOP
}

func newAckReceiver(ch chan<- (*ImcCmd)) *ackReceiver {
	return &ackReceiver{outch: ch, cmdTypeNeeded: 0}
}
func newUnauthorizedReceiver() *unauthorizedReceiver {
	return &unauthorizedReceiver{}
}
func newNotifyReceiver(client MessageReceiver) *notifyReceiver {
	return &notifyReceiver{client: client}
}

func (router *messageRouter) registerReceiver(receiver MessageReceiver) {
	router.Lock()
	defer router.Unlock()
	router.receiver.PushBack(receiver)
}
func (router *messageRouter) HandleMessage(data []byte) error {
	router.RLock()
	defer router.RUnlock()

	msg := &ImcCmd{}
	proto.Unmarshal(data, msg)
	for e := router.receiver.Front(); e != nil; e = e.Next() {
		receiver := e.Value.(MessageReceiver)
		if ret := receiver.Receive(msg); ret == MSG_STOP {
			break
		}
	}
	msg = nil
	return nil
}
func (router *messageRouter) Exit() {
	router.Lock()
	defer router.Unlock()
	for {
		e := router.receiver.Front()
		if e == nil {
			break
		}
		router.receiver.Remove(e)
	}
}

func newMessageRouter() *messageRouter {
	return &messageRouter{receiver: list.New()}
}

func (p *ImcProxy) waitForResponse(reader *bufio.Reader) *ImcCmd {
	cmd := <-p.remotech
	return cmd
}

func (p *ImcProxy) sendAndWait(c *ImcCmd) error {
	// TODO check the whether Start is called

	data, err := proto.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}

	p.ack.cmdTypeNeeded = *c.CmdType + REQUEST_ACK_MARGIN
	writeMsgData(p.writer, data)

	ack := p.waitForResponse(p.reader)
	if *ack.AckCommon.Status != RET_CODE_SUCCESS {
		return fmt.Errorf("%s", *ack.AckCommon.ErrorDesc)
	}
	return nil
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

	return p.sendAndWait(c)
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

	return p.sendAndWait(c)
}
func (p *ImcProxy) DoLogout(name string) error {
	cmd := CMD_TYPE_LOGOUT
	c := &ImcCmd{
		CmdType: &cmd,
		Logout: &CmdLogout{
			UserName: proto.String(name),
		},
	}

	return p.sendAndWait(c)
}
func (p *ImcProxy) DoSendMsg(peer, msg string) error {
	cmd := CMD_TYPE_SENDMSG
	c := &ImcCmd{
		CmdType: &cmd,
		SendMsg: &CmdSendMsg{
			PeerName: proto.String(peer),
			MsgBody:  proto.String(msg),
		},
	}

	return p.sendAndWait(c)
}

func (p *ImcProxy) Start() error {
	if p.looper == nil {
		p.looper = newMsgLooper(p.router, p.con)
	}

	// allocate 3 receivers
	p.ack = newAckReceiver(p.remotech)
	p.notify = newNotifyReceiver(nil)
	p.unknown = newUnauthorizedReceiver()

	p.router.registerReceiver(p.ack)
	p.router.registerReceiver(p.notify)
	p.router.registerReceiver(p.unknown)

	go p.looper.Loop()
	return nil
}
func (p *ImcProxy) Stop() {
	if p.looper != nil {
		p.looper.Exit()
	}
	p.looper = nil
	p.reader = nil
	p.writer = nil
	p.con = nil
	p.router = nil
	close(p.remotech)
}

func NewImcProxy(con net.Conn) *ImcProxy {
	r := bufio.NewReader(con)
	w := bufio.NewWriter(con)

	p := &ImcProxy{
		reader:   r,
		writer:   w,
		con:      con,
		router:   newMessageRouter(),
		looper:   nil,
		ack:      nil,
		notify:   nil,
		unknown:  nil,
		remotech: make(chan *ImcCmd),
	}

	return p
}

//==================================== END ======================================
